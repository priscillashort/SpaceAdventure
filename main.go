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

type SolarSystem struct {
	Name string `json:"name"`
    Planets []Planet `json:"planets"`
}

type Planet struct {
    Name   string `json:"name"`
    Description   string `json:"description"`
}

/*func parseJson(jsonFile var, planets var){
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, planets)
}*/

func initialPrompts(solarSystem *SolarSystem) {
	fmt.Println("Welcome to the " + (*solarSystem).Name + "!")
	fmt.Println("There are " + strconv.Itoa(len((*solarSystem).Planets)) + " planets to explore.")
	fmt.Println("What is your name?")
}

func readEchoName(reader *bufio.Reader) {
	//Read name
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	//Echo name
	fmt.Println("Nice to meet you,", name + ". My name is Eliza, I'm an old friend of Alexa.")
	fmt.Println("Let's go on an adventure!")
}

func validateResponse(reader *bufio.Reader, response *string){
	if *response == "N" || *response == "Y"{ 
		return
	}

	fmt.Println("Sorry, I didn't get that.")
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
    *response, _ = reader.ReadString('\n')
	*response = strings.Replace(*response, "\n", "", -1)
	validateResponse(reader, response)
}

func choosePlanet(reader *bufio.Reader, solarSystem *SolarSystem) (planet Planet){
	
	var planetName string

	//for loop to ensure entering correct planet name
	for {

		//Ask for planet name
		fmt.Println("Name the planet you would like to visit.")
		planetName, _ = reader.ReadString('\n')
		planetName = strings.Replace(planetName, "\n", "", -1)

		//Get planet description
		for i := 0; i < len((*solarSystem).Planets); i++ {
			if planetName == (*solarSystem).Planets[i].Name {
				planet = (*solarSystem).Planets[i]
			}
		}
	
		//Make sure description was found and entered correct planet name
		if len(planet.Description) > 0 {
			break
		} else {

			//Loop (ie. don't break) if correct planet if not entered
			fmt.Println("Sorry, " + planetName + " is not one of the planets in the " + (*solarSystem).Name + ".")
		
		}
	}

	return
}

func selectRandPlanet(solarSystem *SolarSystem) (planet Planet){
	//select random planet
	fmt.Println("Selecting a random planet...")
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	planet = (*solarSystem).Planets[rand.Intn(len((*solarSystem).Planets))]
	return
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
	var solarSystem SolarSystem
	//parseJson(&jsonFile, &planets)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &solarSystem)

	//Initial prompts
	initialPrompts(&solarSystem)

	//Create reader
	reader := bufio.NewReader(os.Stdin)

	//Read name and echo name
	readEchoName(reader)

	//Ask if selecting rand value
	fmt.Println("Shall I randomly choose a planet for you to visit? (Y or N)")
	response, _ := reader.ReadString('\n')
	response = strings.Replace(response, "\n", "", -1)
	validateResponse(reader, &response)

	//Declare planet variable
	var planet Planet

	//Select planet
	if response == "N"{
		planet = choosePlanet(reader, &solarSystem)
	} else {
		planet = selectRandPlanet(&solarSystem)
	}

	//Print name and description of selected planet
	fmt.Println("Traveling to " + planet.Name + "...")
	fmt.Println("Arived at " + planet.Name + ". " + planet.Description)
}
