package tools

import (
	"fmt"
	"strings"
)

func DistributeAnts(paths [][]string, nombreOfAnts int) [][]int {
	antDistribution := make([][]int, len(paths))
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path)
	}
	for i := 1; i <= nombreOfAnts; i++ {
		minIndex := 0
		for j := 1; j < len(paths); j++ {
			if len(antDistribution[j])+pathLengths[j] < len(antDistribution[minIndex])+pathLengths[minIndex] {
				minIndex = j
			}
		}

		antDistribution[minIndex] = append(antDistribution[minIndex], i)
	}
	return antDistribution
}

func SimulateAntMovement(paths [][]string, antDistribution [][]int) (string, int) {
	var finelresult string
	type AntPosition struct {
		ant  int
		path int
		step int
	}

	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}

	moveCount := 0
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			finelresult += strings.Join(moves, " ")
			finelresult += "\n"
		}
		antPositions = newPositions
		moveCount++
	}
	return finelresult, moveCount - 1
}
