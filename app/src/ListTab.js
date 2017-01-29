"use strict";

import React from 'react';
import { ListView, RefreshControl, View } from 'react-native';
import { ErrorView, TabHeader, colors } from './common';

export default class ListTab extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      items: [],
      error: null
    };
    this.update = this.update.bind(this);
    this.renderSeparator = this.renderSeparator.bind(this);
    this.renderRow = this.renderRow.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true, error: null});
      let items = await this.props.appstate.request(
          "GET", this.props.resource);
      this.setState({loading: false, items});
    } catch(error) {
      this.setState({loading: false, error});
    }
  }

  renderSeparator(sectionId, rowId, adjacentRowHighlighted) {
    return (
      <View key={rowId} style={{
          borderBottomWidth: 1,
          borderColor: colors.primary.faint}}/>
    );
  }

  renderRow(rowData, sectionId, rowId, highlightRow) {
    return (
      <View style={{padding: 20, paddingTop: 5, paddingBottom: 5}}>
        {this.props.renderRow(rowData, sectionId, rowId, highlightRow)}
      </View>
    );
  }

  render() {
    const ds = new ListView.DataSource({
      rowHasChanged: (r1, r2) => r1.id !== r2.id});
    let dataSource = ds.cloneWithRows(this.state.items);
    return (
      <View style={{flex:1}}>
        <TabHeader>
          {this.props.header}
        </TabHeader>
        {this.state.error ?
         (<ErrorView msg={this.state.error}/>) :
         (<ListView refreshControl={
              <RefreshControl refreshing={this.state.loading}
                              onRefresh={this.update}/>}
             enableEmptySections={true}
             dataSource={dataSource}
             renderRow={this.renderRow}
             renderSeparator={this.renderSeparator}/>)}
      </View>
    );
  }
}
