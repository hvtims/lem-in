package tools

import "strings"

func Graph(rooms []string, roomlinks []string) map[string][]string {
	graph := make(map[string][]string)
	for _, v := range rooms {
		for _, t := range roomlinks {
			slicpoint := strings.Split(t, "-")
			if Contains(slicpoint, v) {
				if strings.Contains(t, v) {
					nextroom := Getpoint(slicpoint, v) 
					if !Contains(graph[v], nextroom) {
						graph[v] = append(graph[v], nextroom)
					}
				}
			}
		}
	}
	return graph
}

func Getpoint(t []string, v string) string {
	var point string
	for _, c := range t {
		if c != v {
			point = c
		}
	}
	return point
}
