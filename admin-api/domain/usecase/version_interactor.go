package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/konstellation-io/kre/admin-api/domain/usecase/auth"
	"io"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/konstellation-io/kre/admin-api/domain/entity"
	"github.com/konstellation-io/kre/admin-api/domain/repository"
	"github.com/konstellation-io/kre/admin-api/domain/service"
	"github.com/konstellation-io/kre/admin-api/domain/usecase/krt"
	"github.com/konstellation-io/kre/admin-api/domain/usecase/logging"
)

const idLength = 10
const idCharset = "abcdefghijklmnopqrstuvwxyz"
const idCharsetLen = len(idCharset)

var (
	// ErrVersionNotFound error
	ErrVersionNotFound = errors.New("error version not found")
	// ErrVersionDuplicated error
	ErrVersionDuplicated = errors.New("error version duplicated")
	// ErrVersionConfigIncomplete error
	ErrVersionConfigIncomplete = errors.New("version config is incomplete")
	// ErrVersionConfigInvalidKey error
	ErrVersionConfigInvalidKey = errors.New("version config contains an unknown key")
	// ErrUpdatingRunningVersionConfig error
	ErrUpdatingStartedVersionConfig = errors.New("config can't be incomplete for started version")
	// ErrInvalidVersionStatusBeforeStarting error
	ErrInvalidVersionStatusBeforeStarting = errors.New("the version must be stopped before starting")
	// ErrInvalidVersionStatusBeforeStopping error
	ErrInvalidVersionStatusBeforeStopping = errors.New("the version must be started before stopping")
	// ErrInvalidVersionStatusBeforePublishing error
	ErrInvalidVersionStatusBeforePublishing = errors.New("the version must be started before publishing")
	// ErrInvalidVersionStatusBeforeUnpublishing error
	ErrInvalidVersionStatusBeforeUnpublishing = errors.New("the version must be published before unpublishing")
)

// VersionInteractor contains app logic about Version entities
type VersionInteractor struct {
	logger                 logging.Logger
	versionRepo            repository.VersionRepo
	runtimeRepo            repository.RuntimeRepo
	versionService         service.VersionService
	monitoringService      service.MonitoringService
	userActivityInteractor *UserActivityInteractor
	createStorage          repository.CreateStorage
	accessControl          auth.AccessControl
}

// NewVersionInteractor creates a new interactor
func NewVersionInteractor(
	logger logging.Logger,
	versionRepo repository.VersionRepo,
	runtimeRepo repository.RuntimeRepo,
	runtimeService service.VersionService,
	monitoringService service.MonitoringService,
	userActivityInteractor *UserActivityInteractor,
	createStorage repository.CreateStorage,
	accessControl auth.AccessControl,
) *VersionInteractor {
	return &VersionInteractor{
		logger,
		versionRepo,
		runtimeRepo,
		runtimeService,
		monitoringService,
		userActivityInteractor,
		createStorage,
		accessControl,
	}
}

// TODO refactor this in order to use a common id generator
func generateId() string {
	b := make([]byte, idLength)
	for i := range b {
		b[i] = idCharset[rand.Intn(idCharsetLen)]
	}
	return string(b)
}

func (i *VersionInteractor) filterConfigVars(loggedUserID string, version *entity.Version) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		version.Config.Vars = nil
	}
}

// GetByRuntime returns all Versions of the given Runtime
func (i *VersionInteractor) GetByRuntime(loggedUserID, runtimeID string) ([]*entity.Version, error) {
	versions, err := i.versionRepo.GetByRuntime(runtimeID)
	if err != nil {
		return nil, err
	}

	for _, v := range versions {
		i.filterConfigVars(loggedUserID, v)
	}

	return versions, nil
}

// GetByID returns a Version by its ID
func (i *VersionInteractor) GetByID(loggedUserID, id string) (*entity.Version, error) {
	v, err := i.versionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	i.filterConfigVars(loggedUserID, v)

	return v, nil
}

func (i *VersionInteractor) GetByIDs(ids []string) ([]*entity.Version, []error) {
	return i.versionRepo.GetByIDs(ids)
}

