"use strict";

import React, { Component } from 'react';
import { H2, ListItem, List, View } from 'native-base';
import { styles, LoadingView, ErrorView } from './common';

export default class ProfileTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      profile: null,
      error: null
    };
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  update() {
    this.setState({loading: true});
    let req = new Request("https://www.politivate.org/api/v1/profile");
    fetch(req)
      .then((response) => response.json())
      .then((json) => {
        this.setState({
          loading: false,
          profile: json.resp,
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
    if (this.state.loading) {
      return <LoadingView/>;
    }
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>
    }
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.tabheader}>
          <H2>{this.state.profile.name}</H2>
        </View>
      </View>
    );
  }
}
