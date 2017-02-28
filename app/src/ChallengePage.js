/* @flow */
"use strict";

import React from 'react';
import { Text, Image, View, ScrollView, Button } from 'react-native';
import Subpage from './Subpage';
import LoadablePage from './LoadablePage';
import { Link, ErrorView, LoadingView, phonecall } from './common';
import FollowButton from './FollowButton';
import MapView from 'react-native-maps';

const EARTH_RADIUS = 6378137;

function square(val) { return val * val; }

// haversine returns distance between two gps coordinates in
// meters, assuming a spherical earth.
function haversine(lat1, lon1, lat2, lon2) {
  let latr1 = lat1 * Math.PI / 180;
  let latr2 = lat2 * Math.PI / 180;
  let dlatr = (lat2 - lat1) * Math.PI / 180;
  let dlonr = (lon2 - lon1) * Math.PI / 180;

  let a = square(Math.sin(dlatr/2)) +
          Math.cos(latr1) * Math.cos(latr2) *
          square(Math.sin(dlonr/2));
  return EARTH_RADIUS * 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));
}

class ChallengeLocationMap extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      position: null,
      actions: null,
    };
    this.watch_id = null;
    this.updateSuccess = this.updateSuccess.bind(this);
    this.updateFailure = this.updateFailure.bind(this);
    this.completeChallenge = this.completeChallenge.bind(this);
  }

  componentDidMount() {
    if (this.watch_id !== null) { return; }
    this.watch_id = navigator.geolocation.watchPosition(
        this.updateSuccess, this.updateFailure, {
      enableHighAccuracy: true,
      maximumAge: 0,
      distanceFilter: 1,
    });
  }

  componentWillUnmount() {
    if (this.watch_id !== null) {
      navigator.geolocation.clearWatch(this.watch_id);
      this.watch_id = null;
    }
  }

  updateSuccess(position) {
    if (this.watch_id === null) { return; }
    // makes debugging harder but also cheating
    if (position.mocked) { return; }
    this.setState({position});
  }

  updateFailure(error) {
    if (this.watch_id === null) { return; }
    // TODO
    console.log(error);
  }

  async completeChallenge(latitude, longitude) {
    try {
      if (this.props.challenge.database != "direct") {
        // TODO
        console.log("invalid database type");
        return;
      }
      let body = {
        challenge_latitude: this.props.challenge.direct_latitude,
        challenge_longitude: this.props.challenge.direct_longitude,
        user_latitude: latitude,
        user_longitude: longitude,
      };
      let resp = await this.props.appstate.request("POST",
          "/v1/cause/" + this.props.challenge.cause_id + "/challenge/" +
          this.props.challenge.id + "/complete", {body});
      this.setState({actions: resp.actions});
    } catch(error) {
      // TODO
      console.log(error);
    }
  }

  render() {
    let chal = this.props.challenge;
    if (chal.database != "direct") {
      return <ErrorView msg="Unknown challenge type!"/>;
    }
    let completed = (chal.actions.length > 0);
    if (this.state.actions !== null) {
      completed = (this.state.actions.length > 0);
    }
    let pos = this.state.position;
    if (!pos) {
      return <LoadingView/>;
    }
    let distance = haversine(
        pos.coords.latitude, pos.coords.longitude,
        chal.direct_latitude, chal.direct_longitude);
    let in_range = distance <= chal.direct_radius;

    let latitudeDelta = Math.abs(chal.direct_latitude-pos.coords.latitude);
    let longitudeDelta = Math.abs(chal.direct_longitude-pos.coords.longitude);
    let initialRegion = {
      latitude: (chal.direct_latitude + pos.coords.latitude)/2 +
                latitudeDelta * 0.15,
      longitude: (chal.direct_longitude + pos.coords.longitude)/2,
      latitudeDelta: latitudeDelta * 1.5,
      longitudeDelta: longitudeDelta * 1.5,
    };

    let now = Date.now();
    let after_event_start = !(chal.event_start && now < chal.event_start);
    let before_event_end = !(chal.event_end && now > chal.event_end);

    var button;
    if (completed) {
      button = <Button disabled title="Completed" onPress={() => {}}/>;
    } else if (!after_event_start) {
      button = <Button disabled title="Event hasn't started" onPress={() => {}}/>;
    } else if (!before_event_end) {
      button = <Button disabled title="Event is over" onPress={() => {}}/>;
    } else if (!in_range) {
      button = <Button disabled title="Not in range" onPress={() => {}}/>;
    } else {
      button = (
        <Button title="Check in" onPress={() => {
          this.completeChallenge(pos.coords.latitude, pos.coords.longitude);
        }}/>
      );
    }

    return (
      <View>
        <MapView style={{height: 200}} showUserLocation={false}
            rotateEnabled={false} pitchEnabled={false} loadingEnabled={true}
            initialRegion={initialRegion}>
          <MapView.Marker coordinate={{
              latitude: chal.direct_latitude,
              longitude: chal.direct_longitude,
            }}/>
          <MapView.Marker coordinate={{
              latitude: pos.coords.latitude,
              longitude: pos.coords.longitude,
            }} image={require("../images/person.png")}/>
          <MapView.Circle radius={chal.direct_radius} center={{
              latitude: chal.direct_latitude,
              longitude: chal.direct_longitude,
            }} strokeWidth={0} fillColor="#ff000033" geodesic={true}/>
        </MapView>
        <View style={{paddingTop: 10}}/>
        {button}
      </View>
    );
  }
}

