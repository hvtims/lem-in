package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"tools/tools"
)

func main() {
	var shortsteps string
	var file string
	shortmouves := math.MaxInt
	if len(os.Args) != 2 {
		log.Fatal("invalid arguments!")
	}
	file = os.Args[1]
	ant, rooms, roomlinks, start, end := tools.Parsefile(file)
	graph := tools.Graph(rooms, roomlinks)

	paths := tools.FindAllPaths(graph, start, end)

	bublepaths := tools.Bublepath(paths)
	// fmt.Println(bublepaths)
	result := tools.Combinationpaths(bublepaths)
	for _, v := range result {
		antDistribution := tools.DistributeAnts(v, ant)
		// fmt.Println(antDistribution)
		finelresults, mouvecount := tools.SimulateAntMovement(v, antDistribution)
		if mouvecount < shortmouves {
			shortsteps = finelresults
			shortmouves = mouvecount
		}
	}
	if shortmouves < 1 {
		log.Fatal("ERROR: invalid data format")
	}
	File, _ := os.ReadFile(file)
	fmt.Println(string(File))
	fmt.Println()
	fmt.Print(shortsteps)
}
