package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertRouteToArrayOfDirections(route string) []string {
	townsArray := strings.Split(route, "")
	directionsArray := []string{}
	// convert the route into an array of directions
	for index, town := range townsArray {
		if index != len(townsArray)-1 {
			directionsArray = append(directionsArray, fmt.Sprintf("%s%s", town, townsArray[index+1]))
		}
	}
	return directionsArray
}

// takes a route, checks if it exists and then calculates the value if it exists
func calculateRouteValue(route string, routeGraphValueMap map[string]int) string {
	directionsArray := convertRouteToArrayOfDirections(route)
	// check if the route exists in the graph
	for _, direction := range directionsArray {
		if _, ok := routeGraphValueMap[direction]; !ok {
			return fmt.Sprintf("NO SUCH ROUTE")
		}
	}
	// calculate the route value
	var routeValue int
	for _, direction := range directionsArray {
		routeValue += routeGraphValueMap[direction]
	}
	return fmt.Sprintf("%d", routeValue)
}

func calculateNoOfStopsInRoute(routes []string, stops int) string {
	totalStops := 0
	// now we can calculate the routes
	for _, route := range routes {
		routeNodes := strings.Split(route, "")
		if len(routeNodes)-1 <= stops {
			totalStops++
		}
	}
	return fmt.Sprintf("%d", totalStops)
}

func main() {

	routeGraph := []string{
		"AB5",
		"BC4",
		"CD8",
		"DC8",
		"DE6",
		"AD5",
		"CE2",
		"EB3",
		"AE7",
	}

	// take the above graph and synthesise it into a map that can be easily
	// looked up for specific route values
	routeGraphValueMap := make(map[string]int)

	for _, route := range routeGraph {
		splitRoute := strings.Split(route, "")
		routeValue, err := strconv.Atoi(splitRoute[2])
		if err != nil {
			panic(err)
		}
		routeGraphValueMap[splitRoute[0]+splitRoute[1]] = int(routeValue)
	}

	var output []string

	// 1. The distance of the route A-B-C.
	output = append(output, calculateRouteValue("ABC", routeGraphValueMap))

	// 2. The distance of the route A-D.
	output = append(output, calculateRouteValue("AD", routeGraphValueMap))

	// 3. The distance of the route A-D-C.
	output = append(output, calculateRouteValue("ADC", routeGraphValueMap))

	// 4. The distance of the route A-E-B-C-D.
	output = append(output, calculateRouteValue("AEBCD", routeGraphValueMap))

	// 5. The distance of the route A-E-D.
	output = append(output, calculateRouteValue("AED", routeGraphValueMap))

	// 6. The number of trips starting at C and ending at C with a maximum of 3
	// stops. In the sample data below, there are two such trips: C-D-C (2
	// stops). and C-E-B-C (3 stops).

	output6Sample := []string{
		"CDC",
		"CEBC",
		"CEBCDC",
		"CDCEBC",
		"CDEBC",
		"CEBCEBC",
		"CEBCEBCEBC",
	}

	output = append(output, calculateNoOfStopsInRoute(output6Sample, 3))

	// Print the output to the console
	for index, o := range output {
		fmt.Printf("Output #%d: %s\n", index+1, o)
	}
}
