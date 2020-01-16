import React from 'react';
import { Route, Switch, useParams } from 'react-router-dom';
import * as ROUTE from '../../constants/routes';

import StatusIcon from '@material-ui/icons/DeviceHub';
import MetricsIcon from '@material-ui/icons/ShowChart';
import DocumentationIcon from '@material-ui/icons/Toc';
import TimeIcon from '@material-ui/icons/AccessTime';
import ConfigIcon from '@material-ui/icons/Settings';

import RuntimeStatus from './pages/RuntimeStatus/RuntimeStatus';
import RuntimeStatusPreview from './pages/RuntimeStatusPreview/RuntimeStatusPreview';
import RuntimeVersions from './pages/RuntimeVersions/RuntimeVersions';
import RuntimeConfiguration from './pages/RuntimeConfiguration/RuntimeConfiguration';
import SpinnerCircular from '../../components/LoadingComponents/SpinnerCircular/SpinnerCircular';
import ErrorMessage from '../../components/ErrorMessage/ErrorMessage';
import Header from '../../components/Header/Header';
import NavigationBar from '../../components/NavigationBar/NavigationBar';
import Button from '../../components/Button/Button';
import Sidebar from '../../components/Sidebar/Sidebar';
import { Tab } from '../../components/NavBar/NavBar';
import SidebarTitle from './components/SidebarTitle/SidebarTitle';

import { useQuery } from '@apollo/react-hooks';
import {
  GET_RUNTIME,
  GetRuntimeResponse,
  GetRuntimeVars,
  GetVersionConfStatusResponse,
  GetVersionConfStatusVars,
  GET_VERSION_CONF_STATUS
} from './Runtime.graphql';

import styles from './Runtime.module.scss';

function updateTab(
  label: string,
  tabs: Tab[],
  updateFunc: (t: Tab) => void
): Tab[] {
  return tabs.map((tab: Tab) => {
    let tabCp = { ...tab };

    if (tab.label === label) {
      updateFunc(tabCp);
    }

    return tabCp;
  });
}

function disableTab(label: string, tabs: Tab[]): Tab[] {
  return updateTab(label, tabs, function(tab: Tab) {
    tab.disabled = true;
  });
}
function addWarningToTab(label: string, tabs: Tab[], message: string): Tab[] {
  return updateTab(label, tabs, function(tab: Tab) {
    tab.showWarning = true;
    tab.warningTitle = message;
  });
}

function createNavTabs(runtimeId: string, versionId: string): Tab[] {
  const navTabs = [
    {
      label: 'STATUS',
      route: ROUTE.RUNTIME_STATUS,
      Icon: StatusIcon,
      exact: false
    },
    {
      label: 'METRICS',
      route: ROUTE.HOME,
      Icon: MetricsIcon,
      disabled: true
    },
    {
      label: 'DOCUMENTATION',
      route: ROUTE.HOME,
      Icon: DocumentationIcon,
      disabled: true
    },
    {
      label: 'VERSIONS',
      route: ROUTE.RUNTIME_VERSIONS,
      Icon: TimeIcon
    },
    {
      label: 'CONFIGURATION',
      route: ROUTE.RUNTIME_VERSION_CONFIGURATION,
      Icon: ConfigIcon
    }
  ];

  navTabs.forEach(n => {
    n.route = n.route
      .replace(':runtimeId', runtimeId)
      .replace(':versionId', versionId);
  });

  return navTabs;
}

function Runtime() {
  const { runtimeId, versionId } = useParams();
  const noVersion = versionId === undefined;
  const { data, loading, error } = useQuery<GetRuntimeResponse, GetRuntimeVars>(
    GET_RUNTIME,
    {
      variables: { runtimeId },
      fetchPolicy: 'no-cache'
    }
  );
  const {
    data: versionData,
    loading: versionLoading,
    error: versionError,
    refetch: refetchVersion
  } = useQuery<GetVersionConfStatusResponse, GetVersionConfStatusVars>(
    GET_VERSION_CONF_STATUS,
    {
      variables: { versionId },
      skip: noVersion,
      fetchPolicy: 'no-cache'
    }
  );

  if (error || versionError) return <ErrorMessage />;
  if (loading || versionLoading) return <SpinnerCircular />;

  const runtime = data && data.runtime;
  const activeVersion = runtime && runtime.activeVersion;

  let navTabs: Tab[] = createNavTabs(runtimeId || '', versionId || '');

  if (versionData && versionData.version.configurationCompleted === false) {
    navTabs = addWarningToTab(
      'CONFIGURATION',
      navTabs,
      'Configuration is not completed'
    );
  }

  if (noVersion) {
    navTabs = disableTab('CONFIGURATION', navTabs);
  }

  const newVersionRoute = ROUTE.NEW_VERSION.replace(
    ':runtimeId',
    runtimeId || ''
  );

  return (
    <>
      <Header>
        <Button label="ADD VERSION" height={40} to={newVersionRoute} />
      </Header>
      <div className={styles.container} data-testid="runtimeContainer">
        <NavigationBar />
        <Sidebar
          title="Runtime"
          subheader={<SidebarTitle version={activeVersion} />}
          tabs={navTabs}
        />
        <div className={styles.content}>
          <Switch>
            <Route
              exact
              path={ROUTE.RUNTIME_VERSION_STATUS}
              component={RuntimeStatusPreview}
            />
            <Route
              exact
              path={ROUTE.RUNTIME_STATUS}
              component={RuntimeStatus}
            />
            <Route
              exact
              path={ROUTE.RUNTIME_VERSIONS}
              component={RuntimeVersions}
            />
            <Route
              exact
              path={ROUTE.RUNTIME_VERSION_CONFIGURATION}
              render={props => (
                <RuntimeConfiguration
                  {...props}
                  refetchVersion={refetchVersion}
                />
              )}
            />
          </Switch>
          {/*<Route exact path={ROUTE.SETTINGS_SECURITY} component={SecuritySettings} />
          <Route exact path={ROUTE.SETTINGS_AUDIT} component={AuditSettings} /> */}
        </div>
      </div>
    </>
  );
}

export default Runtime;
