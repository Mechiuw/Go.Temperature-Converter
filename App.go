package main

import (
	"GoTempConverter/src/temp"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	userInput()
}
func menu() {
	fmt.Println("a) celsius -> fahrenheit")
	fmt.Println("b) fahrenheit -> kelvin")
	fmt.Println("c) kelvin -> rankine")
	fmt.Println("d) rankine -> reaumur")
}

func userInput() {
	menu()
	fmt.Println("choice : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	switch input {
	case "a":
		fmt.Println("insert the celsius temperature (int)")
		scanner.Scan()
		temperatures, _ := strconv.Atoi(scanner.Text())

		fmt.Println("insert the category ")
		scanner.Scan()
		categories := scanner.Text()
		createC(temperatures, categories)
	case "b":
		fmt.Println("insert the fahrenheit temperature (int)")
		scanner.Scan()
		tempF, _ := strconv.Atoi(scanner.Text())

		fmt.Println("insert the category ")
		scanner.Scan()
		catF := scanner.Text()
		createF(tempF, catF)
	}
}

func createC(temperature int, category string) {
	var c temp.Celsius
	m := make(map[string]int)

	c.Temperature = temperature
	c.Category = category

	m[c.Category] = c.Temperature

	for k, v := range m {
		fmt.Printf("category : %s temperature : %d", k, v)
	}
}

func createF(temp int, categ string) {

}
