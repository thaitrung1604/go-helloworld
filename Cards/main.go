package main

func main() {
	cards := deck{newCard(), "Ace", "Two"}
	cards = append(cards, "Three")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
