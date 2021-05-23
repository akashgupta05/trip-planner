package controllers

import (
	"sort"

	"github.com/akashgupta05/trip-planner/common"
	"github.com/akashgupta05/trip-planner/models"
)

type TravelManager struct {
	ContinentsPath []string
	ProcessedData  *PreProcessor
}

func NewTravelManager(processedData *PreProcessor) *TravelManager {
	return &TravelManager{
		ContinentsPath: make([]string, 0),
		ProcessedData:  processedData,
	}
}

func (tm *TravelManager) FindContinentsPath() {
	currentContinentID := tm.ProcessedData.OriginCity.ContinentID
	loc0 := tm.ProcessedData.OriginCity.Location

	coordList := []models.ConinentCoordinate{}
	for key, loc := range tm.ProcessedData.ContinentsCoordinates {
		if key != currentContinentID {
			coordList = append(coordList,
				models.ConinentCoordinate{ContinentID: key, Location: loc})
		}
	}

	sort.SliceStable(coordList, func(i, j int) bool {
		switch orientation(loc0, coordList[i].Location, coordList[j].Location) {
		case 0:
			dist1 := common.DistanceBetweenCoordinates(loc0.Lat, loc0.Lon,
				coordList[i].Location.Lat, coordList[i].Location.Lon)
			dist2 := common.DistanceBetweenCoordinates(loc0.Lat, loc0.Lon,
				coordList[j].Location.Lat, coordList[j].Location.Lon)
			return dist1 < dist2
		case 1:
			return true
		default:
			return false
		}
	})

	for _, c := range coordList {
		tm.ContinentsPath = append(tm.ContinentsPath, c.ContinentID)
	}
}

func orientation(originLoc, loc1, loc2 models.Location) int {
	val := (loc1.Lon-originLoc.Lon)*(loc2.Lat-loc1.Lat) -
		(loc1.Lat-originLoc.Lat)*(loc2.Lon-loc1.Lon)
	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	}

	return 2
}
