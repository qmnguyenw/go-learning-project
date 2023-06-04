package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello World!!!")
	cards := initDeck()
	cards.printDeck()
	for {
		fmt.Println("------------------------")
		fmt.Println("Deck Project")
		fmt.Println("1. Init new deck")
		fmt.Println("2. Print deck")
		fmt.Println("3. Shuffle deck")
		fmt.Println("4. Deal deck")
		fmt.Println("5. Save deck to File")
		fmt.Println("6. Load deck from File")
		fmt.Println("7. Exit")

		// read user's input for choosing option
		input_reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter option: ")
		input_text, _ := input_reader.ReadString('\n')
		// fmt.Println(input_text)

		switch strings.TrimSpace(input_text) {
		case "1":
			// fmt.Println("1")
			cards = initDeck()
		case "2":
			// fmt.Println("2")
			cards.printDeck()
		case "3":
			// fmt.Println("3")
			cards.shuffleDeck()
		case "4":
			// fmt.Println("4")
			card1, card2 := dealDeck(cards, 3)
			fmt.Println("Card1", card1.toString())
			fmt.Println("Card2", card2.toString())
		case "5":
			// fmt.Println("5")
			cards.saveDeck("deck.txt")
		case "6":
			// fmt.Println("6")
			cards = loadDeck("deck.txt")
		case "7":
			// fmt.Println("7")
			os.Exit(1)
		default:
			fmt.Println(reflect.TypeOf(input_text))
			fmt.Println(strings.TrimSpace(input_text) == "1")
			fmt.Println("Please choose valid option.")
		}

		// read user's input for confirming exit
		continue_reader := bufio.NewReader(os.Stdin)
		fmt.Println("Do you want to continue: ")
		continue_text, _ := continue_reader.ReadString('\n')
		// fmt.Println(continue_text)

		if (strings.ToLower(strings.TrimSpace(continue_text)) != "y") {
			fmt.Println("Have a nice day")
			os.Exit(1)
		}
	}

}
