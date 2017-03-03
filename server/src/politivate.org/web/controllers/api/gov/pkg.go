package gov

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/spacemonkeygo/spacelog"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whmux"
)

var (
	Handler http.Handler = mux
	mux                  = whmux.Dir{}

	logger = spacelog.GetLogger()
)

func apiReq(ctx context.Context, uri string, vals map[string]string) (
	*http.Response, error) {
	query := url.Values{}
	for name, val := range vals {
		query.Add(name, val)
	}
	req, err := http.NewRequest("GET", uri+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := urlfetch.Client(ctx).Do(req)
	if err != nil {
		return nil, wherr.InternalServerError.Wrap(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, wherr.InternalServerError.New(
			"Invalid api status code returned from %s: %s", uri, resp.Status)
	}
	return resp, nil
}

func SunlightAPIReq(ctx context.Context, path string, vals map[string]string) (
	*http.Response, error) {
	return apiReq(ctx, "https://congress.api.sunlightfoundation.com"+path, vals)
}

func OpenStatesAPIReq(ctx context.Context, path string, vals map[string]string) (
	*http.Response, error) {
	return apiReq(ctx, "https://openstates.org/api"+path, vals)
}

func contains(vals map[string]string, name string) bool {
	_, exists := vals[name]
	return exists
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
		FullName: l.FirstName + " " + l.LastName,
		Website:  l.Website,
		Offices: []Office{{
			Phone: l.Phone,
		}}}
}

type OpenStatesOffice struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Type    string `json:"type"`
}

type OpenStatesRole struct {
	Chamber     string `json:"chamber"`
	District    string `json:"district"`
	EndDate     string `json:"end_date"`
	Party       string `json:"party"`
	StartDate   string `json:"start_date"`
	State       string `json:"state"`
	Committee   string `json:"committee"`
	CommitteeId string `json:"committee_id"`
	Position    string `json:"position"`
	Subcomittee string `json:"subcomittee"`
	Term        string `json:"term"`
	Type        string `json:"type"`
}

type OpenStatesLegislator struct {
	Active     bool                `json:"active"`
	BoundaryId string              `json:"boundary_id"`
	Chamber    string              `json:"chamber"`
	CreatedAt  string              `json:"created_at"`
	District   string              `json:"district"`
	FirstName  string              `json:"first_name"`
	FullName   string              `json:"full_name"`
	Id         string              `json:"id"`
	LastName   string              `json:"last_name"`
	LegId      string              `json:"leg_id"`
	MiddleName string              `json:"middle_name"`
	Offices    []*OpenStatesOffice `json:"offices"`
	Party      string              `json:"party"`
	PhotoURL   string              `json:"photo_url"`
	//	OldRoles   map[string][]*OpenStatesRole `json:"old_roles"`
	Roles     []*OpenStatesRole `json:"roles"`
	State     string            `json:"state"`
	Suffixes  string            `json:"suffixes"`
	UpdatedAt string            `json:"updated_at"`
	URL       string            `json:"url"`
}

func (l *OpenStatesLegislator) Convert() *Legislator {
	rv := &Legislator{
		FullName: l.FullName,
		Website:  l.URL}
	for _, office := range l.Offices {
		rv.Offices = append(rv.Offices, Office{Phone: office.Phone})
	}
	return rv
}

type Office struct {
	Phone string `json:"phone"`
}

type Legislator struct {
	FullName string   `json:"full_name"`
	Offices  []Office `json:"offices"`
	Website  string   `json:"website"`
}

func (l *Legislator) MarshalJSON() ([]byte, error) {
	var phone string
	if len(l.Offices) > 0 {
		phone = l.Offices[0].Phone
	}

	m := map[string]interface{}{
		"full_name": l.FullName,
		"offices":   l.Offices,
		"website":   l.Website,
		"comment": `the fields "chamber", "votesmart_id", "first_name",
                "last_name", and "phone" are all deprecated.`,
		"first_name":   "",
		"last_name":    "",
		"chamber":      "house",
		"phone":        phone,
		"votesmart_id": l.FullName + ": " + phone,
	}

	nameParts := strings.SplitN(l.FullName, " ", 2)
	if len(nameParts) > 0 {
		m["first_name"] = nameParts[0]
		if len(nameParts) > 1 {
			m["last_name"] = nameParts[1]
		}
	}

	return json.Marshal(m)
}

func LegislatorsByGPS(ctx context.Context, database string,
	latitude, longitude float64) (rv []*Legislator) {
	rv = make([]*Legislator, 0) // json not null
	var federalDistricts *[]FederalDistrict
	if database == "ushouse" || database == "us" || database == "usandstate" {
		fds := FederalDistrictLocateByGPS(ctx, latitude, longitude)
		federalDistricts = &fds
		for _, fd := range *federalDistricts {
			for _, l := range HouseRepsByFederalDistrict(ctx, fd) {
				rv = append(rv, l.Convert())
			}
		}
	}
	if database == "ussenate" || database == "us" || database == "usandstate" {
		if federalDistricts == nil {
			fds := FederalDistrictLocateByGPS(ctx, latitude, longitude)
			federalDistricts = &fds
		}
		for _, fd := range *federalDistricts {
			for _, l := range SenatorsByFederalDistrict(ctx, fd) {
				rv = append(rv, l.Convert())
			}
		}
	}
	var stateReps *[]*OpenStatesLegislator
	if database == "statehouse" || database == "state" || database == "usandstate" {
		srs := StateRepsLocateByGPS(ctx, latitude, longitude)
		stateReps = &srs
		for _, sr := range *stateReps {
			if sr.Chamber == "lower" {
				rv = append(rv, sr.Convert())
			}
		}
	}
	if database == "statesenate" || database == "state" || database == "usandstate" {
		if stateReps == nil {
			srs := StateRepsLocateByGPS(ctx, latitude, longitude)
			stateReps = &srs
		}
		for _, sr := range *stateReps {
			if sr.Chamber == "upper" {
				rv = append(rv, sr.Convert())
			}
		}
	}
	return rv
}
