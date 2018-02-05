package main

import (
	mongo "dataApi/mongo"
	obj "dataApi/objects"
	"encoding/json"

	"github.com/pkg/errors"
)

func ConsumeSummonerInfo(body []byte) error {
	var summonerInfo obj.SummonerDTO
	err := json.Unmarshal(body, &summonerInfo)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerInfo")
	}

	err = mongo.UpdateSummoner(summonerInfo)
	if err != nil {
		return errors.Wrap(err, "Error in ConsumeSummonerInfo")
	}

	return nil
}
