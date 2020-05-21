import React, { useEffect } from 'react';
import { Route, Switch } from 'react-router-dom';
import ROUTE from '../../constants/routes';

import Status from './pages/Status/Status';
import Configuration from './pages/Configuration/Configuration';
import Metrics from './pages/Metrics/Metrics';
import Documentation from './pages/Documentation/Documentation';
import VersionSideBar from './components/VersionSideBar/VersionSideBar';

import {
  GetVersionConfStatus_versions,
  GetVersionConfStatus_runtime
} from '../../graphql/queries/types/GetVersionConfStatus';

import styles from './Version.module.scss';
import { useApolloClient } from '@apollo/react-hooks';

type Props = {
  version?: GetVersionConfStatus_versions;
  runtime?: GetVersionConfStatus_runtime;
};

function Version({ version, runtime }: Props) {
  const client = useApolloClient();

  useEffect(() => {
    if (runtime && version) {
      client.writeData({
        data: {
          openedVersion: {
            runtimeName: runtime.name,
            versionName: version.name,
            __typename: 'OpenedVersion'
          }
        }
      });
    }
  }, [version, runtime, client]);

  return (
    <div className={styles.container} data-testid="runtimeContainer">
      <VersionSideBar runtime={runtime} version={version} />
      <div className={styles.content}>
        <Switch>
          <Route
            exact
            path={ROUTE.RUNTIME_VERSION_STATUS}
            render={props => <Status {...props} version={version} />}
          />
          <Route
            exact
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
