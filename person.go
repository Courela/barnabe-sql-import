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

type person struct {
	id               uint
	name             string
	gender           string
	birthdate        string
	idCardNr         string
	idCardExpireDate string
	voterNr          string
	phone            string
	email            string
	localBorn        bool
	localTown        bool
	createdAt        string
	lastUpdatedAt    string
}

var (
	persons          []person
	idCardsProcessed []string

	personIdRegex               *regexp.Regexp
	personNameRegex             *regexp.Regexp
	personGenderRegex           *regexp.Regexp
	personBirthdateRegex        *regexp.Regexp
	personIdCardNrRegex         *regexp.Regexp
	personIdCardExpireDateRegex *regexp.Regexp
	personVoterNrRegex          *regexp.Regexp
	personPhoneRegex            *regexp.Regexp
	personEmailRegex            *regexp.Regexp
	personLocalBornRegex        *regexp.Regexp
	personLocalTownRegex        *regexp.Regexp
	personCreatedAtRegex        *regexp.Regexp
	personLastUpdatedAtRegex    *regexp.Regexp
)

func init() {
	personIdRegex = regexp.MustCompile("^.*\"Id\": (.+),\\s*$")
	personNameRegex = regexp.MustCompile("^.*\"Name\": \"(.*)\",\\s*$")
	personGenderRegex = regexp.MustCompile("^.*\"Gender\": (.+),\\s*$")
	personBirthdateRegex = regexp.MustCompile("^.*\"Birthdate\": (.+)\\s*$")
	personIdCardNrRegex = regexp.MustCompile("^.*\"IdCardNr\": \"(.*)\",\\s*$")
	personIdCardExpireDateRegex = regexp.MustCompile("^.*\"IdCardExpireDate\": (.+)\\s*$")
	personVoterNrRegex = regexp.MustCompile("^.*\"VoterNr\": (.+)\\s*$")
	personPhoneRegex = regexp.MustCompile("^.*\"Phone\": (.+)\\s*$")
	personEmailRegex = regexp.MustCompile("^.*\"Email\": (.+)\\s*$")
	personLocalBornRegex = regexp.MustCompile("^.*\"LocalBorn\": (.+)\\s*$")
	personLocalTownRegex = regexp.MustCompile("^.*\"LocalTown\": (.+)\\s*$")
	personCreatedAtRegex = regexp.MustCompile("^.*\"CreatedAt\": \"(.+)\".?\\s*$")
	personLastUpdatedAtRegex = regexp.MustCompile("^.*\"LastUpdatedAt\": \"(.+)\".?\\s*$")
}

