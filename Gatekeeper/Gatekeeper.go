package main

import (
	"bufio"   // Bibliotheek voor buffered I/O, waardoor het efficiënter is om gegevens te lezen en te schrijven.
	"fmt"     // Bibliotheek voor formattering van output en input.
	"os"      // Bibliotheek voor communicatie met het besturingssysteem.
	"strings" // Bibliotheek voor string manipulatie.
	"time"    // Bibliotheek voor tijd manipulatie.
)

// Functie voor het controleren of een kenteken in de lijst met toegestane kentekens zit.
func checkName(kenteken string) bool {
	allowedkenteken := []string{"82nnk5", "nk76b5", "nnk825", "825nnk"}

	kenteken = strings.ToLower(kenteken) // Zet het ingevoerde kenteken in kleine letters.
	for _, allowed := range allowedkenteken {
		if kenteken == allowed {
			return true // Kenteken is toegestaan.
		}
	}

	return false // Kenteken is niet toegestaan.
}

// Functie voor het begroeten van de bezoeker op basis van het tijdstip.
func greetVisitor(welkomStr string) {
	currentTime := time.Now().Format("15:04") // Vind de huidige tijd en zet het in uur:minuut formaat.

	if currentTime >= "07:00" && currentTime <= "12:00" {
		fmt.Println("Goedemorgen!:", welkomStr)
	} else if currentTime > "12:00" && currentTime <= "18:00" {
		fmt.Println("Goedemiddag:", welkomStr)
	} else if currentTime > "18:00" && currentTime <= "23:00" {
		fmt.Println("Goedenavond:", welkomStr)
	} else {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten.")
	}
}

// Functie voor het verwerken van de input van de gebruiker en het uitvoeren van de begroeting.
func processInput() {
	welkomStr := "Welkom bij Fonteyn Vakantieparken"

	fmt.Print("Voer je kenteken in: ")
	reader := bufio.NewReader(os.Stdin)    // Maak een nieuwe reader aan voor input van de gebruiker.
	kenteken, _ := reader.ReadString('\n') // Lees de ingevoerde string tot aan de volgende nieuwe regel.
	kenteken = strings.TrimSpace(kenteken) // Verwijder eventuele extra spaties voor of na het kenteken.

	if checkName(kenteken) {
		greetVisitor(welkomStr)
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein.")
	}
}

func main() {
	processInput()
}
