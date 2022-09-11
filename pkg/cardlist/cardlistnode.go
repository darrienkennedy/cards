package cardlist

import (
	"github.com/darrienkennedy/cards/pkg/card"
)

type CardListNode struct {
	card *card.Card
	next *CardListNode
}

func NewCardListNode(c *card.Card) *CardListNode {
	return &CardListNode{
		card: c,
	}
}
