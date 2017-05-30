package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"

)

func startingCity() {

}

func count(){

}

func driveCity(newNumber int, currentCity int, allCities [4][3][2]string ) (int,string,string,string,string) {

// Variable Initialization
	var pathName string
	var exitName string
	var location [3][2]string
	currentCityName := location[1][0]

//
	if newNumber == 1 { // right side path, int goes up
		if currentCity == 3 { // restricts int size
			currentCity = 0
		}else{
			currentCity = currentCity + 1
		}
		location = allCities[currentCity]
		pathName  = location[2][1]
		exitName  = location[2][0]
	}

//
	if newNumber == 0{ // left side path, int goes down
		if currentCity == 0 { //restricts int size
			currentCity = 3
		}else {
			currentCity = currentCity - 1
		}
		location = allCities[currentCity]
		pathName = location[0][1]
		exitName = location[0][0]
	}

	nextCityName := location[1][0]

// Return the Values that need to be set after the function is complete
	return currentCity,currentCityName,pathName,exitName,nextCityName

}

func driveExit(){

}

func extras(){

}

func printer( ) {

}


/*
func simulation () {
	//------------------- 1. simulation initialization
	var key bool = false

	// new rnd number
	var currentCity int = 0 // rnd

	var location [3][2]string
	location = allCities[currentCity]
	// variable that holds a single city section/quadrant
	//------------------ 2.

	for key == false{  //Main loop for a single

		var akinaCount int = 0

		// new RND number
		var pathSelection int = 0 //rnd

		var currentCityName string = location[1][0]

		if nextCityName == "Akina" {
		akinaCount++
		}

		//------------------- 3. City movement
		if pathSelection == 1 { // right side path, int goes up
			if currentCity == 3 { // restricts int size
				currentCity = 0
			}else{
				currentCity = currentCity + 1
			}
			location = allCities[currentCity]
			var pathName string = location[2][1]
			var exitName string = location[2][0]
		}

		if pathSelection == 0{ // left side path, int goes down
			if currentCity == 0 { //restricts int size
				currentCity = 3
			}else {
				currentCity = currentCity - 1
			}
			location = allCities[currentCity]
			var pathName string = location[0][1]
			var exitName string = location[0][0]
		}

		var nextCityName string = location[1][0]


		//-------------------------- 4. Extras
		// Moved to the Top of the Loop to catch the starting City
		//-------------------------- 5. Exitting and Printing

		// new rnd
		var exitSelection int = 0 //RND

		if exitSelection == 2 { //  number that allows the driver to exit
			// related Printing Info
			key = true
		}else{     // Exiting failed
			     // related Printing info
		}
	//------------ Printing Extras

	fmt.Println("Driver", "<driverNumber>", "Met With Jhon Jamieson", akinaCount, " Times!")

	if akinaCount != 1 || 2{
		if akinaCount == 0{
			fmt.Println("That Driver Missed Out!")
		}
		if akinaCount <= 3{
			fmt.Println("That Driver Needed Lots Of Help!")
		}
	}
}
}
*/

// IF for the counter for additional text output (Handled in Main)

