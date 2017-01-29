"use strict";

import React from 'react';
import { Text, Image, View, TouchableOpacity } from 'react-native';
import ListTab from './ListTab';
import CausePage from './CausePage';
import FollowButton from './FollowButton';

export default class CausesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    let followButton = (
      <FollowButton cause={row} appstate={this.props.appstate} />
    );
    return (
      <TouchableOpacity onPress={() => this.props.appstate
          .navigator.push({component: CausePage, passProps: {
              cause: row, followButton: followButton}})}>
        <View>
          {(row.icon_url != "") ?
            <Image source={{uri: row.icon_url}} /> : null}
          <Text>{row.name}</Text>
          {followButton}
        </View>
      </TouchableOpacity>
    );
  }

  render() {
    return (
      <ListTab resource="/v1/causes/" header="Causes"
        renderRow={this.renderRow} appstate={this.props.appstate} />
    );
  }
}
