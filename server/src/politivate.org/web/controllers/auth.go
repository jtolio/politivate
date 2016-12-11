package controllers

import (
	"politivate.org/web/controllers/auth"
)

func init() {
	Mux["auth"] = auth.Auth
}
