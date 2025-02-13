package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()

	if hour >= 23 || hour < 7 {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten")
	} else if hour < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken")
	} else if hour < 18 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken")
	} else {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken")
	}
}