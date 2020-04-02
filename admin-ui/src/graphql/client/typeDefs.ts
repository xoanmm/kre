import gql from 'graphql-tag';

export enum NotificationType {
  MESSAGE = 'MESSAGE',
  ERROR = 'ERROR'
}

export interface AddNotificationInput {
  id: string;
  message: string;
  type: NotificationType;
  timeout: number;
  to: string;
}

export interface SetCurrentLogPanelInput {
  runtimeId: string;
  nodeId: string;
  nodeName: string;
}

export interface LogPanel extends SetCurrentLogPanelInput {
  __typename: string;
}

export interface RemoveNotificationInput {
  id: string;
}

const typeDefs = gql`
  extend type Query {
    notifications: [Notification!]!
  }

  extend type Mutation {
    addNotification(input: AddNotificationInput!): [Notification!]!
    removeNotification(input: RemoveNotificationInput!): [Notification!]!
  }

  extend type Notification {
    id: ID!
    message: String!
    type: NotificationType!
    to: String!
    timeout: Number!
  }

  extend type AddNotificationInput {
    id: ID!
    message: String!
    type: NotificationType!
    timeout: Number
    to: String!
  }

  extend type RemoveNotificationInput {
    id: ID!
  }

  extend enum NotificationType {
    MESSAGE
    ERROR
  }
`;

export default typeDefs;