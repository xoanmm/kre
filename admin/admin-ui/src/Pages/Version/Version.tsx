import {
  GetVersionConfStatus_runtime,
  GetVersionConfStatus_versions
} from 'Graphql/queries/types/GetVersionConfStatus';
import React, { useEffect } from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';

import Configuration from './pages/Configuration/Configuration';
import Documentation from './pages/Documentation/Documentation';
import Metrics from './pages/Metrics/Metrics';
import { NodeStatus } from 'Graphql/types/globalTypes';
import ROUTE from 'Constants/routes';
import Status from './pages/Status/Status';
import VersionSideBar from './components/VersionSideBar/VersionSideBar';
import styles from './Version.module.scss';
import useOpenedVersion from 'Graphql/hooks/useOpenedVersion';

type Props = {
  versions?: GetVersionConfStatus_versions[];
  version?: GetVersionConfStatus_versions;
  runtime?: GetVersionConfStatus_runtime;
};

function Version({ versions, version, runtime }: Props) {
  const { updateVersion, updateEntrypointStatus } = useOpenedVersion();

  useEffect(() => {
    if (runtime && version) {
      updateVersion(runtime.name, version.name);
    }
  }, [version, runtime, updateVersion]);

  useEffect(
    () => () => updateEntrypointStatus(NodeStatus.STOPPED),
    // We only want to execute this after unmounting
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  return (
    <div className={styles.container} data-testid="runtimeContainer">
      <VersionSideBar runtime={runtime} versions={versions} version={version} />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={ROUTE.RUNTIME_VERSION_STATUS}
            render={props => (
              <Status {...props} version={version} runtime={runtime} />
            )}
          />
          <Redirect
            exact
            from={ROUTE.RUNTIME_VERSION_DOCUMENTATION}
            to={ROUTE.RUNTIME_VERSION_DOCUMENTATION + '/README.md'}
          />
          <Route
            path={ROUTE.RUNTIME_VERSION_DOCUMENTATION}
            component={Documentation}
          />
          <Route
            exact
            path={ROUTE.RUNTIME_VERSION_CONFIGURATION}
            component={Configuration}
          />
          <Route
            exact
            path={ROUTE.RUNTIME_VERSION_METRICS}
            render={props => (
              <Metrics {...props} runtime={runtime} version={version} />
            )}
          />
        </Switch>
      </div>
    </div>
  );
}

export default Version;
