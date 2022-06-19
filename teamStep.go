package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type teamStep struct {
	teamId    uint8
	stepId    uint8
	season    uint16
	createdAt string
}

var (
	teamSteps []teamStep

	teamStepTeamIdRegex    *regexp.Regexp
	teamStepStepIdRegex    *regexp.Regexp
	teamStepSeasonRegex    *regexp.Regexp
	teamStepCreatedAtRegex *regexp.Regexp
)

func init() {
	teamStepTeamIdRegex = regexp.MustCompile("^.*\"TeamId\": (.+),\\s*$")
	teamStepStepIdRegex = regexp.MustCompile("^.*\"StepId\": (.+),\\s*$")
	teamStepSeasonRegex = regexp.MustCompile("^.*\"Season\": (.+),\\s*$")
	teamStepCreatedAtRegex = regexp.MustCompile("^.*\"CreatedAt\": \"(.+)\"\\s*$")
}

func parseTeamSteps(scanner *bufio.Scanner) {
	log.Print("TeamStep started")
	for scanner.Scan() {
		teamId := teamStepTeamIdRegex.FindStringSubmatch(scanner.Text())
		if len(teamId) > 0 {
			log.Printf("Step for Team = %s", teamId[1])
			scanner.Scan()

			stepId := teamStepStepIdRegex.FindStringSubmatch(scanner.Text())
			if len(stepId) > 0 {
				log.Printf("    StepId = %s", stepId[1])
				scanner.Scan()

				season := teamStepSeasonRegex.FindStringSubmatch(scanner.Text())
				if len(season) > 0 {
					log.Printf("    Season = %s", season[1])
					scanner.Scan()

					createdAt := teamStepCreatedAtRegex.FindStringSubmatch(scanner.Text())
					if len(createdAt) > 0 {
						log.Printf("    CreatedAt = %s", createdAt[1])

						teamSteps = append(teamSteps, teamStep{
							teamId:    uint8(parseInt(teamId[1])),
							stepId:    uint8(parseInt(stepId[1])),
							season:    uint16(parseInt(season[1])),
							createdAt: createdAt[1],
						})
					}
				}
			}
		}
		personStart, err := regexp.MatchString("Person", scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if personStart {
			break
		}
	}
	log.Print("TeamStep finished")
}

func writeTeamSteps(f *os.File) {
	for _, teamStep := range teamSteps {
		f.WriteString(
			fmt.Sprintf("INSERT INTO teamstep (TeamId, StepId, Season, CreatedAt) VALUES (%d, %d, %d, '%s');\n",
				teamStep.teamId, teamStep.stepId, teamStep.season, teamStep.createdAt))
	}
	log.Printf("Processed %d team steps", len(teamSteps))
}
