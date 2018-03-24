package objects

type ChampionData struct {
	Data map[string]ChampionDTO `json:"data"`
}

type ChampionDTO struct {
	Title string `json:"title" bson:"Title"`
	Id    int    `json:"id" bson:"Id"`
	Key   string `json:"key" bson:"Key"`
	Name  string `json:"name" bson:"Name"`
}
