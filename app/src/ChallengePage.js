/* @flow */
"use strict";

import React from 'react';
import { Text, Image, View } from 'react-native';
import Subpage from './Subpage';

export default class ChallengePage extends React.Component {
  render() {
    let row = this.props.challenge;
    return (
      <Subpage backPress={this.props.backPress} title={row.title}>
        <View style={{flex:1}}>
            {(row.icon_url != "") ?
              <Image source={{uri: row.icon_url}} /> : null}
            <Text>{row.title}</Text>
            <Text>{row.description}</Text>
        </View>
      </Subpage>
    );
  }
}
