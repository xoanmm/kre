import React from 'react';
import Chip from '../../../../../../../components/Chip/Chip';
import { isEditable, ProcessChip } from './AppliedFilters';
import moment from 'moment';

const filterToLabel = new Map([
  ['workflowId', 'Workflow'],
  ['nodeName', 'Process'],
  ['level', 'Level'],
  ['search', 'Search'],
  ['processes', 'Process'],
  ['startDate', 'From'],
  ['endDate', 'To']
]);
function getFilterLabel(filter: string) {
  return filterToLabel.get(filter);
}

function getValueLabel(filter: string, value: string | ProcessChip) {
  if (filter === 'startDate' || filter === 'endDate') {
    if (typeof value === 'string')
      return moment(value).format('MMM DD, YYYY HH:mm');
  }
  if (filter === 'processes') {
    if (typeof value !== 'string')
      return `${value.processName} [${value.workflowName}]`;
  }

  return value;
}

type Props = {
  filter: string;
  value: string;
  removeFilter: Function;
};
export function Filter({ filter, value, removeFilter }: Props) {
  const label = `${getFilterLabel(filter)}: ${getValueLabel(filter, value)}`;

  function onClose() {
    removeFilter(filter, value);
  }
  return (
    <Chip label={label} onClose={isEditable(filter) ? onClose : undefined} />
  );
}

export default Filter;