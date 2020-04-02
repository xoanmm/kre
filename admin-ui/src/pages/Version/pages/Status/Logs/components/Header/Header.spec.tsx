import React from 'react';
import Header from './Header';

import IconOpen from '@material-ui/icons/ExpandLess';
import IconClose from '@material-ui/icons/ExpandMore';
import IconStickBottom from '@material-ui/icons/VerticalAlignBottom';
import IconClear from '@material-ui/icons/DeleteOutline';
import { shallow, ShallowWrapper } from 'enzyme';

const mockWriteData = jest.fn();
jest.mock('@apollo/react-hooks', () => ({
  useApolloClient: () => ({
    writeData: mockWriteData
  })
}));

describe('Logs Header', () => {
  let wrapper: ShallowWrapper<any, Readonly<{}>, React.Component<{}, {}, any>>;
  const togglePanel = jest.fn();
  const toggleStickToBottomMock = jest.fn();

  beforeEach(() => {
    wrapper = shallow(
      <Header
        togglePanel={togglePanel}
        opened
        stickToBottom
        toggleStickToBottom={toggleStickToBottomMock}
      />
    );
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });

  it('show right components', () => {
    expect(wrapper.find('.buttons > div').length).toBe(3);
  });

  it('handles events', () => {
    wrapper
      .find(IconClose)
      .parent()
      .simulate('click');
    wrapper
      .find(IconStickBottom)
      .parent()
      .simulate('click');
    wrapper
      .find(IconClear)
      .parent()
      .simulate('click');

    wrapper.setProps({ opened: false });
    wrapper
      .find(IconOpen)
      .parent()
      .simulate('click');

    expect(togglePanel).toHaveBeenCalledTimes(2);
    expect(toggleStickToBottomMock).toHaveBeenCalledTimes(1);
    expect(mockWriteData).toHaveBeenCalledTimes(1);
  });
});