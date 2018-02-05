package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controllers "dataApi/controllers"
	util "dataApi/util"
)

func StartWeb() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(util.GetConfigValue("dataApiPort"), router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	controllers.RegisterSummonerHandlers(router, "/summoner")
	router.HandleFunc("/", handleInvalidPath)

	return router
}

func handleInvalidPath(w http.ResponseWriter, r *http.Request) {
	controllers.WriteResponse(w, http.StatusNotFound, false, "not found", nil)
}
