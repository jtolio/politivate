"use strict";

import React from 'react';
import { H2, Text, Card, CardItem, Thumbnail } from 'native-base';
import ListTab from './ListTab';
import CausePage from './CausePage';
import CauseHeader from './CauseHeader';

export default class CausesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    return (
      <Card>
        <CauseHeader cause={row} button onPress={() => this
          .props.navigator.push({component: CausePage, passProps: {cause: row}})}/>
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