func main() {

// --------------------------------------------- Start/Varible Initialization
// Array Creation - Creates 4, 2 Dimentional Arrays.
// During Comments, a single one of these arrays might be refered to as City Sections or Quadrants
// Looking back. Should have used a struct instead of the 2-D arrays

// City 1
	var City1 [3][2] string
	City1 [0][0] = "Exit 1"
	City1 [0][1] = "Path 1"
	City1 [1][0] = "City"
	City1 [1][1] = ""
	City1 [2][0] = "Exit 2"
	City1 [2][1] = "Path2"

//City 2
	var City2 [3][2] string
	City2 [0][0] = "Exit 2"
	City2 [0][1] = "Path 2"
	City2 [1][0] = "City1"
	City2 [1][1] = ""
	City2 [2][0] = "Exit 3"
	City2 [2][1] = "Path 3"

//City 3
	var City3 [3][2] string
	City3 [0][0] = "Exit 3"
	City3 [0][1] = "Path 3"
	City3 [1][0] = "City2"
	City3 [1][1] = ""
	City3 [2][0] = "Exit 4"
	City3 [2][1] = "Path 4"

//City 4
	var City4 [3][2] string
	City4 [0][0] = "Exit 4"
	City4 [0][1] = "Path 4"
	City4 [1][0] = "City3"
	City4 [1][1] = ""
	City4 [2][0] = "Exit 1"
	City4 [2][1] = "Path 1"

// Array Creation - Create a single array to hold the previous 4 arrays.
// Array for holding all city sections/quadrants.
	var allCities [4] [3][2] string
	allCities[0] = City1
	allCities[1] = City2
	allCities[2] = City3
	allCities[3] = City4

// Varible Creation
// Variables will be initialized later, during the loops

// Used to keep a driver in a loop
	var key bool = false;

// Used to hold the Number given by the Random number generator
	//var pathSelection int

// Used to hold the Number given by the Random number generator
	//var exitSelection int

// Used to hold the Number given by the Random number generator
	var newNumber int

//Holds the conveted input from the user
	var seed int

// Holds the number of the current city, within the "allCities" array.
//Used to set the "location" variable
	var currentCity int

// Used to hold an entire City Section
	var location [3][2]string

// Used to track the amount of times the driver has visted "Akina"
	var akinaCount int

// Stores the name of the City the driver is in before the driver moves to the next city
	var currentCityName string

// Used to hold the current driver. Changes depending on the loop number
	var driverNumber int = 0;

// Holds the string value of the path that has been selected
	var pathName string

// Holds the string value of the exit that has been selected
	var exitName string

// Holds name of the city that has been selected next
	var nextCityName string




// --------------------------------------------- Program Begins
// RNG/Seed Initialization

 	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a number:");
	scanner.Scan()
	input := scanner.Text()

	seed = /*input*/ 5

	rand.Seed(seed)


// Loop to complete multiiple driving instances
	for i:=0; i>5; i++{

// Varible re-Initialization to start a driving instance

			key = false
			driverNumber++
			akinaCount = 0

// Finding a Starting City for the driving instance

			startingCity();

// Loop for the moving sections of a single driving instance
			for key == false{

// Path Decision. What path the driver will take out of the current city.
				//new rnd
				//set rnd Variables
				newNumber = rand.Intn(2)

				currentCity,currentCityName,pathName,exitName,nextCityName = driveCity(newNumber, currentCity, allCities );
				// printing stuff (just print Or move the print text to a variable to be printed later)

//Exit Decision. If the Driver will use an exit or continue to the city.
				// new rnd
				// set RND varablies
				newNumber = rand.Intn(6)

				driveExit();
				// printing stuff (just print Or move the print text to a variable to be printed later)
			}

// Additional Dialouge. The exta line required at the end of a driving instance.

				extras();
				printer();
	}




						fmt.Println(seed,location,akinaCount,currentCityName,pathName,exitName,nextCityName,input)

/*



var input int = 0 // if statments for managing the -numbers and over 3 (4+),  over 3 will rest to 0,  - numbers (under 0) will increase to 3
var input2 int = 1 // 2 possiblities, need an if for number changeing, eg 0 =0 and 1 = 2, this is due to the city name being inbetween the 2 path groups
var section = complie[input] // variable that holds a single city section/quadrant

var cityname string = section[1][0]
var pathname string = section[2][1]

if input2 == 1 { // will need to keep the varialbes above for handling the names as the "section" state will change without the statement
	section = complie[input + input2]  // right side path, int goes up
	fmt.Println("The Driver is headed to",section[1][0],"from", cityname, "using the path", pathname)
}else{
	section = complie[input - input2]  // left side path, int goes down
}

	//var input string
/*
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Path Input: ")
	scanner.Scan()
	input = scanner.Text()

	fmt.Println("Enter City Input: ")
	scanner.Scan()
	input = scanner.Text()



	//fmt.Println(city)
	fmt.Println(path1)
	fmt.Println(complie)
	fmt.Println(aCity)
	fmt.Println(section)



	fmt.Println("-------------")
	var new = complie[0]
	fmt.Println(new[1][0])

	*/
}
