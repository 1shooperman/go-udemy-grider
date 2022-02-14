package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["foo"] = "bar"

	// delete(colors, "foo")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v: %v\n", color, hex)
	}
}
