"use strict";

function causeKey(id) {
  return "cause-" + id;
}

function challengeKey(id, cause_id) {
  return "challenge-" + cause_id + "-" + id;
}

export default class Resources {
  constructor(auth_token, logout_cb) {
    this.auth_token = auth_token;
    this.logout_cb = logout_cb;

    this.partials = {};
    this.full = {};
    this.followers = {}

    this.request = this.request.bind(this);
    this.getCauses = this.getCauses.bind(this);
    this.getChallenges = this.getChallenges.bind(this);
    this.getCauseChallenges = this.getCauseChallenges.bind(this);
    this.getPartialCause = this.getPartialCause.bind(this);
    this.getPartialChallenge = this.getPartialChallenge.bind(this);
    this.getFullCause = this.getFullCause.bind(this);
    this.getFullChallenge = this.getFullChallenge.bind(this);
    this.forceFullCause = this.forceFullCause.bind(this);
    this.forceFullChallenge = this.forceFullChallenge.bind(this);
    this.followCause = this.followCause.bind(this);
    this.unfollowCause = this.unfollowCause.bind(this);
    this.causeFollowers = this.causeFollowers.bind(this);
    this.completeChallenge = this.completeChallenge.bind(this);
  }

  async request(method, resource, options) {
    let reqopts = {
      method,
      headers: {"X-Auth-Token": this.auth_token},
    };
    if (options && options.body) {
      // surely there's a better way.
      reqopts.body = Object.keys(options.body).map((key) => (
          encodeURIComponent(key) + "=" +
          encodeURIComponent(options.body[key]))
      ).join("&");
      reqopts.headers["Content-Type"] = "application/x-www-form-urlencoded";
    }
    let resp = await fetch(
        new Request("https://www.politivate.org/api" + resource, reqopts));
    if (!resp.ok) {
      if (resp.status == 401) {
        // TODO: this causes warnings cause it causes logic to happen on
        //  unmounted components.
        this.logout();
      }
      let json = null;
      try {
        json = await resp.json();
      } catch(err) {
        throw resp.statusText;
      }
      if (!json.err) {
        throw resp.statusText;
      }
      throw json.err;
    }
    let json = await resp.json();
    if (json.err) {
      throw json.err;
    }
    return json.resp;
  }

  async getCauses() {
    let items = await this.request("GET", "/v1/causes/");
    for (var item of items) {
      item._key = causeKey(item.id);
      this.partials[causeKey(item.id)] = item;
    }
    return items;
  }

  async getChallenges() {
    let items = await this.request("GET", "/v1/challenges/");
    for (var item of items) {
      item._key = challengeKey(item.id, item.cause_id);
      this.partials[challengeKey(item.id, item.cause_id)] = item;
    }
    return items;
  }

  async getCauseChallenges(id) {
    let items = await this.request("GET", "/v1/cause/" + id + "/challenges/");
    for (var item of items) {
      item._key = challengeKey(item.id, item.cause_id);
      this.partials[challengeKey(item.id, item.cause_id)] = item;
    }
    return items;
  }

  async getPartialCause(id) {
    let key = causeKey(id);
    if (this.partials.hasOwnProperty(key)) {
      return this.partials[key];
    }
    return (await this.getFullCause(id));
  }

  async getPartialChallenge(id, cause_id) {
    let key = challengeKey(id, cause_id);
    if (this.partials.hasOwnProperty(key)) {
      return this.partials[key];
    }
    return (await this.getFullChallenge(id, cause_id));
  }

  async getFullCause(id) {
    let key = causeKey(id);
    if (this.full.hasOwnProperty(key)) {
      return this.full[key];
    }
    return (await this.forceFullCause(id));
  }

  async getFullChallenge(id, cause_id) {
    let key = challengeKey(id, cause_id);
    if (this.full.hasOwnProperty(key)) {
      return this.full[key];
    }
    return (await this.forceFullChallenge(id, cause_id));
  }

  async forceFullCause(id) {
    let cause = await this.request("GET", "/v1/cause/" + id);
    cause._key = causeKey(id);
    this.full[cause._key] = cause;
    return cause;
  }

  async forceFullChallenge(id, cause_id) {
    let challenge = await this.request(
        "GET", "/v1/cause/" + cause_id + "/challenge/" + id);
    challenge._key = challengeKey(id, cause_id);
    this.full[challenge._key] = challenge;
    return challenge;
  }

  async followCause(id) {
    let key = causeKey(id);
    let result = await this.request(
        "POST", "/v1/cause/" + id + "/followers");
    this.followers[key] = result;
    return result;
  }

  async unfollowCause(id) {
    let key = causeKey(id);
    let result = await this.request(
        "DELETE", "/v1/cause/" + id + "/followers");
    this.followers[key] = result;
    return result;
  }

  async causeFollowers(id) {
    let key = causeKey(id);
    if (this.followers.hasOwnProperty(key)) {
      return this.followers[key];
    }
    let result = await this.request(
        "GET", "/v1/cause/" + id + "/followers");
    this.followers[key] = result;
    return result;
  }

  async completeChallenge(id, cause_id, body) {
    let resp = (await this.request("POST",
        "/v1/cause/" + cause_id + "/challenge/" + id + "/complete", {body}));
    return resp.actions;
  }
}
