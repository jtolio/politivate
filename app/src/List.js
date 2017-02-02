"use strict";

import React from 'react';
import { ListView, RefreshControl, View, ScrollView } from 'react-native';
import { ErrorView, colors } from './common';

export default class List extends React.Component {
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
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    if (this.props.children && !this.state.loading &&
        this.state.items.length == 0) {
      return (
        <ScrollView style={{flex: 1}} refreshControl={
            <RefreshControl onRefresh={this.update} refreshing={false}/>}>
          {this.props.children}
        </ScrollView>
      );
    }
    const ds = new ListView.DataSource({
      rowHasChanged: (r1, r2) => r1.id !== r2.id});
    let dataSource = ds.cloneWithRows(this.state.items);
    return (
        <ListView refreshControl={
            <RefreshControl refreshing={this.state.loading}
                            onRefresh={this.update}/>}
           enableEmptySections={true}
           dataSource={dataSource}
           renderRow={this.renderRow}
           renderSeparator={this.renderSeparator}/>
    );
  }
}
