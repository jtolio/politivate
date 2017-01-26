"use strict";

import React, { Component } from 'react';
import { StyleSheet, View, Text, Button } from 'react-native';

var colors = {
  blue: "rgb(0, 117, 255)",
  red: "rgb(255, 66, 48)"
};

var styles = StyleSheet.create({
  appheader: {
    flexDirection: "row",
    padding: 5,
  },
  applogo: {
    resizeMode: "contain",
    flex: 1,
    width: null,
    height: 30,
    borderWidth: 0
  },
  tabheader: {
    padding: 10,
    fontWeight: "bold",
    fontSize: 20
  },
  tabBarText: {},
  tabBarUnderline: {
    backgroundColor: colors.red
  },
  tabBar: {
    borderWidth: 1,
    borderBottomWidth: 0,
    borderLeftWidth: 0,
    borderRightWidth: 0,
    borderColor: colors.blue
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
    return <View><Text>Error: {this.props.msg.toString()}</Text></View>;
  }
}

class Link extends Component {
  render() {
    return (<Text style={{color: colors.blue}} onPress={this.props.onPress}>
      {this.props.children}
    </Text>);
  }
}

class TabHeader extends Component {
  render() {
    return (
      <View>
        <Text style={styles.tabheader}>
          {this.props.children}
        </Text>
      </View>);
  }
}

class StyledButton extends Component {
  render() {
    return (
      <Button onPress={this.props.onPress} title={this.props.title}
              color={colors.blue} />
    );
  }
}

module.exports = {
  styles, LoadingView, ErrorView, Link, TabHeader, colors,
  Button: StyledButton
}
