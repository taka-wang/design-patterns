package main

import (
	"fmt"

	"./motion" // bad habit
)

func main() {
	//var cards []motion.Card
	//cards = make([]motion.Card, 2)
	var cards [2]motion.ICard
	cards[0] = motion.Create("8338")
	cards[1] = motion.Create("7856")

	for _, c := range cards {
		if c != nil {
			fmt.Println("Card Number: " + c.GetName())
			c.Move()
		}
	}
}
