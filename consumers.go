package main

import (
	mongo "dataApi/mongo"
	obj "dataApi/objects"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func ConsumeSummonerInfo(body []byte) error {
	var summonerInfo obj.SummonerDTO
	err := json.Unmarshal(body, &summonerInfo)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerInfo")
	}

	err = mongo.UpdateSummoner(&summonerInfo)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerInfo")
	}

	return nil
}

func ConsumeSummonerStats(body []byte) error {
	var summonerStats obj.SummonerMatchStats
	err := json.Unmarshal(body, &summonerStats)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerStats")
	}

	err = mongo.UpdateSummonerStats(&summonerStats)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerStats")
	}

	return nil
}

func ConsumeMatchReference(body []byte) error {
	var matchRefData obj.MatchReferenceData
	err := json.Unmarshal(body, &matchRefData)
	if err != nil {
		return errors.Wrap(err, "Error unmarshalling ConsumeMatchReference")
	}

	err = mongo.UpdateMatchReference(&matchRefData)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeMatchReference")
	}

	return nil
}

func ConsumeChampionData(body []byte) error {
	var championData obj.ChampionData
	err := json.Unmarshal(body, &championData)
	if err != nil {
		return errors.Wrap(err, "Error unmarshalling to champion data")
	}

	for key, value := range championData.Data {
		fmt.Printf("ChampionKey:%v", key)
		err = mongo.UpdateChampionData(value)
		if err != nil {
			fmt.Printf("Error updating championId:%v, championKey: %v, Original Err: %v", value.Id, key, err)
		}
	}

	// for i := range championData {
	// 	err = mongo.UpdateChampionData(championData[i])
	// 	if err != nil {
	// 		fmt.Printf("Error updating championId:%v Original Err: %v", championData[i].Id, err)
	// 	}
	// }

	return nil
}
