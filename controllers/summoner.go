package controllers

import (
	mongo "dataApi/mongo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterSummonerHandlers(router *mux.Router, route string) {
	summoner := router.PathPrefix(route).Subrouter()
	summoner.HandleFunc("/byName/{summonerName}", GetSummonerInfo)
	summoner.HandleFunc("/byAccId/{accountId}", GetSummonerInfoAcc)
}

func GetSummonerInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	summonerName := vars["summonerName"]

	summonerInfo, err := mongo.GetSummonerInfo(summonerName)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetSummonerInfo: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	WriteResponse(w, 200, true, "", summonerInfo)
}

func GetSummonerInfoAcc(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["accountId"]

	intAccountId, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errLog := fmt.Sprintf("Error parsing string: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	summonerInfo, err := mongo.GetSummonerInfoAcc(intAccountId)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetSummonerInfo: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	WriteResponse(w, 200, true, "", summonerInfo)
}
