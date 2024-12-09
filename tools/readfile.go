package tools

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Parsefile(file string) (int, []string, []string, string, string) {
	var part1offile []string
	var start string
	var end string
	var confile string
	var roomlinks []string
	var room []string
	fileopen, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if len(fileopen) == 0 {
		log.Fatal("Empty File")
	}
	confile = string(fileopen)
	elemnts := strings.Split(confile, "\n")
	ant, eroratoi := strconv.Atoi(elemnts[0])
	if eroratoi != nil || ant <= 0 {
		// iff negativ
		log.Fatal("ERROR: invalid data format")
	}
	// fix mor start end khwia//
	// fix mor start/end link //
	// fix 2-2 link / 2 - 77
	// fix ta room ma link end
	// fix cmnts
	// fix links
	// fix nfs room 2 mrat
	elemnts = elemnts[1:]
	for i := 0; i < len(elemnts); i++ {
		if elemnts[i] == "" {
			continue
		}
		if elemnts[i][0] == '#' && elemnts[i] != "##start" && elemnts[i] != "##end" {
			continue
		}
		if elemnts[i] == "##start" {
			if elemnts[i+1] == "" {
				log.Fatal("ERROR: invalid start format")
			}
			start = Splitt(elemnts[i+1])
		} else if elemnts[i] == "##end" {
			if elemnts[i+1] == "" {
				log.Fatal("ERROR: invalid end format")
			}
			end = Splitt(elemnts[i+1])
		}
	}
	for j := 0; j < len(elemnts); j++ {
		if elemnts[j] == "" {
			continue
		}
		if elemnts[j][0] == '#' && elemnts[j] != "##start" && elemnts[j] != "##end" {
			continue
		}
		if strings.Contains(elemnts[j], "-") {
			if elemnts[j] == "" {
				continue
			}
			if elemnts[j][0] == '#' {
				continue
			}
			for k := 0; k < j; k++ {
				if elemnts[k] == "" {
					continue
				}
				if elemnts[k][0] == '#' {
					continue
				}
				part1offile = append(part1offile, elemnts[k])
			}

			for k := j; k < len(elemnts); k++ {
				if elemnts[k] == "" {
					continue
				}
				if elemnts[k][0] == '#' {
					continue
				}
				roomlinks = append(roomlinks, elemnts[k])
			}
			break
		}
	}

	for i := 0; i < len(part1offile); i++ {
		if part1offile[i] == "" {
			continue
		}
		if part1offile[i][0] == '#' && part1offile[i] != "##start" && part1offile[i] != "##end" {
			continue
		}
		if part1offile[i] != "##start" && part1offile[i] != "##end" {
			room = append(room, Splitt(part1offile[i]))
		}
	}
	if len(start) > 1 || len(end) > 1 {
		log.Fatal("ERROR: invalid end or start")
	}
	return ant, room, roomlinks, start, end
}

func Splitt(line string) string {
	sliceline := strings.Split(line, " ")
	return sliceline[0]
}