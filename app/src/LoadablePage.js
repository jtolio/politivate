"use strict";

import React from 'react';
import { ScrollView, RefreshControl, View } from 'react-native';
import { ErrorView } from './common';

export default class LoadablePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      response: null,
      error: null
    };
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true, error: null});
      let response = await this.props.appstate.request(
          "GET", this.props.resourceURL);
      if (this.props.process) {
        response = await this.props.process(response);
      }
      this.setState({loading: false, response});
    } catch(error) {
      this.setState({loading: false, error});
    }
  }

  render() {
    return (
      <View style={{flex: 1}}>
        { this.props.renderHeader ? this.props.renderHeader() : null }
        <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.loading}
                            onRefresh={this.update}/>}>
            { this.state.loading ? null : (
              this.state.error ? <ErrorView msg={this.state.error}/> :
              this.props.renderLoaded(this.state.response)) }
        </ScrollView>
      </View>
    );
  }
}
