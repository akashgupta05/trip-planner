package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/akashgupta05/trip-planner/common"
	"github.com/akashgupta05/trip-planner/models"
)

type PreProcessor struct {
	Continents     map[string]map[string]*models.City
	CitiesMap      map[string]string
	CitiesDistance map[string]map[string]float64
	FilePath       string
	cities         map[string]*models.City
}

func NewPreProcessor(filePath string) *PreProcessor {
	return &PreProcessor{
		FilePath:       filePath,
		Continents:     make(map[string]map[string]*models.City),
		cities:         make(map[string]*models.City),
		CitiesMap:      make(map[string]string),
		CitiesDistance: make(map[string]map[string]float64),
	}
}

func (pp *PreProcessor) ReadData() error {
	citiesBytes, err := ioutil.ReadFile(pp.FilePath)
	if err != nil {
		fmt.Println("error while openning file", err)
		return err
	}

	cities := map[string]*models.City{}
	json.Unmarshal(citiesBytes, &cities)
	pp.cities = cities
	return nil
}

func (pp *PreProcessor) PreProcessData() {
	// a := []string{"BOM", "CCK", "GAM", "AAE", "LMP", "VDE", "BOM"}
	// distance := 0.0
	// for i, _ := range a {
	// 	if i > 0 {
	// 		distance += common.DistanceBetweenCities(pp.cities[a[i-1]], pp.cities[a[i]])
	// 	}
	// }
	// fmt.Println(distance)
	pp.buildCitiesContinentsMap()
	pp.buildCitiesDistanceMap()
}

func (pp *PreProcessor) buildCitiesContinentsMap() {
	for _, city := range pp.cities {
		continentID := city.ContinentID
		pp.CitiesMap[city.Iata] = continentID
		if _, ok := pp.Continents[continentID]; ok {
			pp.Continents[continentID][city.Iata] = city
			continue
		}

		pp.Continents[continentID] = map[string]*models.City{city.Iata: city}
	}
}

func (pp *PreProcessor) buildCitiesDistanceMap() {
	for _, cities := range pp.Continents {
		for _, city := range cities {
			pp.buildNeighbour(city)
		}
	}
}

func (pp *PreProcessor) buildNeighbour(city *models.City) {
	distanceFromDifferentCities := map[string][]*models.NeighBouringCity{}
	if _, ok := pp.CitiesDistance[city.Iata]; !ok {
		pp.CitiesDistance[city.Iata] = make(map[string]float64)
	}
	cityDistanceMap := pp.CitiesDistance[city.Iata]

	for continentID, cities := range pp.Continents {
		if continentID == city.ContinentID {
			continue
		}

		neighbouringCities := []*models.NeighBouringCity{}
		for _, diffCity := range cities {
			distance := common.DistanceBetweenCities(city, diffCity)
			cityDistanceMap[diffCity.Iata] = distance
			neighbouringCities = append(neighbouringCities, &models.NeighBouringCity{
				Name: diffCity.Name, Iata: diffCity.Iata, Continent: continentID, Distance: distance,
			})
		}
		sort.SliceStable(neighbouringCities, func(i, j int) bool {
			return neighbouringCities[i].Distance < neighbouringCities[j].Distance
		})
		distanceFromDifferentCities[continentID] = neighbouringCities[:10]
	}
	pp.CitiesDistance[city.Iata] = cityDistanceMap
	city.NeighbouringCities = distanceFromDifferentCities
}
