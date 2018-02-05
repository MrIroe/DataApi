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

	<-forever
}