// Create creates a Version on the DB based on the content of a KRT file
func (i *VersionInteractor) Create(ctx context.Context, loggedUserID, runtimeID string, krtFile io.Reader) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, runtimeID)
	if err != nil {
		return nil, fmt.Errorf("error runtime repo GetByID: %w", err)
	}

	tmpDir, err := ioutil.TempDir("", "version")
	if err != nil {
		return nil, fmt.Errorf("error creating temp dir for version: %w", err)
	}
	i.logger.Info("Created temp dir to extract the KRT files at " + tmpDir)

	krtYml, err := krt.ProcessFile(i.logger, krtFile, tmpDir)
	if err != nil {
		return nil, fmt.Errorf("error processing krt: %w", err)
	}

	// Check if the version is duplicated
	versions, err := i.versionRepo.GetByRuntime(runtimeID)
	for _, v := range versions {
		if v.Name == krtYml.Version {
			return nil, ErrVersionDuplicated
		}
	}

	workflows, err := generateWorkflows(krtYml)
	if err != nil {
		return nil, fmt.Errorf("error generating workflows: %w", err)
	}

	existingConfig := readExistingConf(versions)
	config := fillNewConfWithExisting(existingConfig, krtYml)

	versionCreated, err := i.versionRepo.Create(loggedUserID, &entity.Version{
		RuntimeID:   runtimeID,
		Name:        krtYml.Version,
		Description: krtYml.Description,
		Config:      config,
		Entrypoint: entity.Entrypoint{
			ProtoFile: krtYml.Entrypoint.Proto,
			Image:     krtYml.Entrypoint.Image,
			Src:       krtYml.Entrypoint.Src,
		},
		Workflows: workflows,
	})
	if err != nil {
		return nil, err
	}
	i.logger.Info("Version created")

	err = i.storeContent(runtime, krtYml, tmpDir)
	if err != nil {
		return nil, err
	}

	err = i.userActivityInteractor.RegisterCreateAction(loggedUserID, runtime, versionCreated)
	if err != nil {
		return nil, fmt.Errorf("error registering activity: %w", err)
	}
	return versionCreated, nil
}

// Start create the resources of the given Version
func (i *VersionInteractor) Start(ctx context.Context, loggedUserID string, versionID string, comment string) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	i.logger.Infof("The user %s is starting version %s", loggedUserID, versionID)

	version, err := i.versionRepo.GetByID(versionID)
	if err != nil {
		return nil, err
	}

	if version.Status != entity.VersionStatusStopped {
		return nil, ErrInvalidVersionStatusBeforeStarting
	}

	if !version.Config.Completed {
		return nil, ErrVersionConfigIncomplete
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	err = i.versionService.Start(runtime, version)
	if err != nil {
		return nil, err
	}

	version.Status = entity.VersionStatusStarted // TODO: First go to Starting
	err = i.versionRepo.Update(version)
	if err != nil {
		return nil, err
	}

	err = i.userActivityInteractor.RegisterStartAction(loggedUserID, runtime, version, comment)
	if err != nil {
		return nil, err
	}
	return version, nil
}

// Stop removes the resources of the given Version
func (i *VersionInteractor) Stop(ctx context.Context, loggedUserID string, versionID string, comment string) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	i.logger.Infof("The user %s is stopping version %s", loggedUserID, versionID)

	version, err := i.versionRepo.GetByID(versionID)
	if err != nil {
		return nil, err
	}

	if version.Status != entity.VersionStatusStarted {
		return nil, ErrInvalidVersionStatusBeforeStopping
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	err = i.versionService.Stop(runtime, version)
	if err != nil {
		return nil, err
	}

	version.Status = entity.VersionStatusStopped
	err = i.versionRepo.Update(version)
	if err != nil {
		return nil, err
	}

	err = i.userActivityInteractor.RegisterStopAction(loggedUserID, runtime, version, comment)
	if err != nil {
		return nil, err
	}
	return version, nil
}

