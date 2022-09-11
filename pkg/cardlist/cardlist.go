package cardlist

import (
	"fmt"
	"github.com/darrienkennedy/cards/pkg/card"
	"strings"
)

type CardList struct {
	head *CardListNode
	n    int
}

func NewCardList() *CardList {
	return &CardList{}
}

func NewCardListFromSlice(cards []*card.Card) *CardList {
	cl := NewCardList()
	// to preserve order, iterate through slice in reverse
	for i := len(cards) - 1; i >= 0; i-- {
		cl.PushCard(cards[i])
	}
	return cl
}

func (cl *CardList) PopCard() (*card.Card, error) {
	if cl.n == 0 || cl.head == nil {
		return &card.Card{}, fmt.Errorf("card list is empty")
	}

	picked := cl.head.card
	cl.head = cl.head.next
	cl.n -= 1
	return picked, nil
}

func (cl *CardList) PushCard(c *card.Card) {
	node := NewCardListNode(c)
	prevhead := cl.head
	cl.head = node
	cl.head.next = prevhead
	cl.n += 1
}

func (cl *CardList) String() string {
	var sb strings.Builder
	current := cl.head
	for i := 0; i < cl.n; i++ {
		sb.WriteString(current.card.Runes() + " ")
		if current.next != nil {
			current = current.next
		} else {
			break
		}
	}
	return sb.String()
}

func (cl *CardList) equals(other *CardList) bool {
	if cl.n != other.n {
		return false
	}

	a := cl.head
	b := other.head
	for i := 0; i < cl.n; i++ {
		if !a.card.Equals(b.card) {
			return false
		}
		a = a.next
		b = b.next
	}
	return true
}
