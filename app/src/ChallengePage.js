"use strict";

import React from 'react';
import { Card, CardItem, Text, Thumbnail } from 'native-base';
import Subpage from './Subpage';

export default class ChallengePage extends React.Component {
  render() {
    let row = this.props.challenge;
    return (
      <Subpage backPress={this.props.backPress} title={row.title}>
        <Card style={{flex:1}}>
          <CardItem header>
            {(row.icon_url != "") ?
              <Thumbnail source={{uri: row.icon_url}} /> : null}
            <Text>{row.title}</Text>
          </CardItem>
          <CardItem>
            <Text>{row.description}</Text>
          </CardItem>
        </Card>
      </Subpage>
    );
  }
}
