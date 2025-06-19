package routers

import (
	"mediadb/utils"
	"net/http"
)

type DocumentationRouter struct {
	BaseRoute string
	Documentation string
}

func (self *DocumentationRouter) getDocumentation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(self.Documentation))
}

func (self *DocumentationRouter) GetHandler() http.Handler {
	ctx := http.NewServeMux()
	ctx.HandleFunc("GET " + utils.ConcatUrls(self.BaseRoute, "/", true), self.getDocumentation)
	return ctx
}
