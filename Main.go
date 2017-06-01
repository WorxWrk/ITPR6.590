package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"strings"
	"strconv"

)

// --------------------------------------------- Type Creation
// Creates new data structures

	type City struct{
		path1 string
		exit1 string
		cityName string
		path2 string
		exit2 string
	}

	type Driver struct{

// Used to hold the current driver. Changes depending on the loop number
		driverNumber int

// Used to keep a driver in a loop
		key bool

// Used to hold the Number given by the Random number generator
		newNumber int

// Used to track the amount of times the driver has visted "Akina"
		akinaCount int

// Holds the number of the current city, within the "allCities" array.
//Used to set the "currentCity" variable
		currentCityNumber int

// Used to hold an entire City Section
		currentCity City

// Stores the name of the City the driver is in before the driver moves to the next city
		currentCityName string

// Holds the string value of the path that has been selected
		pathName string

// Holds the string value of the exit that has been selected
		exitName string

// Holds name of the city that has been selected next
		nextCityName string

// An Array that holds all of the output the driver instance has creted
		totalOutput []string
		}

	// --------------------------------------------- Function Creation
	// Creates the functions

func startingCity(driver Driver, allCities [4]City) Driver{

	driver.currentCityNumber = driver.newNumber
	driver.currentCity = allCities[driver.currentCityNumber]
	driver.currentCityName = driver.currentCity.cityName

	return driver
}

func count(driver Driver) Driver{

	if driver.currentCityName == "Akina" {
			driver.akinaCount = driver.akinaCount + 1
	}

	return driver
}

func driveCity(driver Driver, allCities [4]City) Driver {

	driver.currentCityName = driver.currentCity.cityName

	if driver.newNumber == 1 { // right side path, int goes up
		driver.pathName = driver.currentCity.path2
		driver.exitName = driver.currentCity.exit2
		if driver.currentCityNumber == 3 { // restricts int size to keep it inside the allCities array
			driver.currentCityNumber = 0
		} else {
			driver.currentCityNumber = driver.currentCityNumber + 1
		}
	}

	if driver.newNumber == 0{ // left side path, int goes down
		driver.pathName = driver.currentCity.path1
		driver.exitName = driver.currentCity.exit1
		if driver.currentCityNumber == 0 { // restricts int size to keep it inside the allCities array
			driver.currentCityNumber = 3
		} else {
			driver.currentCityNumber = driver.currentCityNumber - 1
		}
	}

	driver.currentCity = allCities[driver.currentCityNumber]
	driver.nextCityName = driver.currentCity.cityName

// Return the driver type with updated values
	return driver
}

func driveExit(driver Driver) Driver{

	if driver.newNumber == 2 { //  number that allows the driver to exit
		driver.key = true
		driver.nextCityName = "Outside City"
	}
// Return the driver type with updated values
	return driver
}

func extras(driver Driver) Driver{

	var strDriverNumber string
	strDriverNumber = strconv.Itoa(driver.driverNumber)

	var strAkinaCount string
	strAkinaCount = strconv.Itoa(driver.akinaCount)

	mergedStr := strings.Join([]string{"Driver", strDriverNumber, "Met With Jhon Jamieson", strAkinaCount, "Times!"}, " ")
	driver.totalOutput = append(driver.totalOutput, mergedStr)

	if driver.akinaCount != 1 | 2 {

		if driver.akinaCount == 0{
			driver.totalOutput = append(driver.totalOutput, "That Driver Missed Out!")
		}

		if driver.akinaCount >= 3{
			driver.totalOutput = append(driver.totalOutput, "That Driver Needed Lots Of Help!")
		}
	}
	fmt.Println("-----")
	return driver
}

func driveCityPrint (driver Driver) Driver{

	var strDriverNumber string
	strDriverNumber = strconv.Itoa(driver.driverNumber)

	mergedStr := strings.Join([]string{
		"Driver",
		strDriverNumber,
		"heading from",
		driver.currentCityName,
		"to",
		driver.nextCityName,
		"via",
		driver.pathName },
		 " ")
	driver.totalOutput = append(driver.totalOutput, mergedStr)

	return driver
}

func main() {

// --------------------------------------------- Start/Varible Initialization
// City Creation - Creates 4 varablies of the new type City.
// During Comments, a single one of these cities might be refered to as City Sections or Quadrants

// City 1
	City1 := City {"Path 1","Exit 1","City1","Path2","Exit 2"}

//City 2
	City2 := City {"Path 1","Exit 1","City2","Path2","Exit 2"}

//City 3
	City3 := City {"Path 1","Exit 1","Akina","Path2","Exit 2"}

//City 4
	City4 := City {"Path 1","Exit 1","City4","Path2","Exit 2"}

// Array Creation - Create a single array to hold the previous 4 Cities.
// Array for holding all city sections/quadrants.
	var allCities [4] City
	allCities[0] = City1
	allCities[1] = City2
	allCities[2] = City3
	allCities[3] = City4

// Varible Creation
// Variables will be initialized later, during the loops

//Holds the conveted input from the user
	var seed int

// --------------------------------------------- Program Begins
// User Input And RNG/Seed Initialization

 	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a number:");
	scanner.Scan()
	input := scanner.Text()
	seed, err := strconv.Atoi(input)

		if err != nil{
			fmt.Println("Please provide an input of the Integer type")
			// some form of panic
		}

	rand.Seed(int64(seed))

// Loop to complete multiiple driving instances
	for index:=0; index < 5; index++{

// Varible re-Initialization to start a driving instance

// Varable that holds all the output of the system
			var totalOutput []string

// Varable for the new type "Driver".
			driver := Driver { index + 1, false, 0, 0, 0, City1, "No City", "No Path", "No Exit", "No City", totalOutput}

// Finding a Starting City for the driving instance

			driver.newNumber = rand.Intn(4)
			driver = startingCity(driver, allCities);

// Loop for the moving sections of a single driving instance
			for driver.key == false{

// Path Decision. What path the driver will take out of the current city.

				driver.newNumber = rand.Intn(2)	//Sets a new rnd Variable
				driver = driveCity(driver, allCities);

//Exit Decision. If the Driver will use an exit or continue to the city.
				// set RND variable
				driver.newNumber = rand.Intn(6)

				// Decides if the driver will exit or continue
				driver = driveExit(driver);

				// Stores the string for this loop (within a driving instance)
				driver = driveCityPrint(driver);
				driver = count(driver)
			}

// Additional Dialouge. The exta lines required at the end of a driving instance.

				driver = extras(driver);

// Print the all the output from this drive instance

				for _, printAll := range driver.totalOutput {
					fmt.Println(printAll)
				}
	}
}
