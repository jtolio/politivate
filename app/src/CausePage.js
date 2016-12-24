"use strict";

import React from 'react';
import { Linking } from 'react-native';
import { Card, CardItem, Text, Thumbnail, Button, View } from 'native-base';
import Subpage from './Subpage';
import CauseHeader from './CauseHeader';
import { ErrorView, Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage backPress={this.props.backPress} title={row.name}>
        <Card>
          <CauseHeader cause={this.props.cause}/>
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
