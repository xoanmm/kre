import React, { useState, useEffect } from 'react';
import moment, { Moment } from 'moment';
import { isInclusivelyBeforeDay } from 'react-dates';

import InputLabel from '../InputLabel/InputLabel';
import InputError from '../InputError/InputError';
import ArrowRightIcon from '@material-ui/icons/ArrowForward';
import ArrowLeftIcon from '@material-ui/icons/ArrowBack';
import { DateRangePicker } from 'react-dates';
import CalendarExtension from './components/CalendarExtension';

export type Range = {
  label: string;
  startDate: Moment;
  endDate: Moment;
};
type GenerateActRangeParams = {
  unit: 'day' | 'week' | 'month' | 'year';
  date?: Moment;
  label?: string;
  prefix?: string;
};
function generateRange({
  unit,
  date = moment(),
  label = '',
  prefix = 'This'
}: GenerateActRangeParams): Range {
  return {
    label: label || `${prefix} ${unit}`,
    startDate: date.clone().startOf(unit),
    endDate: date.clone().endOf(unit)
  };
}
function generatePrevRange(
  unit: 'day' | 'week' | 'month' | 'year',
  label?: string
): Range {
  const units = `${unit}s` as 'days' | 'weeks' | 'months' | 'years';
  return generateRange({
    date: moment().subtract(1, units),
    prefix: 'Previous',
    label,
    unit
  });
}

// Generate quick ranges for calendar extension
function getRanges(): Range[] {
  return [
    generateRange({ unit: 'day', label: 'Today' }),
    generateRange({ unit: 'week' }),
    generateRange({ unit: 'month' }),
    generateRange({ unit: 'year' }),
    generatePrevRange('day', 'Yesterday'),
    generatePrevRange('week'),
    generatePrevRange('month'),
    generatePrevRange('year'),
    {
      label: 'Last 2 days',
      startDate: moment()
        .subtract(1, 'days')
        .startOf('day'),
      endDate: moment().endOf('day')
    },
    {
      label: 'Last 6 months',
      startDate: moment()
        .subtract(6, 'months')
        .startOf('day'),
      endDate: moment().endOf('month')
    }
  ];
}

type DateKey = 'startDate' | 'endDate' | null;
type Dates = {
  startDate: Moment | null;
  endDate: Moment | null;
};
type Props = {
  label?: string;
  error?: string | boolean;
  formFromDate?: Moment | null;
  formToDate?: Moment | null;
  onChangeFromDate?: Function;
  onChangeToDate?: Function;
  hideError?: boolean;
  addExtension?: boolean;
};

function Calendar({
  label = '',
  error = '',
  hideError = false,
  formFromDate = null,
  formToDate = null,
  onChangeFromDate = function() {},
  onChangeToDate = function() {},
  addExtension = false
}: Props) {
  const [fromDate, setFromDate] = useState<Moment | null>(null);
  const [toDate, setToDate] = useState<Moment | null>(null);
  const [focusedInput, setFocusedInput] = useState<DateKey>(null);

  useEffect(() => {
    setFromDate(formFromDate);
  }, [formFromDate]);

  useEffect(() => {
    setToDate(formToDate);
  }, [formToDate]);

  function changeFromDate(date: Moment | null) {
    setFromDate(date);
    onChangeFromDate(date);
  }

  function changeToDate(date: Moment | null) {
    setToDate(date);
    onChangeToDate(date);
  }

  function getStartOrEndTime(date: Moment, isEnd = false) {
    return isEnd ? date.endOf('day') : date.startOf('day');
  }

  /*
   * If it is the first time setting a date, sets the time to the
   * start or end of the day, otherwise inherits the time from
   * the previous date
   */
  function getDateAndTime(
    previousDate: Moment | null,
    newDate: Moment | null,
    isEnd = false
  ): Moment | null {
    let newDatetime: Moment | null = null;

    if (newDate !== null) {
      const noPrevDate = previousDate === null;
      newDatetime = noPrevDate
        ? getStartOrEndTime(newDate, isEnd)
        : newDate.set(getTimeUnits(previousDate));
    }

    return newDatetime;
  }

  function getTimeUnits(date: Moment | null) {
    return {
      hour: date?.get('hour'),
      minute: date?.get('minute'),
      second: date?.get('second')
    };
  }

  return (
    <div>
      <InputLabel text={label} />
      <DateRangePicker
        startDate={fromDate}
        startDateId="calendar_from_date"
        startDatePlaceholderText="From Date"
        endDate={toDate}
        endDateId="calendar_to_date"
        endDatePlaceholderText="To Date"
        displayFormat={addExtension ? 'MM/DD/YYYY HH:mm:ss' : 'MM/DD/YYYY'}
        onDatesChange={({
          startDate: newFromDate,
          endDate: newToDate
        }: Dates) => {
          let newFromDatetime = getDateAndTime(fromDate, newFromDate);
          let newToDatetime = getDateAndTime(toDate, newToDate, true);

          changeFromDate(newFromDatetime);
          changeToDate(newToDatetime);

          setFocusedInput(focusedInput);
        }}
        focusedInput={focusedInput}
        onFocusChange={(input: DateKey) => setFocusedInput(input)}
        customArrowIcon={<ArrowRightIcon className="icon-small" />}
        navPrev={
          <div className="DayPickerNavigation_button DayPickerNavigation_button_1 DayPickerNavigation_button__default DayPickerNavigation_button__default_2 DayPickerNavigation_button__horizontal DayPickerNavigation_button__horizontal_3 DayPickerNavigation_button__horizontalDefault DayPickerNavigation_button__horizontalDefault_4 DayPickerNavigation_leftButton__horizontalDefault DayPickerNavigation_leftButton__horizontalDefault_5">
            <ArrowLeftIcon className="icon-small" style={{ margin: '0 8px' }} />
          </div>
        }
        navNext={
          <div className="DayPickerNavigation_button DayPickerNavigation_button_1 DayPickerNavigation_button__default DayPickerNavigation_button__default_2 DayPickerNavigation_button__horizontal DayPickerNavigation_button__horizontal_3 DayPickerNavigation_button__horizontalDefault DayPickerNavigation_button__horizontalDefault_4 DayPickerNavigation_rightButton__horizontalDefault DayPickerNavigation_rightButton__horizontalDefault_5">
            <ArrowRightIcon
              className="icon-small"
              style={{ margin: '0 8px' }}
            />
          </div>
        }
        isOutsideRange={day => !isInclusivelyBeforeDay(day, moment())}
        minimumNights={0}
        renderCalendarInfo={() =>
          addExtension ? (
            <CalendarExtension
              ranges={getRanges()}
              changeFromDate={changeFromDate}
              changeToDate={changeToDate}
              formFromDate={fromDate}
              formToDate={toDate}
            />
          ) : (
            <></>
          )
        }
      />
      {hideError === false && <InputError message={error.toString()} />}
    </div>
  );
}

export default Calendar;
