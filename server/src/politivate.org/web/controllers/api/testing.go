package api

import (
	"net/http"
	"time"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whfatal"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/models"
)

func init() {
	mux["testing"] = whmux.Exact(http.HandlerFunc(serveTest))
}

func serveTest(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)

	switch r.FormValue("action") {
	case "filldb":
		causeSD := models.NewCause(ctx)
		causeSD.Name = "SaveDemocracy.org"
		causeSD.IconURL = "https://www.savedemocracy.org/wp-content/uploads/2016/12/cropped-save_democracy-kilrain-MED-1-192x192.png"
		causeSD.Save(ctx)

		chalSD := causeSD.NewChallenge(ctx)
		chalSD.Title = "Go to your state's capitol building!"
		chalSD.ShortDesc = "On December 19th, we will go to every state capitol " +
			"building, join the peaceful protest, and call on the electoral college " +
			"to do their job and prevent a demagogue from gaining power."
		var err error
		chalSD.Deadline, err = time.Parse(time.RFC822, "19 Dec 16 23:59 PST")
		if err != nil {
			whfatal.Error(err)
		}
		chalSD.IconURL = "https://cdn2.iconfinder.com/data/icons/the-urban-hustle-and-bustle/60/townhall-256.png"
		chalSD.Points = 100
		chalSD.Save(ctx)

		user := models.NewUser(ctx)
		user.Name = "Test User"
		user.Save(ctx)

	case "followtest":
		u := models.GetUsers(ctx)[0]
		c := models.GetCauses(ctx)[0]

		whjson.Render(w, r, map[string]interface{}{
			"user_is_following":   u.CauseIds(ctx),
			"cause_has_followers": c.UserCount(ctx),
		})
		return
	}

	whjson.Render(w, r, "ok")
}
