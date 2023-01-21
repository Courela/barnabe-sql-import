package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type player struct {
	id            uint
	season        uint16
	teamId        uint8
	stepId        uint8
	personId      uint
	roleId        uint8
	resident      bool
	careTakerId   string
	comments      string
	photoFileName string
	docFilename   string
	createdAt     string
	lastUpdatedAt string
}

var (
	players []player

	playerIdRegex            *regexp.Regexp
	playerSeasonRegex        *regexp.Regexp
	playerTeamIdRegex        *regexp.Regexp
	playerStepIdRegex        *regexp.Regexp
	playerPersonIdRegex      *regexp.Regexp
	playerResidentRegex      *regexp.Regexp
	playerRoleIdRegex        *regexp.Regexp
	playerCareTakerIdRegex   *regexp.Regexp
	playerCommentsRegex      *regexp.Regexp
	playerPhotoFilenameRegex *regexp.Regexp
	playerDocFilenameRegex   *regexp.Regexp
	playerCreatedAtRegex     *regexp.Regexp
	playerLastUpdatedAtRegex *regexp.Regexp
)

func init() {
	playerIdRegex = regexp.MustCompile("^.*\"Id\": (.+),\\s*$")
	playerSeasonRegex = regexp.MustCompile("^.*\"Season\": (.+),\\s*$")
	playerTeamIdRegex = regexp.MustCompile("^.*\"TeamId\": (.+),\\s*$")
	playerStepIdRegex = regexp.MustCompile("^.*\"StepId\": (.+),\\s*$")
	playerPersonIdRegex = regexp.MustCompile("^.*\"PersonId\": (.+),\\s*$")
	playerResidentRegex = regexp.MustCompile("^.*\"Resident\": (.*),\\s*$")
	playerRoleIdRegex = regexp.MustCompile("^.*\"RoleId\": (.+),\\s*$")
	playerCareTakerIdRegex = regexp.MustCompile("^.*\"CareTakerId\": (.*),\\s*$")
	playerCommentsRegex = regexp.MustCompile("^.*\"Comments\": (.*).?\\s*$")
	playerPhotoFilenameRegex = regexp.MustCompile("^.*\"PhotoFilename\": (.*).?\\s*$")
	playerDocFilenameRegex = regexp.MustCompile("^.*\"DocFilename\": (.*).?\\s*$")
	playerCreatedAtRegex = regexp.MustCompile("^.*\"CreatedAt\": \"(.+)\".?\\s*$")
	playerLastUpdatedAtRegex = regexp.MustCompile("^.*\"LastUpdatedAt\": \"(.+)\".?\\s*$")
}

