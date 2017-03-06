package gov

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

func OpenStatesAPIReq(ctx context.Context, path string, vals map[string]string) (
	*http.Response, error) {
	return apiReq(ctx, "https://openstates.org/api"+path, vals)
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
		FullName: titlePrefix(l.Chamber) + l.FullName,
		Website:  l.URL}
	for _, office := range l.Offices {
		rv.Offices = append(rv.Offices, Office{Phone: office.Phone})
	}
	return rv
}
