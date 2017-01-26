"use strict";

import React from 'react';
import { View, TouchableOpacity, Text } from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';

export default class Subpage extends React.Component {
  render() {
    return (
      <View>
        <View>
          <TouchableOpacity onPress={this.props.backPress}>
            <Icon name='ios-arrow-back' />
          </TouchableOpacity>
          <Text>{this.props.title}</Text>
        </View>
        <View>
          {this.props.children}
        </View>
      </View>
    );
  }
}
