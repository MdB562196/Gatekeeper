package main

import (
	"fmt"
	"time"
)

func main() {
	kentekens := map[string]bool{
		"HL-932-K": true,
		"X-789-PT": true,
		"94-GN-RP": true,
	}

	var input string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&input)

	if !kentekens[input] {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
		return
	}

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