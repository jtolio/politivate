"use strict";

import React from 'react';
import { ListView, RefreshControl, View } from 'react-native';
import { ErrorView, TabHeader } from './common';

export default class ListTab extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      items: [],
      error: null
    };
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true});
      let req = new Request(this.props.url,
          {headers: {"X-Auth-Token": this.props.appstate.authtoken}});
      let json = await (await fetch(req)).json();
      if (json.err) {
        this.setState({
          loading: false,
          error: json.err,
        });
        return;
      }
      this.setState({
        loading: false,
        items: json.resp,
      });
    } catch(error) {
      this.setState({
        loading: false,
        error: error,
      });
    }
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
             dataSource={dataSource} renderRow={this.props.renderRow}/>)}
      </View>
    );
  }
}
