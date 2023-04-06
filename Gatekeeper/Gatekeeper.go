package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	timeStr := currentTime.Format("15:04")
	welkomStr := ("Welkom bij Fonteyn Vakantieparken")

	if timeStr >= "07:00" && timeStr <= "12:00" {
		fmt.Println("Goedemorgen!:", welkomStr)
	} else if timeStr > "12:00" && timeStr <= "18:00" {
		fmt.Println("Goedemiddag:", welkomStr)
	} else if timeStr > "18:00" && timeStr <= "23:00" {
		fmt.Println("Goedenavond:", welkomStr)
	} else {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten. The current time is:")
	}
}
