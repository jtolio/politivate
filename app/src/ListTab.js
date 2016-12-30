"use strict";

import React from 'react';
import { ListView, RefreshControl } from 'react-native';
import { H2, View, Text, Card, CardItem, Thumbnail } from 'native-base';
import { styles, ErrorView } from './common';

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
      <View>
        <View style={styles.tabheader}>
          {this.props.header}
        </View>
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
