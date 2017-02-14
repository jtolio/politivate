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
	mux["district"] = whmux.Dir{
		"locate": whmux.Exact(http.HandlerFunc(districtLocate)),
		"reps":   whmux.Exact(http.HandlerFunc(repsByDistrict)),
	}
}

func contains(vals map[string]string, name string) bool {
	_, exists := vals[name]
	return exists
}

type District struct {
	State    string `json:"state"`
	District int    `json:"district"`
}

func DistrictLocateByGPS(ctx context.Context,
	latitude, longitude float64) []District {
	resp, err := SunlightAPIReq(ctx, "/districts/locate", map[string]string{
		"latitude":  fmt.Sprint(latitude),
		"longitude": fmt.Sprint(longitude),
	})
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}
	defer resp.Body.Close()

	var districts struct {
		Results []District `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&districts)
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}
	return districts.Results
}

func districtLocate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	whjson.Render(w, r, DistrictLocateByGPS(ctx, latitude, longitude))
}

type Legislator struct {
	BioguideId     string   `json:"bioguide_id"`
	Birthday       string   `json:"birthday"`
	Chamber        string   `json:"chamber"`
	ContactForm    string   `json:"contact_form"`
	CRPId          string   `json:"crp_id"`
	District       int      `json:"district"`
	FacebookId     string   `json:"facebook_id"`
	Fax            string   `json:"fax"`
	FECIds         []string `json:"fec_ids"`
	FirstName      string   `json:"first_name"`
	Gender         string   `json:"gender"`
	GovTrackId     string   `json:"govtrack_id"`
	ICPSRId        int64    `json:"icpsr_id"`
	InOffice       bool     `json:"in_office"`
	LastName       string   `json:"last_name"`
	LeadershipRole string   `json:"leadership_role"`
	LISMemberId    string   `json:"lis_id"`
	MiddleName     string   `json:"middle_name"`
	NameSuffix     string   `json:"name_suffix"`
	NickName       string   `json:"nickname"`
	OCDId          string   `json:"ocd_id"`
	OCEmail        string   `json:"oc_email"`
	Office         string   `json:"office"`
	Party          string   `json:"party"`
	Phone          string   `json:"phone"`
	SenateClass    int      `json:"senate_class"`
	State          string   `json:"state"`
	StateName      string   `json:"state_name"`
	StateRank      string   `json:"state_rank"`
	TermEnd        string   `json:"term_end"`
	TermStart      string   `json:"term_start"`
	ThomasId       string   `json:"thomas_id"`
	Title          string   `json:"title"`
	TwitterId      string   `json:"twitter_id"`
	VotesmartId    int64    `json:"votesmart_id"`
	Website        string   `json:"website"`
	YoutubeId      string   `json:"youtube_id"`
}

func loadLegislators(ctx context.Context, query map[string]string) (
	rv []*Legislator) {

	type legislatorResults struct {
		Count int `json:"count"`
		Page  struct {
			Count   int `json:"count"`
			PerPage int `json:"per_page"`
			Page    int `json:page"`
		}
		Results []*Legislator `json:"results"`
	}

	var vals = make(map[string]string, len(query))
	for key, val := range query {
		vals[key] = val
	}

	var legislators []*Legislator
	page := 0

	for page >= 0 {
		page += 1
		vals["page"] = fmt.Sprint(page)
		func() {
			resp, err := SunlightAPIReq(ctx, "/legislators", vals)
			if err != nil {
				whfatal.Error(wherr.InternalServerError.Wrap(err))
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

func SenatorsByDistrict(ctx context.Context, d District) []*Legislator {
	return loadLegislators(ctx, map[string]string{
		"chamber": "senate",
		"state":   d.State,
	})
}

func HouseRepsByDistrict(ctx context.Context, d District) []*Legislator {
	return loadLegislators(ctx, map[string]string{
		"chamber":  "house",
		"state":    d.State,
		"district": fmt.Sprint(d.District),
	})
}

func repsByDistrict(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)

	if r.FormValue("state") == "" || r.FormValue("district") == "" {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	district, err := strconv.Atoi(r.FormValue("district"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("bad argument"))
	}
	d := District{
		State:    r.FormValue("state"),
		District: district,
	}

	whjson.Render(w, r, append(
		HouseRepsByDistrict(ctx, d),
		SenatorsByDistrict(ctx, d)...))
}
