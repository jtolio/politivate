package api

import (
	"net/http"
	"time"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

func init() {
	mux["test"] = webhelp.Exact(http.HandlerFunc(serveTest))
}

func serveTest(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)

	switch r.FormValue("action") {
	case "filldb":
		causeSD := models.NewCause(ctx)
		causeSD.Name = "SaveDemocracy.org"
		causeSD.IconURL = "https://www.savedemocracy.org/wp-content/uploads/2016/12/cropped-save_democracy-kilrain-MED-1-192x192.png"
		err := causeSD.Save(ctx)
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
		}

		chalSD, err := causeSD.NewChallenge(ctx)
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
		}
		chalSD.Title = "Go to your state's capitol building!"
		chalSD.ShortDesc = "On December 19th, we will go to every state capitol " +
			"building, join the peaceful protest, and call on the electoral college " +
			"to do their job and prevent a demogogue from gaining power."
		chalSD.Deadline, err = time.Parse(time.RFC822, "19 Dec 16 23:59 PST")
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
		}
		chalSD.IconURL = "https://cdn2.iconfinder.com/data/icons/the-urban-hustle-and-bustle/60/townhall-256.png"
		chalSD.Points = 100
		err = chalSD.Save(ctx)
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
		}
	}

	webhelp.RenderJSON(w, r, "ok")
}
