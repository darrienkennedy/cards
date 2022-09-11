package main

import (
	"fmt"
	"github.com/darrienkennedy/cards/pkg/card"
	"github.com/darrienkennedy/cards/pkg/cardlist"
)

func main() {
	deck := cardlist.NewCardList()
	deck.PushCard(card.NewCard(card.Three, card.Clubs))

	fmt.Println(deck.String())
}
