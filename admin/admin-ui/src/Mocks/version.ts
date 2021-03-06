import {
  NodeStatus,
  RuntimeStatus,
  VersionStatus
} from 'Graphql/types/globalTypes';

import { loader } from 'graphql.macro';

const GetConfigurationVariablesQuery = loader(
  'Graphql/queries/getConfigurationVariables.graphql'
);
const GetVersionWorkflowsQuery = loader(
  'Graphql/queries/getVersionWorkflows.graphql'
);
const GetVersionNodeStatusQuery = loader(
  'Graphql/subscriptions/watchVersionNodeStatus.graphql'
);
const NodeStatusSubscription = loader(
  'Graphql/subscriptions/watchVersionNodeStatus.graphql'
);
const PublishVersionMutation = loader(
  'Graphql/mutations/publishVersion.graphql'
);
const StartVersionMutation = loader('Graphql/mutations/startVersion.graphql');
const StopVersionMutation = loader('Graphql/mutations/stopVersion.graphql');
const UnpublishVersionMutation = loader(
  'Graphql/mutations/unpublishVersion.graphql'
);

export const runtime = {
  id: 'runtimeId',
  name: 'runtimeName',
  description: 'Some Description',
  creationAuthor: { email: 'some@user.com' },
  status: RuntimeStatus.STARTED,
  measurementsUrl: 'measurementsUrl',
  __typename: 'Runtime'
};
export const version = {
  __typename: 'Version',
  id: 'versionId',
  name: 'versionName',
  description: 'versionDescription',
  status: VersionStatus.STARTED,
  creationDate: '2020/01/01',
  creationAuthor: {
    __typename: 'User',
    id: 'userId',
    email: 'user@domain.com'
  },
  publicationDate: '2020/01/01',
  publicationAuthor: {
    __typename: 'User',
    id: 'userId',
    email: 'user@domain.com'
  },
  config: {
    vars: [],
    completed: false
  },
  hasDoc: false,
  errors: []
};
export const workflow = {
  __typename: 'Workflow',
  id: 'workflowId',
  name: 'Workflow Nane',
  nodes: [
    {
      __typename: 'Node',
      id: 'nodeId',
      name: 'Node Name',
      status: NodeStatus.STARTED
    }
  ],
  edges: [
    {
      __typename: 'Edge',
      id: 'edgeId',
      fromNode: 'node1',
      toNode: 'node2'
    }
  ]
};

export const confVarsMock = {
  request: {
    query: GetConfigurationVariablesQuery,
    variables: {
      versionId: 'id'
    }
  },
  result: {
    data: {
      version: {
        status: 'STARTED',
        config: {
          completed: true,
          vars: [
            {
              key: 'var2',
              value: 'value1',
              type: 'VARIABLE'
            },
            {
              key: 'var1',
              value: 'value2',
              type: 'VARIABLE'
            },
            {
              key: 'var3',
              value: 'value3',
              type: 'FILE'
            }
          ]
        }
      }
    }
  }
};

export const workflowsMock = {
  request: {
    query: GetVersionWorkflowsQuery
  },
  result: {
    data: {
      version: {
        name: 'version01',
        status: 'STARTED',
        configurationCompleted: true,
        workflows: [
          {
            name: 'workflow01',
            nodes: [
              {
                id: 'node01',
                name: 'one node',
                status: 'STARTED'
              },
              {
                id: 'node02',
                name: 'another node',
                status: 'STARTED'
              }
            ],
            edges: [
              {
                id: 'edge01',
                fromNode: 'node01',
                toNode: 'node02'
              }
            ]
          }
        ]
      }
    }
  }
};

export const errorMorkflowsMock = {
  request: {
    query: GetVersionWorkflowsQuery
  },
  error: 'Some error',
  result: {
    data: null
  }
};

export const nodeStatus = {
  request: {
    query: GetVersionNodeStatusQuery
  },
  result: {
    data: {
      versionNodeStatus: [
        {
          date: '2020-01-01',
          nodeId: 'nodeId1',
          status: 'STARTED',
          message: 'some message 1'
        },
        {
          date: '2020-01-01',
          nodeId: 'nodeId2',
          status: 'STARTED',
          message: 'some message 2'
        }
      ]
    }
  }
};

export const nodeStatusMock = {
  request: {
    query: NodeStatusSubscription
  },
  result: {
    data: {
      versionNodeStatus: {
        date: '2020/01/01',
        nodeId: 'node02',
        status: 'STARTED',
        message: 'some message'
      }
    }
  }
};

export const publishVersionMock = {
  request: {
    query: PublishVersionMutation
  },
  result: {
    data: {
      version
    }
  }
};

export const startVersionMock = {
  request: {
    query: StartVersionMutation
  },
  result: {
    data: {
      version
    }
  }
};

export const stopVersionMock = {
  request: {
    query: StopVersionMutation
  },
  result: {
    data: {
      version
    }
  }
};

export const unpublishVersionMock = {
  request: {
    query: UnpublishVersionMutation
  },
  result: {
    data: {
      version
    }
  }
};
