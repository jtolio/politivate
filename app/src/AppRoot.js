"use strict";

import React, { Component } from 'react';
import { Navigator, BackAndroid, AsyncStorage } from 'react-native';
import Tabs from './Tabs';
import { LoadingView, ErrorView } from './common';

class BackHandler extends Component {
  constructor(props) {
    super(props);
    this.backPress = this.backPress.bind(this);
  }

  backPress() {
    if (this.props.route.hasOwnProperty("_isRoot")) {
      return false;
    }
    this.props.navigator.pop();
    return true;
  }

  componentDidMount() {
    BackAndroid.addEventListener('hardwareBackPress', this.backPress);
  }

  componentWillUnmount() {
    BackAndroid.removeEventListener('hardwareBackPress', this.backPress);
  }

  render() {
    return (<this.props.route.component navigator={this.props.navigator}
               appstate={this.props.appstate} backPress={this.backPress}
               {...this.props.route.passProps}/>);
  }
}

export default class AppRoot extends Component {
  constructor(props) {
    super(props)
    this.state = {
      loading: false,
      error: null
    };
    this.renderScene = this.renderScene.bind(this);
    this.logout = this.logout.bind(this);
  }

  renderScene(route, navigator) {
    return (<BackHandler route={route} navigator={navigator}
                         appstate={{logout: this.logout}} />)
  }

  logout() {
    this.setState({loading: true});
    AsyncStorage.removeItem("@v1/auth/token")
      .then(() => AsyncStorage.removeItem("@v2/auth/token"))
      .then(() => this.setState({loading: false}))
      .catch((err) => {
        this.setState({loading: false, error: err})
      });
  }

  render() {
    if (this.state.loading) {
      return <LoadingView/>;
    }
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    return (<Navigator initialRoute={{component: Tabs, _isRoot: true}}
                       renderScene={this.renderScene} />);
  }
}
