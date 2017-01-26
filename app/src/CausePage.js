"use strict";

import React from 'react';
import { Linking, View, Text, Image, TouchableOpacity } from 'react-native';
import Subpage from './Subpage';
import { ErrorView, Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage backPress={this.props.backPress} title={row.name}>
        <View>
          {(this.props.cause.icon_url != "") ?
            <Image source={{uri: this.props.cause.icon_url}} /> : null}
          <Text>{this.props.cause.name}</Text>
          {this.props.followButton}

        <TouchableOpacity onPress={() => Linking.openURL(row.url).catch(err => {})}>
          <Link>
            {row.url}
          </Link>
        </TouchableOpacity>

          <Text>{row.description}</Text>

        </View>
      </Subpage>
    );
  }
}
