package controllers

import (
	"fmt"

	"github.com/akashgupta05/trip-planner/models"
)

type Traveller struct {
	OriginCity      *models.City
	MinimumDistance float64
	VisitedOrder    []models.City
	ProcessedData   *PreProcessor
}

func NewTraveller(ppData *PreProcessor) *Traveller {
	return &Traveller{
		ProcessedData:   ppData,
		MinimumDistance: 100000000,
		OriginCity:      &models.City{},
		VisitedOrder:    []models.City{},
	}
}

func getVisitName(city *models.City) string {
	return fmt.Sprintf("%s (%s ,%s)", city.Iata, city.Name, city.ContinentID)
}

func (t *Traveller) StartTrip(iata string) {
	t.OriginCity = t.ProcessedData.Continents[t.ProcessedData.CitiesMap[iata]][iata]
	otherContinents := []string{}
	for continentID, _ := range t.ProcessedData.Continents {
		if t.OriginCity.ContinentID != continentID {
			otherContinents = append(otherContinents, continentID)
		}
	}

	visitedOrder := []models.City{*t.OriginCity}

	t.travelNextCity(t.OriginCity, visitedOrder, otherContinents, 0)
	t.VisitedOrder = append(t.VisitedOrder, *t.OriginCity)

}

func (t *Traveller) travelNextCity(
	visitingCity *models.City, visitedOrder []models.City, otherContinents []string, distance float64) bool {
	flag := false
	if len(otherContinents) == 0 {
		distance += t.ProcessedData.CitiesDistance[visitingCity.Iata][t.OriginCity.Iata]
		if t.MinimumDistance > distance {
			if len(visitedOrder) == 6 {
				t.MinimumDistance = distance
				t.VisitedOrder = visitedOrder
			}
			return true
		}
		return false
	}

	remainingContinents := otherContinents
	for _, continentID := range remainingContinents {
		otherContinents = otherContinents[1:]

		nearestNeighbour := false
		for _, neighbour := range visitingCity.NeighbouringCities[continentID] {
			if distance+neighbour.Distance > t.MinimumDistance {
				return false
			}

			nextNeighbour := t.ProcessedData.Continents[continentID][neighbour.Iata]
			visitedOrder = append(visitedOrder, *nextNeighbour)

			nearestNeighbour = !t.travelNextCity(
				nextNeighbour, visitedOrder, otherContinents, distance+neighbour.Distance)
			newOrder := []models.City{}
			for _, visitedCity := range visitedOrder {
				if visitedCity.ID != nextNeighbour.ID {
					newOrder = append(newOrder, visitedCity)
				}
			}
			visitedOrder = newOrder
			if nearestNeighbour {
				break
			}
		}
		flag = nearestNeighbour
	}

	return flag
}
