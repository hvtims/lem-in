package tools

func Combinationpaths(paths [][]string) [][][]string {
	var result [][][]string
	manupaths := paths
	for i := 0; i < len(manupaths); i++ {
		var combpaths [][]string
		for j := 0; j < len(manupaths); j++ {
			if !Checkpath(manupaths[i], manupaths[j]) {
				if !Checkallpaths(combpaths, manupaths[j]) {
					combpaths = append(combpaths, manupaths[j])
				}
			}
		}
		combpaths = append(combpaths, manupaths[i])
		result = append(result, combpaths)
	}
	return result
}

func Checkallpaths(combpath [][]string, manupstr []string) bool {
	for i := 0; i < len(combpath); i++ {
		if Checkpath(combpath[i], manupstr) {
			return true
		}
	}
	return false
}

func Checkpath(phatfixe []string, phathscompare []string) bool {
	for i := 1; i < len(phatfixe)-1; i++ {
		for j := 1; j < len(phathscompare)-1; j++ {
			if phatfixe[i] == phathscompare[j] {
				return true
			}
		}
	}
	return false
}


func Bublepath(paths [][]string) [][]string {
	n := len(paths)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if len(paths[j-1]) > len(paths[j]) {
				paths[j-1], paths[j] = paths[j], paths[j-1]
			}
		}
	}
	return paths
}

func FindAllPaths(graph map[string][]string, start string, end string) [][]string {
	var paths [][]string

	stack := []struct {
		room        string
		currentPath []string
	}{
		{start, []string{start}},
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		room := current.room
		currentPath := current.currentPath
		if room == end {
			pathCopy := currentPath
			paths = append(paths, pathCopy)
			continue
		}

		for _, neighbor := range graph[room] {
			if !Contains(currentPath, neighbor) {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				stack = append(stack, struct {
					room        string
					currentPath []string
				}{neighbor, newPath})
			}
		}
	}
	return paths
}

func Contains(slice []string, element string) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}
