package objects

// Platform struct describes a Riot Platform
type SummonerDTO struct {
	// ID of the summoner icon associated with the summoner.
	ProfileIconId int `json:"profileIconId" bson:"ProfileIconId"`

	// Summoner name.
	Name string `json:"name" bson:"Name"`

	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel" bson:"SummonerLevel"`

	// Date summoner was last modified specified as epoch milliseconds.
	// The following events will update this timestamp: profile icon change,
	// playing the tutorial or advanced tutorial, finishing a game, summoner name change
	RevisionDate int64 `json:"revisionDate" bson:"RevisionDate"`

	// Summoner ID.
	Id int64 `json:"id" bson:"Id"`

	// Account ID.
	AccountId int64 `json:"accountId" bson:"AccountId"`
}

type SummonerMatchStats struct {
	AccountId  int64                  `json:"accountId" bson:"AccountId"`
	MatchId    int64                  `json:"matchId" bson:"MatchId"`
	ChampionId int                    `json:"championId" bson:"ChampionId"`
	Stats      ParticipantStatsDto    `json:"stats" bson:"Stats"`
	Timeline   ParticipantTimelineDto `json:"timeline" bson:"Timeline"`
}
