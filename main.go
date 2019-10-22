package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
    "io/ioutil"
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
	// Open the jsonFile
	jsonFile, err := os.Open("planetarySystem.json")

	// if os.Open returns an error then print it
	if err != nil {
    	fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var planets Planets
	json.Unmarshal(byteValue, &planets)

	//Print a list of the available planets
	/*for i := 0; i < len(planets.Planets); i++ {
    	fmt.Println("Planet Name: " + planets.Planets[i].Name)
    	fmt.Println("Planet Description: " + planets.Planets[i].Description)
	}*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the " + planets.Name + "!")
	fmt.Println("There are 9 planets to explore.")
	fmt.Println("What is your name?")
    name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("Nice to meet you,", name + ". My name is Eliza, I'm an old friend of Alexa.")
	fmt.Println("Let's go on an adventure!")
	var planet string
	var description string

	for {
		fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
    	yesno, _ := reader.ReadString('\n')
		yesno = strings.Replace(yesno, "\n", "", -1)

		if yesno == "N" {

			for {
				fmt.Println("Name the planet you would like to visit.")
				planet, _ = reader.ReadString('\n')
				planet = strings.Replace(planet, "\n", "", -1)

				for i := 0; i < len(planets.Planets); i++ {
					if planet == planets.Planets[i].Name {
						description = planets.Planets[i].Description
					}
				}
			
				if len(description) > 0 {
					fmt.Println("Traveling to " + planet + "...")
					fmt.Println("Arived at " + planet + ". " + description)
					break
				} else {
			
					fmt.Println("Sorry, " + planet + " is not one of the planets in the solar system")
				}

			}
			break
		} else if yesno == "Y" {
			fmt.Println("Selecting a random planet...")
			planet = "Neptune"

			for i := 0; i < len(planets.Planets); i++ {
				if planet == planets.Planets[i].Name {
					description = planets.Planets[i].Description
				}
			}

			fmt.Println("Traveling to " + planet + "...")
			fmt.Println("Arived at " + planet + ". " + description)
			break
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}
