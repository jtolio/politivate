import React from 'react';
import { AppRegistry } from 'react-native';
import AppRoot from './src/AppRoot';

export default class Politivate extends React.Component {
  render() { return <AppRoot/>; }
}

AppRegistry.registerComponent('politivate', () => Politivate);
