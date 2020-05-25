import React, { useState, useRef } from 'react';

import InputLabel from '../InputLabel/InputLabel';
import InputError from '../InputError/InputError';
import Selections from './Selections';
import Group from './Group';

import cx from 'classnames';
import styles from './GroupSelect.module.scss';
import { get, isEmpty } from 'lodash';

const MAX_HEIGHT = 400;
const GROUP_HEIGHT = 30;
const GROUP_ELEMENT_HEIGHT = 45;

export type GroupSelectData = { [key: string]: string[] };

type Props = {
  onChange?: (selection: GroupSelectData) => void;
  label?: string;
  error?: string;
  placeholder?: string;
  options: GroupSelectData;
  formSelectedOptions?: GroupSelectData;
  hideError?: boolean;
  className?: string;
  hideSelections?: boolean;
};

function GroupSelect({
  options,
  onChange = function() {},
  label = '',
  error = '',
  placeholder = '',
  formSelectedOptions,
  hideError = false,
  className = '',
  hideSelections = false
}: Props) {
  const inputEl = useRef<HTMLInputElement>(null);
  const containerEl = useRef<HTMLDivElement>(null);
  const [optionsOpened, setOptionsOpened] = useState(false);

  /*
   * Adds or removes event listeners and updates options visibility
   */
  function changeOptionsState(show: boolean = false) {
    const listenerAction = show ? 'addEventListener' : 'removeEventListener';

    document[listenerAction]('contextmenu', handleClickOutside);
    document[listenerAction]('click', handleClickOutside);

    if (containerEl.current !== null) {
      containerEl.current[listenerAction]('scroll', closeOptions);
    }

    setOptionsOpened(show);
  }

  function openOptions() {
    if (!optionsOpened) changeOptionsState(true);
  }

  function closeOptions() {
    changeOptionsState(false);
  }

  function handleClickOutside(e: Event) {
    const target = e.target as HTMLElement;

    // Has the user clicked outside the selector?
    if (document.contains(target) && !containerEl.current?.contains(target)) {
      closeOptions();
    }
  }

  function onClear() {
    onChange({});
  }

  function onSelectGroup(group: string) {
    const newSelections = {
      ...formSelectedOptions,
      [group]: [...options[group]]
    };

    onChange(newSelections);
  }

  function onDeselectGroup(group: string) {
    const newSelections = { ...formSelectedOptions };
    delete newSelections[group];

    onChange(newSelections);
  }

  function onSelect(group: string, element: string) {
    const newSelections = {
      ...formSelectedOptions,
      [group]: [...get(formSelectedOptions, group, []), element]
    };

    onChange(newSelections);
  }
  function onDeselect(group: string, element: string) {
    const groupElements = [...get(formSelectedOptions, group, [])];
    groupElements.splice(groupElements.indexOf(element), 1);

    const newSelections = {
      ...formSelectedOptions,
      [group]: groupElements
    };

    if (isEmpty(groupElements)) {
      delete newSelections[group];
    }

    onChange(newSelections);
  }

  const optionList = Object.entries(options).map(([group, elements]) => (
    <Group
      key={group}
      group={group}
      elements={elements}
      selections={formSelectedOptions}
      onSelectGroup={onSelectGroup}
      onDeselectGroup={onDeselectGroup}
      onSelect={onSelect}
      onDeselect={onDeselect}
    />
  ));

  const nGroups = Object.keys(options).length;
  const nGroupElements = Object.values(options).flat().length;
  const optionsHeight = Math.min(
    nGroups * GROUP_HEIGHT + nGroupElements * GROUP_ELEMENT_HEIGHT,
    MAX_HEIGHT
  );

  const hasSelectedElements = !isEmpty(formSelectedOptions);

  return (
    <div className={cx(className, styles.container)} ref={containerEl}>
      {label && <InputLabel text={label} />}
      <div className={styles.inputContainer}>
        <div
          className={cx(styles.input, {
            [styles.error]: error !== '',
            [styles.opened]: optionsOpened,
            [styles.placeholder]: placeholder !== ''
          })}
          onClick={openOptions}
          ref={inputEl}
        >
          {placeholder}
        </div>
        <div
          className={cx(styles.optionsContainer, {
            [styles.opened]: optionsOpened
          })}
          style={{ maxHeight: optionsOpened ? optionsHeight : 0 }}
        >
          {hasSelectedElements && !hideSelections && (
            <Selections
              selections={formSelectedOptions}
              onDeselect={onDeselect}
              onClear={onClear}
            />
          )}
          <div className={styles.options}>{optionList}</div>
        </div>
      </div>
      {!hideError && <InputError message={error} />}
    </div>
  );
}

export default GroupSelect;