func parsePersons(scanner *bufio.Scanner) {
	log.Print("Persons started")
	idCardNrSeq := 0
	for scanner.Scan() {
		id := personIdRegex.FindStringSubmatch(scanner.Text())
		if len(id) > 0 {
			log.Printf("Person id = %s", id[1])

			nameVal := ""
			genderVal := "NULL"
			birthdateVal := "1970-01-01T00:00:00"
			idCardNrVal := fmt.Sprintf("Empty%d", idCardNrSeq)
			idCardExpireDateVal := "NULL"
			voterNrVal := "NULL"
			phoneVal := "NULL"
			emailVal := "NULL"
			lb := false
			lt := false
			createdAtVal := ""
			lastUpdatedAtVal := "NULL"
			scanner.Scan()

			endPerson, err := regexp.MatchString("^\\s*},\\s*$", scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			for !endPerson {
				name := personNameRegex.FindStringSubmatch(scanner.Text())
				if len(name) > 0 {
					log.Printf("    Name = %s", name[1])
					nameVal = strings.Replace(name[1], "'", "\\'", -1)
					scanner.Scan()
				}

				gender := personGenderRegex.FindStringSubmatch(scanner.Text())
				if len(gender) > 0 {
					log.Printf("    Gender = %s", gender[1])
					if !strings.Contains(gender[1], "null") {
						genderVal = fmt.Sprintf("'%s'", strings.Trim(gender[1], ",\""))
					}
					scanner.Scan()
				}

				birthdate := personBirthdateRegex.FindStringSubmatch(scanner.Text())
				if len(birthdate) > 0 {
					log.Printf("    Birthdate = %s", birthdate[1])
					if !strings.Contains(birthdate[1], "null") {
						birthdateVal = strings.Trim(birthdate[1], ",\"")
					}
					scanner.Scan()
				}

				idCardNr := personIdCardNrRegex.FindStringSubmatch(scanner.Text())
				if len(idCardNr) > 0 {
					log.Printf("    IdCardNr = %s", idCardNr[1])
					if idCardNr[1] == "" {
						idCardsProcessed = append(idCardsProcessed, idCardNrVal)
						idCardNrSeq++
					} else if contains(idCardsProcessed, strings.ToUpper(idCardNr[1])) {
						c := "+"
						for contains(idCardsProcessed, fmt.Sprintf("%s%s", strings.ToUpper(idCardNr[1]), c)) {
							c += "+"
						}
						idCardNrVal = fmt.Sprintf("%s%s", strings.ToUpper(idCardNr[1]), c)
						idCardsProcessed = append(idCardsProcessed, strings.ToUpper(idCardNrVal))
					} else {
						idCardNrVal = strings.ToUpper(idCardNr[1])
						idCardsProcessed = append(idCardsProcessed, strings.ToUpper(idCardNrVal))
					}
					scanner.Scan()
				}

				idCardExpireDate := personIdCardExpireDateRegex.FindStringSubmatch(scanner.Text())
				if len(idCardExpireDate) > 0 {
					log.Printf("    IdCardExpireDate = %s", idCardExpireDate[1])
					if !strings.Contains(idCardExpireDate[1], "null") {
						idCardExpireDateVal = fmt.Sprintf("'%s'", strings.Trim(idCardExpireDate[1], ",\""))
					}
					scanner.Scan()
				}

				voterNr := personVoterNrRegex.FindStringSubmatch(scanner.Text())
				if len(voterNr) > 0 {
					log.Printf("    VoterNr = %s", voterNr[1])
					if !strings.Contains(voterNr[1], "null") {
						voterNrVal = fmt.Sprintf("'%s'", strings.Trim(voterNr[1], ",\""))
					}
					scanner.Scan()
				}

				phone := personPhoneRegex.FindStringSubmatch(scanner.Text())
				if len(phone) > 0 {
					log.Printf("    Phone = %s", phone[1])
					if !strings.Contains(phone[1], "null") {
						phoneVal = fmt.Sprintf("'%s'", strings.Trim(phone[1], ",\""))
					}
					scanner.Scan()
				}

				email := personEmailRegex.FindStringSubmatch(scanner.Text())
				if len(email) > 0 {
					log.Printf("    Email = %s", email[1])
					if !strings.Contains(email[1], "null") {
						phoneVal = fmt.Sprintf("'%s'", strings.Trim(email[1], ",\""))
					}
					scanner.Scan()
				}

				localBorn := personLocalBornRegex.FindStringSubmatch(scanner.Text())
				if len(localBorn) > 0 {
					log.Printf("    LocalBorn = %s", localBorn[1])
					lb, _ = strconv.ParseBool(strings.Trim(localBorn[1], ","))
					scanner.Scan()
				}

				localTown := personLocalTownRegex.FindStringSubmatch(scanner.Text())
				if len(localTown) > 0 {
					log.Printf("    LocalTown = %s", localTown[1])
					lt, _ = strconv.ParseBool(strings.Trim(localTown[1], ","))
					scanner.Scan()
				}

				createdAt := personCreatedAtRegex.FindStringSubmatch(scanner.Text())
				if len(createdAt) > 0 {
					log.Printf("    CreatedAt = %s", createdAt[1])
					createdAtVal = createdAt[1]
					scanner.Scan()
				}

				lastUpdatedAt := personLastUpdatedAtRegex.FindStringSubmatch(scanner.Text())
				if len(lastUpdatedAt) > 0 {
					log.Printf("    LastUpdatedAt = %s", lastUpdatedAt[1])
					lastUpdatedAtVal = fmt.Sprintf("'%s'", lastUpdatedAt[1])
					scanner.Scan()
				}

				log.Printf("Checking end of person: %s", scanner.Text())
				endPerson, err = regexp.MatchString("^\\s*}[,]?\\s*$", scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				//time.Sleep(100000000)
			}

			idVal := parseInt(id[1])
			persons = append(persons, person{
				id:               uint(idVal),
				name:             nameVal,
				gender:           genderVal,
				birthdate:        birthdateVal,
				idCardNr:         idCardNrVal,
				idCardExpireDate: idCardExpireDateVal,
				voterNr:          voterNrVal,
				phone:            phoneVal,
				email:            emailVal,
				localBorn:        lb,
				localTown:        lt,
				createdAt:        createdAtVal,
				lastUpdatedAt:    lastUpdatedAtVal,
			})
			log.Print("Appended 1 person")
			//time.Sleep(100000000)
		}

		playerStart, err := regexp.MatchString("Player", scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if playerStart {
			break
		}
	}
	log.Print("Persons finished")
}

func writePersons(f *os.File) {
	for _, person := range persons {
		f.WriteString(
			fmt.Sprintf("INSERT INTO person (Id, Name, Gender, Birthdate, IdCardNr, IdCardExpireDate, VoterNr, Phone, Email, LocalBorn, LocalTown, CreatedAt, LastUpdatedAt)"+
				" VALUES (%d, '%s', %s, '%s', '%s', %s, %s, %s, %s, %t, %t, '%s', %s);\n",
				person.id, person.name, person.gender, person.birthdate, person.idCardNr, person.idCardExpireDate, person.voterNr, person.phone, person.email, person.localBorn, person.localTown, person.createdAt, person.lastUpdatedAt))
	}
	log.Printf("Processed %d persons", len(persons))
}
