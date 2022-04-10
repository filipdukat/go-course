package main

import "fmt"

func main() {
	cards := newDeck()
	//	hand, remainigCards := deal(cards, 5)
	shuffledDeck := shuffle(cards)

	//	shuffledDeck.print()
	fmt.Println(shuffledDeck.toString())
}
