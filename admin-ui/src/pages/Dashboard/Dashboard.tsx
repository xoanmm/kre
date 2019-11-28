import { get } from 'lodash';

import React from 'react';
import { useHistory } from 'react-router';
import { History } from 'history';
import * as PAGES from '../../constants/routes';

import Header from '../../components/Header/Header';
import NavigationBar from '../../components/NavigationBar/NavigationBar';
import HexagonPanel from '../../components/Layout/HexagonPanel/HexagonPanel';
import Hexagon from '../../components/Shape/Hexagon/Hexagon';
import AlertMessage from '../../components/Alert/Alert';
import Spinner from '../../components/Spinner/Spinner';
import ErrorMessage from '../../components/ErrorMessage/ErrorMessage';
import Button from '../../components/Button/Button';

import styles from './Dashboard.module.scss';

import { useQuery } from '@apollo/react-hooks';
import { GET_DASHBOARD, GetDashboardResponse } from './Dashboard.graphql';
import { Alert, Runtime } from '../../graphql/models';
import { ApolloError } from 'apollo-client';

type Props = {
  data: GetDashboardResponse;
  error?: ApolloError;
  loading: boolean;
  history: History;
};

function getDashboardContent({ data, error, loading, history }: Props) {
  if (error) return <ErrorMessage />;
  if (loading) return <Spinner />;

  const runtimes = data.runtimes.map((runtime: Runtime, idx: number) => (
    <Hexagon
      key={`runtimeHexagon-${idx}`}
      onClick={() => {
        const runtimePath = PAGES.RUNTIME.replace(':runtimeId', runtime.id);
        history.push(runtimePath);
      }}
      id={runtime.id}
      status={runtime.status.toLowerCase()}
      title={runtime.name}
      info={[
        {
          type: 'active',
          date: runtime.creationDate
        }
      ]}
    />
  ));
  const alerts = data.alerts.map((alert: Alert, idx: number) => (
    <AlertMessage
      key={`runtimeAlert-${idx}`}
      type={alert.type}
      message={alert.message}
      runtimeId={alert.runtime.id}
    />
  ));

  return (
    <>
      <div>{alerts}</div>
      <HexagonPanel>{runtimes}</HexagonPanel>
    </>
  );
}

function Dashboard() {
  const history = useHistory();
  const { data, loading, error } = useQuery(GET_DASHBOARD);
  const nRuntimes = get(data, 'runtimes', []).length;

  return (
    <>
      <Header>
        <Button label="ADD RUNTIME" to={PAGES.NEW_RUNTIME} height={40} />
        <div>{`${nRuntimes} runtimes shown`}</div>
      </Header>
      <div className={styles.container} data-testid="dashboardContainer">
        <NavigationBar />
        <div className={styles.content}>
          <div className={styles.hexagons}>
            {getDashboardContent({ data, error, loading, history })}
          </div>
        </div>
      </div>
    </>
  );
}

export default Dashboard;
