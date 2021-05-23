package common

import (
	"math"
)

const radius float64 = 6371

func DistanceBetweenCoordinates(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
