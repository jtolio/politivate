"use strict";

import React from 'react';
import { Card, CardItem, Text, Thumbnail } from 'native-base';
import Subpage from './Subpage';

export default class Cause extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage backPress={this.props.backPress} title={row.name}>
        <Card style={{flex:1}}>
          <CardItem header>
            {(row.icon_url != "") ?
              <Thumbnail source={{uri: row.icon_url}} /> : null}
            <Text>{row.name}</Text>
          </CardItem>
          <CardItem>
            <Text>Description</Text>
          </CardItem>
        </Card>
      </Subpage>
    );
  }
}
