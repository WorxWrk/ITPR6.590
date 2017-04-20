package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pathExtraction() {

}
func cityExtraction() {

}
func startCity() {

}
func pathDecision() {

}
func exitDecision() {

}
func pathReverse() {

}

// IF for the counter for additional text output (Handled in Main)

func main() {

	var input string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Path Input: ")
	scanner.Scan()
	input = scanner.Text()
	PathInt := strconv.ParseInt(input, 10, 32)

	fmt.Println("Enter City Input: ")
	scanner.Scan()
	input = scanner.Text()
	CityInt := strconv.ParseInt(input, 10, 32)

	var path [4]string
	path[0] = "Path 1 "
	path[1] = "Path 2"
	path[2] = "Path 3"
	path[3] = "Path 4"

	var city [4]string
	city[0] = "City One"
	city[1] = "City Two"
	city[2] = "City Three"
	city[3] = "City Four"

	fmt.Println(city)

	fmt.Println("You took ", path[PathInt], " to get to ", city[CityInt])

	fmt.Println(path)
}
