import React from 'react';
import Filters, { Filter } from './Filters';
import { shallow } from 'enzyme';

describe('Filters', () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallow(<Filters filters={{ filter: 'value' }} />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });

  it('show right components', () => {
    expect(wrapper.exists('.title')).toBeTruthy();
    expect(wrapper.find(Filter).length).toBe(1);
  });

  it('can have multiple filters', () => {
    wrapper.setProps({
      filters: { filter1: 'value1', filter2: 'value2', filter3: 'value3' }
    });

    expect(wrapper.exists('.title')).toBeTruthy();
    expect(wrapper.find(Filter).length).toBe(3);
  });
});