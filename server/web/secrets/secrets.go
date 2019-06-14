package secrets

import (
	"sync"

	"github.com/jtolds/goth"
	"github.com/jtolds/goth/providers/facebook"
	"github.com/jtolds/goth/providers/gplus"
	"github.com/jtolds/goth/providers/twitter"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

var (
	providerTypes = map[string]func(id, secret, url string) goth.Provider{
		"gplus": func(id, secret, url string) goth.Provider {
			return gplus.New(id, secret, url)
		},
		"facebook": func(id, secret, url string) goth.Provider {
			return facebook.New(id, secret, url)
		},
		"twitter": func(id, secret, url string) goth.Provider {
			return twitter.New(id, secret, url)
		},
	}
)

type secret_Provider struct {
	Provider     string
	ClientId     string
	ClientSecret string `datastore:",noindex"`
	CallbackURL  string `datastore:",noindex"`
}

func Providers(ctx context.Context) (rv []goth.Provider, err error) {
	var providers []*secret_Provider
	_, err = datastore.NewQuery("secret_Provider").GetAll(ctx, &providers)
	if err != nil {
		return nil, err
	}
	for _, provider := range providers {
		if fn, ok := providerTypes[provider.Provider]; ok {
			rv = append(rv, fn(provider.ClientId, provider.ClientSecret,
				provider.CallbackURL))
		}
	}
	if len(rv) == 0 {
		if appengine.IsDevAppServer() {
			return []goth.Provider{
				gplus.New(
					"776290794690-642buldhsscc4s5mnr5k24bj6damr9j6.apps.googleusercontent.com",
					"f7EhN_yEh2VF1YkWzqdWyvmT",
					"http://localhost:8080/auth/provider/gplus/callback"),
			}, nil
		}
	}
	return rv, nil
}

type secret_Setting struct {
	Name  string
	Value string `datastore:",noindex"`
}

var (
	secretSettingsMtx sync.Mutex
	secretSettings    = map[string]interface{}{}
)

func loadSecretSetting(ctx context.Context, name string) (string, error) {
	secretSettingsMtx.Lock()
	defer secretSettingsMtx.Unlock()
	if val, ok := secretSettings[name].(string); ok {
		return val, nil
	}
	var settings []*secret_Setting
	_, err := datastore.NewQuery("secret_Setting").Filter("Name =", name).
		GetAll(ctx, &settings)
	if err != nil {
		return "", err
	}
	val := ""
	if len(settings) > 0 {
		val = settings[0].Value
	}
	secretSettings[name] = val
	return val, nil
}

func CookieSecret(ctx context.Context) ([]byte, error) {
	val, err := loadSecretSetting(ctx, "cookie_secret")
	if err != nil {
		return nil, err
	}
	if val == "" {
		if appengine.IsDevAppServer() {
			return []byte("7dba11055c55257ffce1abdb97033839b6e57870"), nil
		}
		return nil, wherr.InternalServerError.New("no cookie secret configured")
	}
	return []byte(val), nil
}

type secret_WikiUser struct {
	Username   string
	BcryptPass string `datastore:",noindex"`
}

func WikiAuth(ctx context.Context, user, pass string) (valid bool) {
	var users []*secret_WikiUser
	_, err := datastore.NewQuery("secret_WikiUser").Filter("Username =", user).
		GetAll(ctx, &users)
	if err != nil {
		whfatal.Error(err)
	}
	return len(users) == 1 && nil == bcrypt.CompareHashAndPassword(
		[]byte(users[0].BcryptPass), []byte(pass))
}
