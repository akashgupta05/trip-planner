package main

import (
	"fmt"
	"time"

	"github.com/akashgupta05/trip-planner/controllers"
)

var RESOURCE_PATH = "./resources/cities.json"

func main() {
	preprocessor := controllers.NewPreProcessor(RESOURCE_PATH)
	if err := preprocessor.ReadData(); err != nil {
		fmt.Println("error while reading json file", err)
		return
	}

	t := time.Now()
	preprocessor.PreProcessData()
	fmt.Println("preprocessing finished", time.Since(t))

	fmt.Println("Enter city iata")

	city := ""
	fmt.Scanln(&city)

	traveller := controllers.NewTraveller(preprocessor)
	traveller.StartTrip(city)

	for _, city := range traveller.VisitedOrder {
		fmt.Printf("%s (%s, %s) ->", city.ID, city.Name, city.ContinentID)
	}
	fmt.Println("\nDistance travelled:", traveller.MinimumDistance)
}
