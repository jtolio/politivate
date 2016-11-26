"use strict";

import React, { Component } from 'react';
import {
  H1, ListItem, List, View, Text, Card, Icon, CardItem } from 'native-base';
import { styles } from './common';
import Challenge from './Challenge';

export default class ChallengesTab extends Component {
  render() {
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.header}>
          <H1>Challenges</H1>
        </View>
        <Card>
          <CardItem button onPress={() => this.props.navigator.push({
                component: Challenge, name: "challenge",
                passProps: {iconname: "logo-googleplus"}})}>
            <Icon name="logo-googleplus" style={{color: "#dd5044"}}/>
            <Text>Google Plus</Text>
          </CardItem>
          <CardItem button onPress={() => this.props.navigator.push({
                component: Challenge, name: "challenge",
                passProps: {iconname: "logo-facebook"}})}>
            <Icon name="logo-facebook" style={{color: "#0000ff"}}/>
            <Text>Facebook</Text>
          </CardItem>
          <CardItem><Challenge iconname="logo-googleplus"/></CardItem>
        </Card>
      </View>
    );
  }
}
