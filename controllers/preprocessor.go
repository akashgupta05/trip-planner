package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/akashgupta05/trip-planner/models"
)

type PreProcessor struct {
	Continents            map[string][]models.City
	FilePath              string
	cities                map[string]models.City
	ContinentsCoordinates map[string]models.Location
	citiesContinentMap    map[string]string
	OriginCity            models.City
}

func NewPreProcessor(filePath string) *PreProcessor {
	return &PreProcessor{
		FilePath:              filePath,
		Continents:            make(map[string][]models.City, 0),
		cities:                make(map[string]models.City),
		ContinentsCoordinates: make(map[string]models.Location),
		citiesContinentMap:    make(map[string]string),
	}
}

func (pp *PreProcessor) ReadData() error {
	citiesBytes, err := ioutil.ReadFile(pp.FilePath)
	if err != nil {
		fmt.Println("error while openning file", err)
		return err
	}

	cities := map[string]models.City{}
	json.Unmarshal(citiesBytes, &cities)
	pp.cities = cities
	return nil
}

func (pp *PreProcessor) PreProcessData() {
	pp.buildCitiesContinentsMap()
}

func (pp *PreProcessor) RemoveConinentForCity(city string) {
	for continentID, cities := range pp.Continents {
		if continentID == pp.citiesContinentMap[city] {
			pp.fillOriginCity(cities, city)
			delete(pp.Continents, continentID)
			return
		}
	}

}

func (pp *PreProcessor) fillOriginCity(cities []models.City, iata string) {
	for _, cityData := range cities {
		if cityData.Iata == iata {
			pp.OriginCity = cityData
		}
	}
}

func (pp *PreProcessor) buildCitiesContinentsMap() {
	continentCoordinatesLat := map[string]float64{}
	continentCoordinatesLon := map[string]float64{}
	continentCitiesCount := map[string]float64{}
	for _, city := range pp.cities {
		pp.citiesContinentMap[city.Iata] = city.ContinentID
		continentID := city.ContinentID
		if _, ok := pp.Continents[continentID]; ok {
			pp.Continents[continentID] = append(pp.Continents[continentID], city)
			continentCoordinatesLat[continentID] += city.Location.Lat
			continentCoordinatesLon[continentID] += city.Location.Lon
			continentCitiesCount[continentID]++
			continue
		}

		continentCoordinatesLat[continentID] = city.Location.Lat
		continentCoordinatesLon[continentID] = city.Location.Lon
		continentCitiesCount[continentID] = 1
		pp.Continents[continentID] = []models.City{city}
	}

	for key, val := range continentCitiesCount {
		pp.ContinentsCoordinates[key] = models.Location{
			Lat: continentCoordinatesLat[key] / val,
			Lon: continentCoordinatesLon[key] / val,
		}
	}
}
