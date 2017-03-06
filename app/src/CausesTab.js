"use strict";

import React from 'react';
import { Text, Image, View, TouchableOpacity } from 'react-native';
import List from './List';
import CausePage from './CausePage';
import { TabHeader } from './common';
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
        <View style={{flexDirection: "row",
                      alignItems: "center",
                      justifyContent: "space-between"}}>
          {(row.icon_url != "") ?
            <Image source={{uri: row.icon_url}}
                   style={{width: 50, height: 50, borderRadius: 10}} /> : null}
          <View style={{flex: 1, alignItems: "flex-start",
                        paddingLeft: 10, paddingRight: 10}}>
            <Text style={{fontWeight: "bold"}}>{row.name}</Text>
          </View>
          {followButton}
        </View>
      </TouchableOpacity>
    );
  }

  render() {
    return (
      <View style={{flex:1}}>
        <TabHeader>Causes</TabHeader>
        <List resource="/v1/causes/" renderRow={this.renderRow}
              appstate={this.props.appstate} keyFunc={(cause) => cause.id}/>
      </View>
    );
  }
}
