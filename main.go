package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		teamRegex := regexp.MustCompile("^\\s*\"Team\": \\[\\s*$")
		teamMatch := teamRegex.MatchString(scanner.Text())
		if teamMatch {
			parseTeams(scanner)
		}

		teamStepRegex := regexp.MustCompile("^\\s*\"TeamStep\": \\[\\s*$")
		teamStepMatch := teamStepRegex.MatchString(scanner.Text())
		if teamStepMatch {
			parseTeamSteps(scanner)
		}

		personRegex := regexp.MustCompile("^\\s*\"Person\": \\[\\s*$")
		personMatch := personRegex.MatchString(scanner.Text())
		if personMatch {
			parsePersons(scanner)
		}

		playerRegex := regexp.MustCompile("^\\s*\"Player\": \\[\\s*$")
		playerMatch := playerRegex.MatchString(scanner.Text())
		if playerMatch {
			parsePlayers(scanner)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("insert.sql")
	if err != nil {
		log.Fatal(err)
	}
	writeTeams(f)
	writeTeamSteps(f)
	writePersons(f)
	writePlayers(f)
	f.Close()
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
