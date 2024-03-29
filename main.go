package main

import (
	mongo "dataApi/mongo"
	util "dataApi/util"
)

func main() {
	mongo.InitMongoRepo()

	forever := make(chan bool)
	go StartWeb()
	go util.RegisterConsumer("dataApi/summoner", ConsumeSummonerInfo)
	go util.RegisterConsumer("dataApi/summonerStats", ConsumeSummonerStats)
	go util.RegisterConsumer("dataApi/matchReferenceData", ConsumeMatchReference)
	go util.RegisterConsumer("dataApi/championData", ConsumeChampionData)
	<-forever
}
