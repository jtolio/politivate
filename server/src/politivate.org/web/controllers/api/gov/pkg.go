package gov

import (
	"net/http"
	"net/url"

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

type Office struct {
	Phone string `json:"phone"`
}

type Legislator struct {
	FullName string   `json:"full_name"`
	Offices  []Office `json:"offices"`
	Website  string   `json:"website"`
}
