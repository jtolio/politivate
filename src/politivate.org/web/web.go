package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jtolds/webhelp"
	"github.com/jtolds/webhelp-oauth2"
	"github.com/jtolds/webhelp/sessions"
)

type Cause struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Challenge struct {
	Id         int    `json:"id"`
	Cause      Cause  `json:"cause"`
	Title      string `json:"title"`
	ShortDesc  string `json:"short_desc"`
	PostedTS   int64  `json:"posted_ts"`
	DeadlineTS *int64 `json:"deadline_ts,omitempty"`
	Icon       string `json:"icon"`
	Points     int    `json:"points"`
}

func maybeInt64(val int64) *int64 { return &val }

var (
	now = time.Now()

	SierraClub = Cause{
		Id:   1,
		Name: "Sierra Club",
		Icon: "http://66.media.tumblr.com/avatar_cdbb9208e450_128.png"}

	challengeData = []Challenge{
		{
			Id:         2,
			Cause:      SierraClub,
			Title:      "Call your local representative",
			ShortDesc:  "We need you to tell them how important the environment is!",
			PostedTS:   now.UnixNano(),
			DeadlineTS: nil,
			Icon:       "http://www.iconsdb.com/icons/preview/black/office-phone-xxl.png",
			Points:     10,
		},
		{
			Id:         3,
			Cause:      SierraClub,
			Title:      "Show up to town hall",
			ShortDesc:  "We need you to tell them how important the environment is!",
			PostedTS:   now.UnixNano(),
			DeadlineTS: maybeInt64(now.UnixNano() + (7 * 24 * 60 * 60 * 1000000000)),
			Icon:       "https://cdn2.iconfinder.com/data/icons/the-urban-hustle-and-bustle/60/townhall-256.png",
			Points:     100,
		},
	}
)

func challenges(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"response": challengeData,
	}, "", "  ")
	if err != nil {
		webhelp.FatalError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	webhelp.FatalError(err)
}

type LoginHandler struct {
	Group *oauth2.ProviderGroup
}

func (l *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h3>Login required</h3>`)
	fmt.Fprintf(w, "<p>Log in with:<ul>")
	for name, provider := range l.Group.Providers() {
		fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`,
			provider.LoginURL(r.FormValue("redirect_to"), false), name)
	}
	fmt.Fprintf(w, "</ul></p>")
}

type SettingsHandler struct {
	Group *oauth2.ProviderGroup
}

func (l *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h3>Logged in!</h3>`)
	fmt.Fprintf(w, "<p>Settings!</p><p><ul>")
	fmt.Fprintf(w, `<li><a href="%s">Logout</a></li>`,
		l.Group.LogoutAllURL("/"))
	fmt.Fprintf(w, "</ul></p>")
}

func init() {
	store := sessions.NewCookieStore(secret)
	group, err := oauth2.NewProviderGroup(
		"auth", "/auth", oauth2.RedirectURLs{},
		oauth2.Google(oauth2.Config{
			ClientID:     googleClientId,
			ClientSecret: googleClientSecret,
			Scopes:       []string{"profile", "email"},
			RedirectURL:  "https://www.politivate.org/auth/google/_cb"}),
		oauth2.Facebook(oauth2.Config{
			ClientID:     facebookClientId,
			ClientSecret: facebookClientSecret,
			RedirectURL:  "https://www.politivate.org/auth/facebook/_cb"}))
	if err != nil {
		panic(err)
	}

	http.Handle("/", webhelp.LoggingHandler(
		RequireHTTPS("www.politivate.org", sessions.HandlerWithStore(store,
			webhelp.DirMux{
				"challenges": http.HandlerFunc(challenges),
				"settings": group.LoginRequired(&SettingsHandler{Group: group},
					func(redirect_to string) string {
						return "/login?" + url.Values{
							"redirect_to": {redirect_to}}.Encode()
					}),
				"auth":  group,
				"login": &LoginHandler{Group: group},
				"legal": webhelp.DirMux{
					"tos":     http.HandlerFunc(tos),
					"privacy": http.HandlerFunc(privacy),
				},
				"": webhelp.RedirectHandler("/settings"),
			},
		))))
}

func tos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tos"))
}

func privacy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("privacy"))
}

func RequireHTTPS(host string, handler http.Handler) http.Handler {
	return webhelp.RouteHandlerFunc(handler,
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Scheme != "https" || r.URL.Host != host {
				u := *r.URL
				u.Scheme = "https"
				u.Host = host
				webhelp.Redirect(w, r, u.String())
			} else {
				handler.ServeHTTP(w, r)
			}
		})
}
