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
	mux["federal"] = whmux.Dir{
		"district": whmux.Dir{
			"locate": whmux.Exact(http.HandlerFunc(federalDistrictLocate)),
			"reps":   whmux.Exact(http.HandlerFunc(repsByFederalDistrict)),
		}}
}

type FederalDistrict struct {
	State    string `json:"state"`
	District int    `json:"district"`
}

func FederalDistrictLocateByGPS(ctx context.Context,
	latitude, longitude float64) []FederalDistrict {
	resp, err := SunlightAPIReq(ctx, "/districts/locate", map[string]string{
		"latitude":  fmt.Sprint(latitude),
		"longitude": fmt.Sprint(longitude),
	})
	if err != nil {
		whfatal.Error(err)
	}
	defer resp.Body.Close()

	var districts struct {
		Results []FederalDistrict `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&districts)
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}
	return districts.Results
}

func federalDistrictLocate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	whjson.Render(w, r, FederalDistrictLocateByGPS(ctx, latitude, longitude))
}

func loadFederalLegislators(ctx context.Context, query map[string]string) (
	rv []*SunlightLegislator) {

	type legislatorResults struct {
		Count int `json:"count"`
		Page  struct {
			Count   int `json:"count"`
			PerPage int `json:"per_page"`
			Page    int `json:page"`
		}
		Results []*SunlightLegislator `json:"results"`
	}

	var vals = make(map[string]string, len(query))
	for key, val := range query {
		vals[key] = val
	}

	var legislators []*SunlightLegislator
	page := 0

	for page >= 0 {
		page += 1
		vals["page"] = fmt.Sprint(page)
		func() {
			resp, err := SunlightAPIReq(ctx, "/legislators", vals)
			if err != nil {
				whfatal.Error(err)
			}
			defer resp.Body.Close()

			result := legislatorResults{}
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				whfatal.Error(wherr.InternalServerError.Wrap(err))
			}

			legislators = append(legislators, result.Results...)

			if result.Page.Page*result.Page.PerPage >= result.Count {
				page = -1
			}
		}()
	}

	return legislators
}

func SenatorsByFederalDistrict(ctx context.Context,
	d FederalDistrict) []*SunlightLegislator {
	return loadFederalLegislators(ctx, map[string]string{
		"chamber": "senate",
		"state":   d.State,
	})
}

func HouseRepsByFederalDistrict(ctx context.Context,
	d FederalDistrict) []*SunlightLegislator {
	return loadFederalLegislators(ctx, map[string]string{
		"chamber":  "house",
		"state":    d.State,
		"district": fmt.Sprint(d.District),
	})
}

func repsByFederalDistrict(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)

	if r.FormValue("state") == "" || r.FormValue("district") == "" {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	district, err := strconv.Atoi(r.FormValue("district"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("bad argument"))
	}
	d := FederalDistrict{
		State:    r.FormValue("state"),
		District: district,
	}

	whjson.Render(w, r, append(
		HouseRepsByFederalDistrict(ctx, d),
		SenatorsByFederalDistrict(ctx, d)...))
}
