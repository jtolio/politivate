"use strict";

import React, { Component } from 'react';
import { Navigator, BackAndroid } from 'react-native';
import Tabs from './Tabs';

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
               backPress={this.backPress} {...this.props.route.passProps}/>);
  }
}

export default class AppRoot extends Component {
  constructor(props) {
    super(props)
    this.renderScene = this.renderScene.bind(this);
  }

  renderScene(route, navigator) {
    return (<BackHandler route={route} navigator={navigator}/>)
  }

  render() {
    return (<Navigator initialRoute={{component: Tabs, _isRoot: true}}
                       renderScene={this.renderScene} />);
  }
}
