package main

func main() {
	cards := newDeckFromFile("cards_list")
	cards.shuffle()
	cards.print()
}
