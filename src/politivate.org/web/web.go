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
	"golang.org/x/net/context"
	xoauth2 "golang.org/x/oauth2"
	"google.golang.org/appengine"
)

func Twitter(conf oauth2.Config) *oauth2.Provider {
	conf.Endpoint = xoauth2.Endpoint{
		AuthURL:  "https://api.twitter.com/oauth2/token",
		TokenURL: "https://api.twitter.com/oauth/request_token",
	}
	return &oauth2.Provider{
		Name:   "twitter",
		Config: xoauth2.Config(conf)}
}

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

func challenges(ctx context.Context,
	w webhelp.ResponseWriter, r *http.Request) error {
	data, err := json.MarshalIndent(map[string]interface{}{
		"response": challengeData,
	}, "", "  ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	return err
}

func settings(ctx context.Context,
	w webhelp.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("settings!"))
	return err
}

type LoginHandler struct {
	Group *oauth2.ProviderGroup
}

func (l *LoginHandler) HandleHTTP(ctx context.Context,
	w webhelp.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<h3>Login required</h3>`)
	fmt.Fprintf(w, "<p>Log in with:<ul>")
	for name, provider := range l.Group.Providers() {
		fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`,
			provider.LoginURL(r.FormValue("redirect_to"), false), name)
	}
	fmt.Fprintf(w, "</ul></p>")
	return nil
}

func init() {
	store := sessions.NewCookieStore(secret)
	group, err := oauth2.NewProviderGroup(
		"oauth", "/auth", oauth2.RedirectURLs{},
		oauth2.Google(oauth2.Config{
			ClientID:     googleClientId,
			ClientSecret: googleClientSecret}),
		Twitter(oauth2.Config{
			ClientID:     twitterClientId,
			ClientSecret: twitterClientSecret,
			RedirectURL:  "https://www.politivate.org/auth/twitter/_cb"}),
		oauth2.Facebook(oauth2.Config{
			ClientID:     facebookClientId,
			ClientSecret: facebookClientSecret,
			RedirectURL:  "https://www.politivate.org/auth/facebook/_cb"}))
	if err != nil {
		panic(err)
	}

	http.Handle("/", webhelp.Base{Root: webhelp.LoggingHandler(
		RequireHTTPS("www.politivate.org", sessions.HandlerWithStore(store,
			webhelp.DirMux{
				"challenges": webhelp.HandlerFunc(challenges),
				"settings": group.LoginRequired(webhelp.HandlerFunc(settings),
					func(redirect_to string) string {
						return "/login?" + url.Values{
							"redirect_to": {redirect_to}}.Encode()
					}),
				"auth":  group,
				"login": &LoginHandler{Group: group},
			},
		)))})
}

func RequireHTTPS(host string, handler webhelp.Handler) webhelp.Handler {
	return webhelp.HandlerFunc(func(ctx context.Context,
		w webhelp.ResponseWriter, r *http.Request) error {
		ctx = appengine.WithContext(ctx, r)
		if r.URL.Scheme != "https" || r.URL.Host != host {
			u := *r.URL
			u.Scheme = "https"
			u.Host = host
			return webhelp.Redirect(w, r, u.String())
		}
		return handler.HandleHTTP(ctx, w, r)
	})
}
