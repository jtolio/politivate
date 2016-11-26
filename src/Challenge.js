"use strict";

import React, { Component } from 'react';
import {
  H1, ListItem, List, View, Text, Card, Icon, CardItem } from 'native-base';
import { styles } from './common';

export default class Challenge extends Component {
  render() {
    return (
      <View>
      <Card>
        <CardItem>
          <Icon name={this.props.iconname}/>
          <Text>Challenge</Text>
        </CardItem>
      </Card>
      </View>
    );
  }
}
