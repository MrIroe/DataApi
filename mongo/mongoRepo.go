package mongo

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	obj "dataApi/objects"
	util "riotApi/util"
)

var primarySess *mgo.Session

const entityDbName = "Entities"
const summonerColName = "Summoners"
const summonerStatsColName = "SummonerStats"
const matchRefColName = "MatchReferences"

func InitMongoRepo() {
	var err error
	primarySess, err = mgo.Dial(util.GetConfigValue("mongoAddress"))
	if err != nil {
		fmt.Println(errors.Wrap(err, "Couldnt connect to mongo"))
		os.Exit(-1)
	}

	primarySess.SetMode(mgo.Monotonic, true)
}

func UpdateSummoner(summoner *obj.SummonerDTO) error {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(summonerColName)

	_, err := c.Upsert(bson.M{"AccountId": summoner.AccountId}, &summoner)
	if err != nil {
		return errors.Wrap(err, "Error updating summoner:"+summoner.Name)
	}

	return nil
}

func UpdateSummonerStats(summonerStats *obj.SummonerMatchStats) error {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(summonerStatsColName)

	_, err := c.Upsert(bson.M{"$and": []bson.M{bson.M{"AccountId": summonerStats.AccountId}, bson.M{"MatchId": summonerStats.MatchId}}}, &summonerStats)
	if err != nil {
		return errors.Wrap(err, "Error updating summoner stats")
	}

	return nil
}

func UpdateMatchReference(matchReference *obj.MatchReferenceData) error {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(matchRefColName)

	_, err := c.Upsert(bson.M{"$and": []bson.M{bson.M{"AccountId": matchReference.AccountId}, bson.M{"MatchId": matchReference.MatchId}}}, &matchReference)
	if err != nil {
		return errors.Wrap(err, "Error updating match reference")
	}

	return nil
}

func GetSummonerInfo(summonerName string) (*obj.SummonerDTO, error) {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(summonerColName)

	var summonerInfo obj.SummonerDTO
	err := c.Find(bson.M{"Name": summonerName}).One(&summonerInfo)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting SummonerInfo")
	}

	return &summonerInfo, nil
}

func GetSummonerInfoAcc(accountId int64) ([]obj.SummonerDTO, error) {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(summonerColName)

	var summonerInfo []obj.SummonerDTO
	err := c.Find(bson.M{"AccountId": accountId}).All(&summonerInfo)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting SummonerInfo")
	}

	return summonerInfo, nil
}
