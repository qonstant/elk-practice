package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"elk-practice/controller"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Info("Welcome Home Log")
		fmt.Fprintf(w, "Welcome Home")
	})

	router.GET("/api/base", controller.Base)

	return router
}
