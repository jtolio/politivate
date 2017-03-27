package forms

import (
	"net/http"

	"politivate.org/web/models"
)

func NewCauseForm() *Form {
	return &Form{
		Template: "cause_form",
		Form:     map[string]interface{}{},
	}
}

func EditCauseForm(c *models.Cause) *Form {
	return &Form{
		Template: "cause_form",
		Form: map[string]interface{}{
			"name":        c.Info.Name,
			"url":         c.Info.URL,
			"icon_url":    c.Info.IconURL,
			"description": c.Data.Description,
			"short_desc":  c.Data.ShortDescription,
		},
	}
}

func ProcessCauseForm(c *models.Cause, r *http.Request) (ok bool, f *Form) {
	c.Info.Name = r.FormValue("name")
	c.Info.URL = r.FormValue("url")
	c.Info.IconURL = r.FormValue("icon_url")
	c.Data.Description = r.FormValue("description")
	c.Data.ShortDescription = r.FormValue("short_desc")
	if c.Info.Name == "" || c.Info.URL == "" || c.Info.IconURL == "" ||
		c.Data.Description == "" {
		f := EditCauseForm(c)
		f.Error = "Required field missing"
		return false, f
	}
	return true, nil
}
