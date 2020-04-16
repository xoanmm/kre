import { get } from 'lodash';
import React, { useEffect, useRef, useState, useCallback } from 'react';
import { loader } from 'graphql.macro';
import { useQuery } from '@apollo/react-hooks';
import { GET_LOGS } from '../../../../../../../graphql/client/queries/getLogs.graphql';
import LogItem from './LogItem';
import { GetServerLogs_logs_items } from '../../../../../../../graphql/queries/types/GetServerLogs';
import styles from './LogsList.module.scss';
import {
  GetServerLogs,
  GetServerLogsVariables
} from '../../../../../../../graphql/queries/types/GetServerLogs';
import moment from 'moment';
import LoadMore from './LoadMore';
import SpinnerLinear from '../../../../../../../components/LoadingComponents/SpinnerLinear/SpinnerLinear';
import { LocalState } from '../../../../../../..';
import { FilterTypes } from '../LogsTab/LogsTab';
const GetLogsSubscription = loader(
  '../../../../../../../graphql/subscriptions/getLogsSubscription.graphql'
);
const GetServerLogsQuery = loader(
  '../../../../../../../graphql/queries/getServerLogs.graphql'
);

const getFilters = ({ startDate, endDate }: FilterTypes) => ({
  startDate: startDate.toISOString(true),
  endDate: endDate.toISOString(true)
});

function scrollToBottom(component: HTMLDivElement) {
  if (component) {
    component.scrollTo({
      top: component.scrollHeight,
      behavior: 'smooth'
    });
  }
}

type Props = {
  nodeId: string;
  runtimeId: string;
  workflowId?: string;
  filterValues: FilterTypes;
  onNewLogs: Function;
  logs: GetServerLogs_logs_items[];
};

function LogsList({
  nodeId,
  runtimeId,
  workflowId = '',
  filterValues,
  logs,
  onNewLogs
}: Props) {
  const [nextPage, setNextPage] = useState<string>('');
  const listRef = useRef<HTMLDivElement>(null);
  const unsubscribeRef = useRef<Function | null>(null);

  const { data: localData } = useQuery<LocalState>(GET_LOGS);

  const { loading, fetchMore, subscribeToMore } = useQuery<
    GetServerLogs,
    GetServerLogsVariables
  >(GetServerLogsQuery, {
    variables: { workflowId, runtimeId, nodeId, ...getFilters(filterValues) },
    onCompleted: data => {
      onNewLogs(data.logs.items.reverse());
      setNextPage(data.logs.cursor || '');
      handleSubscription();
    },
    onError: handleSubscription,
    fetchPolicy: 'no-cache'
  });

  // Subscription query
  const subscribe = () =>
    subscribeToMore({
      document: GetLogsSubscription,
      variables: { runtimeId, nodeId },
      updateQuery: (prev, { subscriptionData }) => {
        const newLog = get(subscriptionData.data, 'nodeLogs');
        onNewLogs((oldLogs: GetServerLogs_logs_items[]) => [
          ...oldLogs,
          newLog
        ]);
        return prev;
      }
    });

  function handleSubscription() {
    const { endDate } = filterValues;
    if (moment().isBefore(endDate)) {
      unsubscribeRef.current && unsubscribeRef.current();
      unsubscribeRef.current = subscribe();
    }
  }

  const handleScroll = useCallback(() => {
    if (localData?.logsAutoScroll && listRef.current !== null) {
      scrollToBottom(listRef.current);
    }
  }, [localData, listRef]);

  useEffect(() => {
    handleScroll();
  }, [logs, localData?.logsAutoScroll]);

  if (loading) return <SpinnerLinear />;

  // Pagination query
  function loadPreviousLogs() {
    fetchMore({
      variables: {
        runtimeId,
        nodeId,
        cursor: nextPage
      },
      updateQuery: (prev, { fetchMoreResult }) => {
        const newData = fetchMoreResult && fetchMoreResult.logs;

        if (newData) {
          onNewLogs((oldLogs: GetServerLogs_logs_items[]) => [
            ...newData.items.reverse(),
            ...oldLogs
          ]);
          setNextPage(newData.cursor || '');
        }

        return prev;
      }
    });
  }

  const logElements = logs.map((log: GetServerLogs_logs_items) => (
    <LogItem {...log} key={log.id} />
  ));
  return (
    <div
      ref={listRef}
      className={styles.listContainer}
      id="VersionLogsListContainer"
    >
      {nextPage && <LoadMore onClick={loadPreviousLogs} />}
      {logElements}
    </div>
  );
}

export default LogsList;
