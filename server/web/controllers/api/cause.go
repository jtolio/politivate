package api

import (
	"politivate.org/web/controllers/api/cause"
)

func init() {
	authedMux["cause"] = cause.Handler
}
