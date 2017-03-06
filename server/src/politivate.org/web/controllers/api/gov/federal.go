package gov

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

func SunlightAPIReq(ctx context.Context, path string, vals map[string]string) (
	*http.Response, error) {
	return apiReq(ctx, "https://congress.api.sunlightfoundation.com"+path, vals)
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

type SunlightLegislator struct {
	Chamber   string `json:"chamber"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Title     string `json:"title"`

	Birthday    string `json:"birthday"`
	ContactForm string `json:"contact_form"`
	InOffice    bool   `json:"in_office"`
	District    int    `json:"district"`
	FacebookId  string `json:"facebook_id"`
	Gender      string `json:"gender"`
	MiddleName  string `json:"middle_name"`
	NameSuffix  string `json:"name_suffix"`
	NickName    string `json:"nickname"`
	Office      string `json:"office"`
	Party       string `json:"party"`
	SenateClass int    `json:"senate_class"`
	State       string `json:"state"`
	StateName   string `json:"state_name"`
	StateRank   string `json:"state_rank"`
	TermEnd     string `json:"term_end"`
	TermStart   string `json:"term_start"`
	TwitterId   string `json:"twitter_id"`
	Website     string `json:"website"`
	YoutubeId   string `json:"youtube_id"`
}

func (l *SunlightLegislator) Convert() *Legislator {
	return &Legislator{
		FullName: titlePrefix(l.Chamber) + l.FirstName + " " + l.LastName,
		Website:  l.Website,
		Offices: []Office{{
			Phone: l.Phone,
		}}}
}
