package models

type City struct {
	ID                 string   `json:"id"`
	Iata               string   `json:"iata"`
	Location           Location `json:"location"`
	Name               string   `json:"name"`
	ContinentID        string   `json:"contId"`
	NeighbouringCities map[string][]NeighBouringCity
}

type ConinentCoordinate struct {
	ContinentID string
	Location
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type NeighBouringCity struct {
	Name      string
	Iata      string
	Distance  float64
	Continent string
}