// Publish set a Version as published on DB and K8s
func (i *VersionInteractor) Publish(ctx context.Context, loggedUserID string, versionID string, comment string) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	i.logger.Infof("The user %s is publishing version %s", loggedUserID, versionID)

	version, err := i.versionRepo.GetByID(versionID)
	if err != nil {
		return nil, err
	}

	if version.Status != entity.VersionStatusStarted {
		return nil, ErrInvalidVersionStatusBeforePublishing
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	// Unpublish previous published version
	previousPublishedVersion, err := i.unpublishPreviousVersion(runtime)
	if err != nil {
		return nil, fmt.Errorf("error unpublishing previous version: %w", err)
	}

	err = i.versionService.Publish(runtime, version)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	version.PublicationDate = &now
	version.PublicationUserID = &loggedUserID
	version.Status = entity.VersionStatusPublished
	err = i.versionRepo.Update(version)
	if err != nil {
		return nil, err
	}

	err = i.userActivityInteractor.RegisterPublishAction(loggedUserID, runtime, version, previousPublishedVersion, comment)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (i *VersionInteractor) unpublishPreviousVersion(runtime *entity.Runtime) (*entity.Version, error) {
	prevVersion := &entity.Version{}
	versions, err := i.versionRepo.GetByRuntime(runtime.ID)
	if err != nil {
		return nil, err
	}
	if len(versions) > 0 {
		for _, v := range versions {
			if v.Status == entity.VersionStatusPublished {
				prevVersion = v
				v.Status = entity.VersionStatusStarted
				v.PublicationUserID = nil
				v.PublicationDate = nil
				err = i.versionRepo.Update(v)
				if err != nil {
					return nil, err
				}
				break
			}
		}
	}
	return prevVersion, nil
}

// Unpublish set a Version as not published on DB and K8s
func (i *VersionInteractor) Unpublish(ctx context.Context, loggedUserID string, versionID string, comment string) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	i.logger.Infof("The user %s is unpublishing version %s", loggedUserID, versionID)

	version, err := i.versionRepo.GetByID(versionID)
	if err != nil {
		return nil, err
	}

	if version.Status != entity.VersionStatusPublished {
		return nil, ErrInvalidVersionStatusBeforeUnpublishing
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	err = i.versionService.Unpublish(runtime, version)
	if err != nil {
		return nil, err
	}

	version.PublicationUserID = nil
	version.PublicationDate = nil
	version.Status = entity.VersionStatusStarted
	err = i.versionRepo.Update(version)
	if err != nil {
		return nil, err
	}

	err = i.userActivityInteractor.RegisterUnpublishAction(loggedUserID, runtime, version, comment)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (i *VersionInteractor) UpdateVersionConfig(ctx context.Context, loggedUserID string, version *entity.Version, config []*entity.ConfigurationVariable) (*entity.Version, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActEdit); err != nil {
		return nil, err
	}

	err := i.validateNewConfig(version.Config.Vars, config)
	if err != nil {
		return nil, err
	}

	isStarted := version.PublishedOrStarted()

	newConfig, newConfigIsComplete := generateNewConfig(version.Config.Vars, config)

	if isStarted && newConfigIsComplete == false {
		return nil, ErrUpdatingStartedVersionConfig
	}

	version.Config.Vars = newConfig
	version.Config.Completed = newConfigIsComplete

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	// No need to call runtime-api if there are no resources running
	if isStarted {
		err = i.versionService.UpdateConfig(runtime, version)
		if err != nil {
			return nil, err
		}
	}

	err = i.versionRepo.Update(version)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (i *VersionInteractor) WatchVersionStatus(ctx context.Context, loggedUserID, versionId string, stopCh <-chan bool) (<-chan *entity.Node, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResVersion, auth.ActView); err != nil {
		return nil, err
	}

	version, err := i.versionRepo.GetByID(versionId)
	if err != nil {
		return nil, err
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, version.RuntimeID)
	if err != nil {
		return nil, err
	}

	return i.monitoringService.VersionStatus(runtime, version.Name, stopCh)
}

func (i *VersionInteractor) WatchNodeLogs(
	ctx context.Context,
	loggedUserID, runtimeID, versionID string,
	filters entity.LogFilters,
	stopChannel chan bool,
) (<-chan *entity.NodeLog, error) {
	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResLogs, auth.ActView); err != nil {
		return nil, err
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, runtimeID)
	if err != nil {
		return nil, err
	}

	return i.monitoringService.NodeLogs(runtime, versionID, filters, stopChannel)
}

func (i *VersionInteractor) SearchLogs(
	ctx context.Context,
	loggedUserID string,
	runtimeID string,
	versionID string,
	filters entity.LogFilters,
	cursor *string,
) (entity.SearchLogsResult, error) {
	var result entity.SearchLogsResult

	if err := i.accessControl.CheckPermission(loggedUserID, auth.ResLogs, auth.ActView); err != nil {
		return result, err
	}

	runtime, err := i.runtimeRepo.GetByID(ctx, runtimeID)
	if err != nil {
		return result, err
	}

	return i.monitoringService.SearchLogs(ctx, runtime, versionID, filters, cursor)
}
