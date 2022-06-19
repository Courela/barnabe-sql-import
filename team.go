package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type team struct {
	id         int
	name       string
	shortDescr string
}

var (
	teams []team

	teamIdRegex         *regexp.Regexp
	teamNameRegex       *regexp.Regexp
	teamShortDescrRegex *regexp.Regexp
)

func init() {
	teamIdRegex = regexp.MustCompile("^.*\"Id\": (.+),\\s*$")
	teamNameRegex = regexp.MustCompile("^.*\"Name\": \"(.+)\",\\s*$")
	teamShortDescrRegex = regexp.MustCompile("^.*\"ShortDescription\": \"(.+)\"\\s*$")
}

func parseTeams(scanner *bufio.Scanner) {
	log.Print("Teams started")
	for scanner.Scan() {
		id := teamIdRegex.FindStringSubmatch(scanner.Text())
		if len(id) > 0 {
			log.Printf("Team id = %s", id[1])
			scanner.Scan()

			endTeam, err := regexp.MatchString("^\\s*},\\s*$", scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			nameVal := ""
			shortDescrVal := ""
			for !endTeam {
				name := teamNameRegex.FindStringSubmatch(scanner.Text())
				if len(name) > 0 {
					log.Printf("    Team = %s", name[1])
					nameVal = name[1]
					scanner.Scan()
				}

				sd := teamShortDescrRegex.FindStringSubmatch(scanner.Text())
				if len(sd) > 0 {
					log.Printf("    ShortDescription = %s", sd[1])
					shortDescrVal = sd[1]
					scanner.Scan()
				}

				endTeam, err = regexp.MatchString("^\\s*},?\\s*$", scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
			}

			teams = append(teams, team{id: parseInt(id[1]), name: nameVal, shortDescr: shortDescrVal})
		}

		seasonStart, err := regexp.MatchString("Season", scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if seasonStart {
			break
		}
	}
	log.Print("Teams finished")
}

func writeTeams(f *os.File) {
	for _, team := range teams {
		f.WriteString(fmt.Sprintf("INSERT INTO team (Id, Name, ShortDescription) VALUES (%d, '%s', '%s');\n", team.id, team.name, team.shortDescr))
	}
	log.Printf("Processed %d teams", len(teams))
}
