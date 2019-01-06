package main

func main()  {

	cards := newDeck()
	//cards.print()
	cards.shuffle()
	cards.print()
	//hand,remainingCards := deal(cards,5)
	//
	//hand.print()
	//fmt.Println("------------------")
	//remainingCards.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//
	//cards := newDeckFromFile("my_cds")
	//cards.print()

}
