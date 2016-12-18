"use strict";

import React from 'react';
import { H2, Text, Card, CardItem, Thumbnail } from 'native-base';
import ListTab from './ListTab';

export default class CausesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    return (
      <Card>
        <CardItem button header>
          {(row.icon_url != "") ?
            <Thumbnail source={{uri: row.icon_url}} /> : null}
          <Text>{row.name}</Text>
        </CardItem>
        <CardItem>
          <Text>Description</Text>
        </CardItem>
      </Card>
    );
  }

  render() {
    return (
      <ListTab url="https://www.politivate.org/api/v1/causes/"
        header={<H2>Causes</H2>} renderRow={this.renderRow} />
    );
  }
}
