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

		createC(temperatures)
	case "b":
		fmt.Println("insert the fahrenheit temperature (int)")
		scanner.Scan()
		tempF, _ := strconv.Atoi(scanner.Text())

		createF(tempF)

	case "c":
		fmt.Println("insert the kelvin temperature (int)")
		scanner.Scan()
		tempK, _ := strconv.Atoi(scanner.Text())

		createK(tempK)
	}

}

func createC(temperature int) {
	var c temp.Celsius

	if temperature > 23 {
		c = temp.Celsius{Temperature: temperature, Category: "hot temperature"}
		fmt.Println("celsius temperature : ", c)
	} else {
		c = temp.Celsius{Temperature: temperature, Category: "cold temperature"}
		fmt.Println("celsius temperature : ", c)
	}
	toFahrenheit := temp.ToFahrenheit(c)
	fmt.Println("celsius -> fahrenheit :", toFahrenheit, "Fahrenheit")
}

func createF(tempF int) {
	var f temp.Fahrenheit

	if tempF > 23 {
		f = temp.Fahrenheit{Temperature: tempF, Category: "hot temperature"}
		fmt.Println("fahrenheit temperature :", f)
	} else {
		f = temp.Fahrenheit{Temperature: tempF, Category: "cold temperature"}
		fmt.Println("fahrenheit tempereture :", f)
	}
	toKelvin := temp.ToKelvin(f)
	fmt.Println("fahrenheit -> kelvin", toKelvin, "Kelvin")
}

func createK(tempK int) {
	var k temp.Kelvin

	if tempK > 22 {
		k = temp.Kelvin{Temperature: tempK, Category: "hot temperature"}
		fmt.Println("kelvin temperature :", k)
	} else {
		k = temp.Kelvin{Temperature: tempK, Category: "cold temperature"}
		fmt.Println("kelvin temperature :", k)
	}

	toRankine := temp.ToRankine(k)
	fmt.Println("kelvin -> rankine", toRankine, "Rankine")
}

func createRan(tempRan int) {
	var ran temp.Rankine
}
