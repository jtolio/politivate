package controllers

import (
	"politivate.org/web/controllers/cause"
)

func init() {
	mux["cause"] = Beta(cause.Handler)
}
