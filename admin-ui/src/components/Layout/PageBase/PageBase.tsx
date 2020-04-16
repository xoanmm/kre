import React, { ReactElement, useEffect, useState } from 'react';
import styles from './PageBase.module.scss';
import Header from '../../Header/Header';
import NavigationBar from '../../NavigationBar/NavigationBar';
import cx from 'classnames';
import { useQuery } from '@apollo/react-hooks';
import { GET_LOG_TABS } from '../../../graphql/client/queries/getLogs.graphql';

type PageBaseProps = {
  children: ReactElement | ReactElement[] | null;
  headerChildren?: ReactElement | ReactElement[] | null;
  customClassname?: string;
};

function PageBase({
  children,
  headerChildren,
  customClassname
}: PageBaseProps) {
  const [opened, setOpened] = useState<boolean>(false);

  const { data: localData } = useQuery(GET_LOG_TABS);

  useEffect(() => {
    setOpened(localData && localData.logTabs && !!localData.logTabs.length);
  }, [localData]);
  return (
    <>
      <Header>{headerChildren}</Header>
      <div
        className={cx(styles.content, customClassname, {
          [styles.withLogs]: opened
        })}
      >
        <NavigationBar />
        {children}
      </div>
    </>
  );
}

export default PageBase;
