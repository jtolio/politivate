"use strict";

import React, { Component } from 'react';
import { ListView, Image, RefreshControl } from 'react-native';
import {
  Header, H2, ListItem, List, View, Text, Card, Icon, CardItem, Title,
  Thumbnail
} from 'native-base';
import { styles, LoadingView, ErrorView } from './common';
import Challenge from './Challenge';

export default class ChallengesTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      challenges: [],
      error: "",
    };
    this.renderRow = this.renderRow.bind(this);
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  renderRow(row) {
    return (
      <Card>
        <CardItem button header onPress={() => this.props.navigator.push({
              component: Challenge, passProps: {challenge: row}})}>
          {(row.icon_url != "") ?
            <Thumbnail source={{uri: row.icon_url}} /> : null}
          <Text>{row.title}</Text>
        </CardItem>
        <CardItem>
          <Text>{row.short_desc}</Text>
        </CardItem>
      </Card>
    );
  }

  update() {
    this.setState({loading: true});
    let req = new Request("https://www.politivate.org/api/v1/challenges/");
    req.headers.append("X-Auth-Token-Google", this.props.appstate.token);
    fetch(req)
      .then((response) => response.json())
      .then((json) => {
        console.log(json);
        this.setState({
          loading: false,
          challenges: json.resp,
        });
      })
      .catch((error) => {
        this.setState({
          loading: false,
          error: error,
        });
      });
  }

  render() {
    const ds = new ListView.DataSource({
      rowHasChanged: (r1, r2) => r1.id !== r2.id});
    let dataSource = ds.cloneWithRows(this.state.challenges);
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.tabheader}>
          <H2>Challenges</H2>
        </View>
        {this.state.error.length > 0 ?
         (<ErrorView msg={this.state.error}/>) :
         (<ListView refreshControl={
              <RefreshControl refreshing={this.state.loading}
                              onRefresh={this.update}/>}
             enableEmptySections={true}
             dataSource={dataSource} renderRow={this.renderRow}/>)}
      </View>
    );
  }
}
