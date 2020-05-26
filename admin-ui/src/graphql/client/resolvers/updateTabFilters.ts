import ApolloClient from 'apollo-client';
import { NormalizedCacheObject } from 'apollo-cache-inmemory';
import {
  UpdateTabFiltersVariables,
  UpdateTabFiltersInput_newFilters
} from '../mutations/updateTabFilters.graphql';
import {
  GET_LOG_TABS,
  GetLogTabs,
  GetLogTabs_logTabs
} from '../queries/getLogs.graphql';
import { dateFilterOptions } from '../../../pages/Version/pages/Status/LogsPanel/components/Filters/components/DatesFilter/DateFilter';
import moment from 'moment';
import { ProcessSelection, LogPanelFilters } from '../typeDefs';

export const defaultFilters: {
  [key: string]: string | ProcessSelection[] | null;
} = {
  dateOption: dateFilterOptions.lastTwentyFourHours,
  startDate: moment()
    .subtract(1, 'day')
    .startOf('day')
    .toISOString(true),
  endDate: moment()
    .endOf('day')
    .toISOString(true),
  search: '',
  processes: [],
  level: null
};

export function getDefaultFilters():
  | LogPanelFilters
  | UpdateTabFiltersInput_newFilters {
  return {
    ...defaultFilters,
    __typename: 'logTabFilters'
  };
}

export default function updateTabFilters(
  _: any,
  variables: UpdateTabFiltersVariables,
  { cache }: ApolloClient<NormalizedCacheObject>
) {
  const { tabId, newFilters } = variables.input;
  const data = cache.readQuery<GetLogTabs>({
    query: GET_LOG_TABS
  });
  const logTabs = data?.logTabs;

  if (logTabs) {
    const updatedTabs = logTabs.map((logTab: GetLogTabs_logTabs) => {
      if (logTab.uniqueId === tabId) {
        return {
          ...logTab,
          filters: newFilters
        };
      }
      return logTab;
    });

    cache.writeData({
      data: {
        logTabs: updatedTabs
      }
    });
  }

  return null;
}