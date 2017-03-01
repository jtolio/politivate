package gov

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["state"] = whmux.Dir{
		"reps": whmux.Exact(http.HandlerFunc(stateReps))}
}

func StateRepsLocateByGPS(ctx context.Context,
	latitude, longitude float64) []*OpenStatesLegislator {
	resp, err := OpenStatesAPIReq(ctx, "/v1/legislators/geo/", map[string]string{
		"lat":  fmt.Sprint(latitude),
		"long": fmt.Sprint(longitude),
	})
	if err != nil {
		whfatal.Error(err)
	}
	defer resp.Body.Close()

	var results []*OpenStatesLegislator
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}
	return results
}

func stateReps(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	whjson.Render(w, r, StateRepsLocateByGPS(ctx, latitude, longitude))
}
