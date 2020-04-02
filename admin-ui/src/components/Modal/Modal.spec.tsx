import React from 'react';
import Modal from './Modal';
import { testid } from '../../utils/testUtilsEnzyme';
import HorizontalBar from '../Layout/HorizontalBar/HorizontalBar';
import Button from '../Button/Button';
import { shallow } from 'enzyme';

const ACTION_BUTTON_LABEL = 'Action Button';

const mockOnAccept = jest.fn();
const mockOnClose = jest.fn();

describe('Modal', () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallow(
      <Modal
        title="Modal"
        message="Message"
        actionButtonLabel={ACTION_BUTTON_LABEL}
        onAccept={mockOnAccept}
        onClose={mockOnClose}
      />
    );
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });

  it('works with default props', () => {
    expect(wrapper.exists(HorizontalBar)).toBeTruthy();
    expect(wrapper.find('.title').text()).toBe('Modal');
    expect(wrapper.find('.message').text()).toBe('Message');
    expect(
      wrapper.find({ label: 'Action Button' }).exists(Button)
    ).toBeTruthy();
  });

  it('calls onClose when closing', () => {
    wrapper.find({ label: 'CANCEL' }).simulate('click');

    expect(mockOnClose).toHaveBeenCalledTimes(1);
  });

  it('hides when clicking on Cancel', () => {
    wrapper.setProps({ onClose: null });
    wrapper.find({ label: 'CANCEL' }).simulate('click');

    expect(
      wrapper.find(testid('modal-container')).exists('.visible')
    ).toBeFalsy();
  });

  it('shows shield when blocking', () => {
    wrapper.setProps({ blocking: true });

    expect(wrapper.exists('.bg')).toBeTruthy();
  });

  it('handles click event on Accept click', () => {
    wrapper.find({ label: 'Action Button' }).simulate('click');

    expect(mockOnAccept).toHaveBeenCalledTimes(1);
  });
});