class ChallengeLocationAction extends React.Component {
  render() {
    let chal = this.props.challenge;
    if (chal.database != "direct") {
      return <ErrorView msg="Unknown challenge type!"/>;
    }
    return (
      <View>
        { chal.event_start ? (
          <View style={{flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>Start:</Text>
            <Text>{(new Date(chal.event_start)).toLocaleString()}</Text>
          </View>
        ) : null }
        { chal.event_end ? (
          <View style={{flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>End:</Text>
            <Text>{(new Date(chal.event_end)).toLocaleString()}</Text>
          </View>
        ) : null }
        <View style={{flexDirection: "row", paddingBottom: 20}}>
          <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>Address:</Text>
          <Text>{chal.direct_address.replace(", ", "\n").replace(", ", "\n")}</Text>
        </View>
        <ChallengeLocationMap challenge={chal} appstate={this.props.appstate}/>
      </View>
    );
  }
}

class ChallengePhonecallAction extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      actions: null,
    };
    this.complete = this.complete.bind(this);
  }

  async complete() {
    if (!await phonecall(this.props.who, this.props.phone)) {
      return
    }
    try {
      let resp = await this.props.appstate.request("POST",
          "/v1/cause/" + this.props.challenge.cause_id + "/challenge/" +
          this.props.challenge.id + "/complete",
          {body: {phone_number: this.props.phone}});
      this.setState({actions: resp.actions});
    } catch(err) {
      // TODO
      console.log(err);
    }
  }

  render() {
    let actions = this.props.challenge.actions;
    if (this.state.actions !== null) {
      actions = this.state.actions;
    }
    let completed = false;
    for (var action of actions) {
      if (action.phone == this.props.phone) {
        completed = true;
        break;
      }
    }
    let who = this.props.who;
    let phone = this.props.phone;
    let message = "Call " + phone;
    if (who) {
      message = who + ": " + message;
    }
    if (completed) {
      message += " (complete)";
    }
    return (
      <Button title={message} onPress={this.complete}/>
    );
  }
}

class ChallengePhonecallActions extends React.Component {
  render() {
    let chal = this.props.challenge;
    let results = [];

    let now = Date.now();
    let after_event_start = !(chal.event_start && now < chal.event_start);
    let before_event_end = !(chal.event_end && now > chal.event_end);

    if (!after_event_start) {
      results.push(
        <Button disabled title="Event hasn't started" onPress={() => {}}/>
      );
    } else if (!before_event_end) {
      results.push(
        <Button disabled title="Event is over" onPress={() => {}}/>
      );
    }

    if (results.length == 0) {
      if (chal.database == "direct") {
        results.push(
          <View style={{paddingTop: 10}} key="view-direct"/>
        );
        results.push(
          <ChallengePhonecallAction challenge={chal} phone={chal.direct_phone}
              key="button-direct" appstate={this.props.appstate}/>
        );
      } else {
        for (var legislator of chal.legislators) {
          results.push(
            <View style={{paddingTop: 10}}
                  key={"view-" + legislator.votesmart_id}/>
          );
          let title = {"senate": "Sen.", "house": "Rep."}[legislator.chamber];
          let name = legislator.first_name + " " + legislator.last_name;
          results.push(
            <ChallengePhonecallAction key={"button-" + legislator.votesmart_id}
                  who={title + " " + name} phone={legislator.phone}
                  appstate={this.props.appstate} challenge={chal}/>
          );
        }
      }
    }

    return (
      <View>
        {results}
      </View>
    );
  }
}


export default class ChallengePage extends React.Component {
  resourceURL() {
    return "/v1/cause/" + this.props.challenge.cause_id +
           "/challenge/" + this.props.challenge.id;
  }

  renderLoaded(chal) {
    return (
      <View style={{
          padding: 20
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 10}}>
          <Image
            source={{uri: this.props.cause.icon_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10, flex: 1}}>
            <Text style={{fontWeight: "bold"}}>{this.props.cause.name}</Text>
            <Link url={this.props.cause.url}>
              {this.props.cause.url}
            </Link>
          </View>
          <FollowButton cause={this.props.cause}
                appstate={this.props.appstate} />
        </View>
        <Text>{chal.description}</Text>
        <View style={{paddingTop: 10}}/>
        { chal.type == "phonecall" ?
            <ChallengePhonecallActions challenge={chal}
                appstate={this.props.appstate}/> :
          chal.type == "location" ?
            <ChallengeLocationAction challenge={chal}
                appstate={this.props.appstate}/> :
          null }
      </View>
    );
  }

  render() {
    return (
      <Subpage appstate={this.props.appstate}
               title={this.props.challenge.title}>
        <LoadablePage renderLoaded={this.renderLoaded.bind(this)}
                      resourceURL={this.resourceURL()}
                      appstate={this.props.appstate}/>
      </Subpage>
    );
  }
}
