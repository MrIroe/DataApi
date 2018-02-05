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
