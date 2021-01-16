package main

func main() {
  cards := newDeckFromFile("cards_test")
  cards.shuffle()
  cards.print()

}


