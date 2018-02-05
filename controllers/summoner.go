package controllers

import (
	mongo "dataApi/mongo"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterSummonerHandlers(router *mux.Router, route string) {
	summoner := router.PathPrefix(route).Subrouter()
	summoner.HandleFunc("/byName/{summonerName}", GetSummonerInfo)
}

func GetSummonerInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	summonerName := vars["summonerName"]

	summonerInfo, err := mongo.GetSummonerInfo(summonerName)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetSummonerInfo: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
	}

	WriteResponse(w, 200, true, "", summonerInfo)
}
