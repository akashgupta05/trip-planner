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

	fmt.Println("Enter the city: ")
	city := ""
	fmt.Scanf("%s", &city)

	preprocessor.RemoveConinentForCity(city)

	travelManager := controllers.NewTravelManager(preprocessor)

	travelManager.FindContinentsPath()

}
