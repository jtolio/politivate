"use strict";

import React, { Component } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import lightTheme from 'native-base/Components/Themes/light';

var theme = lightTheme;

var styles = StyleSheet.create({
  appheader: {
    padding: 10,
    backgroundColor: theme.toolbarDefaultBg,
  },
  appheadertext: {
    color: theme.toolbarTextColor,
    fontSize: theme.fontSizeH1,
  },
  tabheader: {
    padding: 10,
  },
  tabBarText: {},
  tabBarUnderline: {
    backgroundColor: theme.topTabBarActiveTextColor,
  },
  tabBar: {
    borderWidth: 0,
  },
});

class LoadingView extends Component {
  constructor(props) {
    super(props);
    this.state = {counter: 0};
    this.tick = this.tick.bind(this);
  }

  componentDidMount() {
    this.timer = setInterval(this.tick, 300);
  }

  componentWillUnmount() {
    clearInterval(this.timer);
  }

  tick() {
    this.setState((prevState, props) => {
      return {counter: prevState.counter + 1};
    });
  }

  render() {
    let dots = "";
    for (let i = 0; i < this.state.counter % 4; i++) {
      dots += ".";
    }
    for (let i = dots.length; i < 4; i++) {
      dots += " ";
    }
    return <View alignItems="center"><Text>Loading{dots}</Text></View>;
  }
}

class ErrorView extends Component {
  render() {
    return <View><Text>Error: {this.props.msg}</Text></View>;
  }
}

module.exports = {
  "styles": styles,
  "LoadingView": LoadingView,
  "ErrorView": ErrorView,
  "theme": theme,
}
