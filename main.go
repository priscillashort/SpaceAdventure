package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"math/rand"
	"time"
)

type Planets struct {
	Name string `json:"name"`
    Planets []Planet `json:"planets"`
}

type Planet struct {
    Name   string `json:"name"`
    Description   string `json:"description"`
}

func main() {

	//Check for the correct number of arguements
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments.")
		fmt.Println("Expected 2 arguements. Recieved " + strconv.Itoa(len(os.Args)) + ".")
		return
	}

	// Open the json file and check for errors opening file
	jsonFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	//Close and parse json
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var planets Planets
	json.Unmarshal(byteValue, &planets)

	//Start prompts
	fmt.Println("Welcome to the " + planets.Name + "!")
	fmt.Println("There are " + strconv.Itoa(len(planets.Planets)) + " planets to explore.")
	fmt.Println("What is your name?")

	//Read name
	reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	//Echo name
	fmt.Println("Nice to meet you,", name + ". My name is Eliza, I'm an old friend of Alexa.")
	fmt.Println("Let's go on an adventure!")

	//Set planet variables
	var planet string
	var description string

	//main for loop to get Yes or No value
	for {

		//Ask if selecting rand value
		fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
    	yesno, _ := reader.ReadString('\n')
		yesno = strings.Replace(yesno, "\n", "", -1)

		if yesno == "N" {

			//for loop to ensure entering correct planet name
			for {

				//Ask for planet name
				fmt.Println("Name the planet you would like to visit.")
				planet, _ = reader.ReadString('\n')
				planet = strings.Replace(planet, "\n", "", -1)

				//Get planet description
				for i := 0; i < len(planets.Planets); i++ {
					if planet == planets.Planets[i].Name {
						description = planets.Planets[i].Description
					}
				}
			
				//Make sure description was found and entered correct planet name
				if len(description) > 0 {

					//print planet name and description
					fmt.Println("Traveling to " + planet + "...")
					fmt.Println("Arived at " + planet + ". " + description)
					break

				} else {

					//Loop (ie. don't break) if correct planet if not entered
					fmt.Println("Sorry, " + planet + " is not one of the planets in the " + planets.Name + ".")
				
				}

			}

			break

		} else if yesno == "Y" {

			//select random planet
			fmt.Println("Selecting a random planet...")
			rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
			var randPlanet Planet
			randPlanet = planets.Planets[rand.Intn(len(planets.Planets))]
			planet = randPlanet.Name
			description = randPlanet.Description

			//Print name and description of selected planet
			fmt.Println("Traveling to " + planet + "...")
			fmt.Println("Arived at " + planet + ". " + description)
			break

		} else {

			//Loop (ie. don't break) if Y or N was not entered
			fmt.Println("Sorry, I didn't get that.")

		}
	}
}
