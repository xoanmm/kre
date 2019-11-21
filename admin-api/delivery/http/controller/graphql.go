package controller

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/config"
	"gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/gql"
	"gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/usecase"
	"gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/usecase/logging"
)

type GraphQLController struct {
	cfg               *config.Config
	logger            logging.Logger
	runtimeInteractor *usecase.RuntimeInteractor
	userInteractor    *usecase.UserInteractor
}

func NewGraphQLController(
	cfg *config.Config,
	logger logging.Logger,
	runtimeInteractor *usecase.RuntimeInteractor,
	userInteractor *usecase.UserInteractor,
) *GraphQLController {
	return &GraphQLController{
		cfg,
		logger,
		runtimeInteractor,
		userInteractor,
	}
}

func (g *GraphQLController) GraphQLHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	g.logger.Info("Request from user " + userID)

	h := gql.NewHttpHandler(g.logger, g.runtimeInteractor, g.userInteractor)

	r := c.Request()
	ctx := context.WithValue(r.Context(), "userID", userID)

	h.ServeHTTP(c.Response(), r.WithContext(ctx))

	return nil
}
