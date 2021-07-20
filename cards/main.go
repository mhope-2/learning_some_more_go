package main


func main() {
	cards := Deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	cards.Print()
}

func newCard() string{
	return "Five of Diamonds"	

}  