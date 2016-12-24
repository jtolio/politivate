"use strict";

import React from 'react';
import { Linking } from 'react-native';
import { Card, CardItem, Text, Thumbnail, Button, View } from 'native-base';
import Subpage from './Subpage';
import FollowButton from './FollowButton';
import { ErrorView, Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage backPress={this.props.backPress} title={row.name}>
        <Card>
          <CardItem header>
            {(this.props.cause.icon_url != "") ?
              <Thumbnail source={{uri: this.props.cause.icon_url}} /> : null}
            <Text>{this.props.cause.name}</Text>
            {this.props.followButton}
          </CardItem>
          <CardItem onPress={() => Linking
                .openURL(row.url).catch(err => this.setState({error: err}))}>
            <Link>
              {row.url}
            </Link>
          </CardItem>
          <CardItem>
            <Text>{row.description}</Text>
          </CardItem>
        </Card>
      </Subpage>
    );
  }
}
