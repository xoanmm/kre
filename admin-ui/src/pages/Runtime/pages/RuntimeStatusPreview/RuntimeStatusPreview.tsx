import { get, cloneDeep } from 'lodash';

import React, { useRef, useState, useEffect } from 'react';
import { useParams } from 'react-router';

import HorizontalBar from '../../../../components/Layout/HorizontalBar/HorizontalBar';
import Button from '../../../../components/Button/Button';
import SpinnerCircular from '../../../../components/LoadingComponents/SpinnerCircular/SpinnerCircular';
import ErrorMessage from '../../../../components/ErrorMessage/ErrorMessage';
import VersionStatusViewer from '../../../../components/VersionStatusViewer/VersionStatusViewer';
import Node, { TYPES, STATUS } from '../../../../components/Shape/Node/Node';

import { useMutation, useQuery } from '@apollo/react-hooks';
import {
  ACTIVATE_VERSION,
  DEPLOY_VERSION,
  GET_VERSION_WORKFLOWS,
  ActivateVersionResponse,
  DeployVersionResponse,
  ActivateDeployVersionVars,
  GetVersionWorkflowsResponse,
  GetVersionWorkflowsVars
} from './RuntimeStatusPreview.graphql';

import styles from './RuntimeStatusPreview.module.scss';

const data = [
  {
    name: 'MAKE_PREDICTION',
    nodes: [
      {
        id: 'W1InputNode',
        name: 'TICKET ASSET',
        status: '',
        type: TYPES.INPUT
      },
      {
        id: 'W1InnerNode1',
        name: 'TICKET STATUS TRANSFORMER',
        status: '',
        type: TYPES.DEFAULT_2
      },
      {
        id: 'W1InnerNode2',
        name: 'TICKET STATUS NORMALIZATOR',
        status: '',
        type: TYPES.DEFAULT_2
      },
      {
        id: 'W1InnerNode3',
        name: 'TICKET CLASIFICATOR NN',
        status: '',
        type: TYPES.DEFAULT_2
      },
      {
        id: 'W1OutputNode',
        name: 'TNBA ORDERED',
        status: '',
        type: TYPES.OUTPUT
      }
    ],
    edges: [
      {
        id: 'Edge1',
        name: 'Edge1',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W1InputNode',
        to: 'W1InnerNode1'
      },
      {
        id: 'Edge2',
        name: 'Edge2',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W1InnerNode1',
        to: 'W1InnerNode2'
      },
      {
        id: 'Edge3',
        name: 'Edge3',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W1InnerNode2',
        to: 'W1InnerNode3'
      },
      {
        id: 'Edge4',
        name: 'Edge4',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W1InnerNode3',
        to: 'W1OutputNode'
      }
    ]
  },
  {
    name: 'SAVE_CLIENT_METRICS',
    nodes: [
      {
        id: 'W2InputNode',
        name: 'TICKET ASSET',
        status: '',
        type: TYPES.INPUT
      },
      {
        id: 'W2InnerNode1',
        name: 'TICKET STATUS TRANSFORMER',
        status: '',
        type: TYPES.DEFAULT
      },
      {
        id: 'W2InnerNode2',
        name: 'TICKET STATUS NORMALIZATOR',
        status: '',
        type: TYPES.DEFAULT
      },
      {
        id: 'W2OutputNode',
        name: 'TNBA ORDERED',
        status: '',
        type: TYPES.OUTPUT
      }
    ],
    edges: [
      {
        id: 'Edge1',
        name: 'Edge1',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W2InputNode',
        to: 'W2InnerNode1'
      },
      {
        id: 'Edge2',
        name: 'Edge2',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W2InnerNode1',
        to: 'W2InnerNode2'
      },
      {
        id: 'Edge3',
        name: 'Edge3',
        status: STATUS.ACTIVE,
        value: 0,
        from: 'W2InnerNode2',
        to: 'W2OutputNode'
      }
    ]
  }
];

function formatData(workflows: any) {
  let formattedData = cloneDeep(workflows);

  formattedData = formattedData.map((workflow: any, idx: number) => {
    workflow.nodes.unshift({
      id: `W${idx}InputNode`,
      name: 'DATA INPUT',
      status: '',
      type: TYPES.INPUT
    });
    workflow.nodes.push({
      id: `W${idx}OutputNode`,
      name: 'DATA OUTPUT',
      status: '',
      type: TYPES.OUTPUT
    });
    workflow.edges.push({
      id: 'InputEdge',
      status: STATUS.ACTIVE,
      fromNode: `W${idx}InputNode`,
      toNode: workflow.nodes[1].id
    });
    workflow.edges.push({
      id: 'OutputEdge',
      status: STATUS.ACTIVE,
      fromNode: workflow.nodes[workflow.nodes.length - 2].id,
      toNode: `W${idx}OutputNode`
    });

    return workflow;
  });

  return formattedData;
}

function Viewer({ data }: any) {
  const container = useRef(null);
  const [dimensions, setDimensions] = useState({
    width: 0,
    height: 0
  });

  useEffect(() => {
    const containerDOM = container.current;

    if (containerDOM) {
      setDimensions({
        // @ts-ignore
        width: containerDOM.clientWidth,
        // @ts-ignore
        height: containerDOM.clientHeight
      });
    }
  }, [container]);

  const { width, height } = dimensions;

  return (
    <div ref={container} className={styles.container}>
      <VersionStatusViewer
        width={width}
        height={height * 0.6}
        margin={{
          top: 10,
          right: 10,
          bottom: 10,
          left: 10
        }}
        data={data}
        preview={false}
      />
    </div>
  );
}

function RuntimeStatusPreview() {
  const { versionId } = useParams();
  const { data, loading, error } = useQuery<
    GetVersionWorkflowsResponse,
    GetVersionWorkflowsVars
  >(GET_VERSION_WORKFLOWS, { variables: { versionId } });
  // TODO: loading and error check
  const [activateMutation] = useMutation<
    ActivateVersionResponse,
    ActivateDeployVersionVars
  >(ACTIVATE_VERSION);
  const [deployMutation] = useMutation<
    DeployVersionResponse,
    ActivateDeployVersionVars
  >(DEPLOY_VERSION);

  if (error) return <ErrorMessage />;
  if (loading) return <SpinnerCircular />;

  function getMutationVars() {
    return {
      variables: {
        input: {
          versionId: versionId || ''
        }
      }
    };
  }
  function onDeployVersion() {
    deployMutation(getMutationVars());
  }
  function onActivateVersion() {
    activateMutation(getMutationVars());
  }

  return (
    <div className={styles.container}>
      <HorizontalBar>
        <Button label="DEPLOY" onClick={onDeployVersion} />
        <Button label="ACTIVATE" onClick={onActivateVersion} />
      </HorizontalBar>
      STATUS PREVIEW
      <Viewer data={formatData(get(data, 'version.workflows', []))} />
      <Node type={TYPES.INPUT} status={STATUS.INACTIVE} />
      <Node type={TYPES.DEFAULT} status={STATUS.INACTIVE} />
      <Node type={TYPES.DEFAULT_2} status={STATUS.INACTIVE} />
      <Node type={TYPES.OUTPUT} status={STATUS.INACTIVE} />
    </div>
  );
}

export default RuntimeStatusPreview;
