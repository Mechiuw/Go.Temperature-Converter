package main

import (
	"GoTempConverter/src/temp"
	"fmt"
)

func main() {
	userInput()
}

func userInput() {
	var c temp.Celsius
	c.Temperature = 20
	c.Category = "cold"

	fmt.Println(c)
}
