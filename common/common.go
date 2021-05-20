package common

import (
	"math"

	"github.com/akashgupta05/trip-planner/models"
)

const radius float64 = 6371

func DistanceBetweenCities(city1, city2 *models.City) float64 {
	dLat := deg2rad(city2.Location.Lat - city1.Location.Lat)
	dLon := deg2rad(city2.Location.Lon - city1.Location.Lon)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(city1.Location.Lat))*math.Cos(deg2rad(city2.Location.Lat))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
