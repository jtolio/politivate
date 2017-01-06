package api

import (
	"politivate.org/web/controllers/api/gov"
)

func init() {
	unauthedMux["gov"] = gov.Handler
}
