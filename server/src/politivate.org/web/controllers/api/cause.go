package api

import (
	"politivate.org/web/controllers/api/cause"
)

func init() {
	mux["cause"] = cause.Handler
}