func parsePlayers(scanner *bufio.Scanner) {
	log.Print("Players started")
	for scanner.Scan() {
		id := playerIdRegex.FindStringSubmatch(scanner.Text())
		if len(id) > 0 {
			log.Printf("Player id = %s", id[1])
			scanner.Scan()

			endPlayer, err := regexp.MatchString("^\\s*},\\s*$", scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			seasonVal := 0
			teamIdVal := 0
			stepIdVal := 0
			personIdVal := 0
			residentVal := false
			roleIdVal := 0
			careTakerIdVal := "NULL"
			commentsVal := "NULL"
			photoFilenameVal := "NULL"
			docFilenameVal := "NULL"
			createdAtVal := ""
			lastUpdatedAtVal := "CURRENT_TIMESTAMP"

			for !endPlayer {
				season := playerSeasonRegex.FindStringSubmatch(scanner.Text())
				if len(season) > 0 {
					log.Printf("    Season = %s", season[1])
					seasonVal = parseInt(season[1])
					scanner.Scan()
				}

				teamId := playerTeamIdRegex.FindStringSubmatch(scanner.Text())
				if len(teamId) > 0 {
					log.Printf("    TeamId = %s", teamId[1])
					teamIdVal = parseInt(teamId[1])
					scanner.Scan()
				}

				stepId := playerStepIdRegex.FindStringSubmatch(scanner.Text())
				if len(stepId) > 0 {
					log.Printf("    StepId = %s", stepId[1])
					stepIdVal = parseInt(stepId[1])
					scanner.Scan()
				}

				personId := playerPersonIdRegex.FindStringSubmatch(scanner.Text())
				if len(personId) > 0 {
					log.Printf("    PersonId = %s", personId[1])
					personIdVal = parseInt(personId[1])
					scanner.Scan()
				}

				resident := playerResidentRegex.FindStringSubmatch(scanner.Text())
				if len(resident) > 0 {
					log.Printf("    Resident = %s", resident[1])
					residentVal, err = strconv.ParseBool(resident[1])
					if err != nil {
						log.Fatal(err)
					}
					scanner.Scan()
				}

				roleId := playerRoleIdRegex.FindStringSubmatch(scanner.Text())
				if len(roleId) > 0 {
					log.Printf("    RoleId = %s", roleId[1])
					roleIdVal = parseInt(roleId[1])
					scanner.Scan()
				}

				careTakerId := playerCareTakerIdRegex.FindStringSubmatch(scanner.Text())
				if len(careTakerId) > 0 {
					log.Printf("    CareTakerId = %s", careTakerId[1])
					if !strings.Contains(careTakerId[1], "null") {
						careTakerIdVal = strings.Trim(careTakerId[1], ",\"")
						if err != nil {
							log.Fatal(err)
						}
					}
					scanner.Scan()
				}

				comments := playerCommentsRegex.FindStringSubmatch(scanner.Text())
				if len(comments) > 0 {
					log.Printf("    Comments = %s", comments[1])
					if !strings.Contains(comments[1], "null") {
						commentsVal = fmt.Sprintf("'%s'", strings.Trim(comments[1], ","))
					}
					scanner.Scan()
				}

				photoFilename := playerPhotoFilenameRegex.FindStringSubmatch(scanner.Text())
				if len(photoFilename) > 0 {
					log.Printf("    PhotoFilename = %s", photoFilename[1])
					if !strings.Contains(photoFilename[1], "null") {
						photoFilenameVal = fmt.Sprintf("'%s'", strings.Trim(photoFilename[1], ",\""))
					}
					scanner.Scan()
				}

				docFilename := playerDocFilenameRegex.FindStringSubmatch(scanner.Text())
				if len(docFilename) > 0 {
					log.Printf("    DocFilename = %s", docFilename[1])
					if !strings.Contains(docFilename[1], "null") {
						docFilenameVal = fmt.Sprintf("'%s'", strings.Trim(docFilename[1], ",\""))
					}
					scanner.Scan()
				}

				createdAt := playerCreatedAtRegex.FindStringSubmatch(scanner.Text())
				if len(createdAt) > 0 {
					log.Printf("    CreatedAt = %s", createdAt[1])
					if !strings.Contains(createdAt[1], "null") {
						createdAtVal = fmt.Sprintf("'%s'", strings.Split(strings.Trim(createdAt[1], ",\""), ".")[0])
					}
					scanner.Scan()
				}

				lastUpdatedAt := playerLastUpdatedAtRegex.FindStringSubmatch(scanner.Text())
				if len(lastUpdatedAt) > 0 {
					log.Printf("    LastUpdatedAt = %s", lastUpdatedAt[1])
					if !strings.Contains(lastUpdatedAt[1], "null") {
						lastUpdatedAtVal = fmt.Sprintf("'%s'", strings.Split(strings.Trim(lastUpdatedAt[1], ",\""), ".")[0])
					}
					scanner.Scan()
				}

				endPlayer, err = regexp.MatchString("^\\s*}.?\\s*$", scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
			}

			idVal := parseInt(id[1])
			players = append(players, player{
				id:            uint(idVal),
				season:        uint16(seasonVal),
				teamId:        uint8(teamIdVal),
				stepId:        uint8(stepIdVal),
				personId:      uint(personIdVal),
				resident:      residentVal,
				roleId:        uint8(roleIdVal),
				careTakerId:   careTakerIdVal,
				comments:      commentsVal,
				photoFileName: photoFilenameVal,
				docFilename:   docFilenameVal,
				createdAt:     createdAtVal,
				lastUpdatedAt: lastUpdatedAtVal,
			})
		}
	}
	log.Print("Players finished")
}

func writePlayers(f *os.File) {
	dateFormat := "%Y-%c-%dT%T"
	for _, player := range players {
		f.WriteString(
			fmt.Sprintf("INSERT INTO player (Id, Season, TeamId, StepId, PersonId, Resident, RoleId, CareTakerId, Comments, PhotoFilename, DocFilename, CreatedAt, LastUpdatedAt) \n"+
				"VALUES (%d, %d, %d, %d, %d, %t, %d, %s, %s, %s, %s, STR_TO_DATE(%s,'%s'), %s);\n",
				player.id, player.season, player.teamId, player.stepId, player.personId, player.resident, player.roleId, player.careTakerId, player.comments, player.photoFileName, player.docFilename, player.createdAt, dateFormat, player.lastUpdatedAt))
	}
	log.Printf("Processed %d players", len(players))
}
