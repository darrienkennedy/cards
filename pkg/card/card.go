package card

import (
	"fmt"
)

type Card struct {
	Value CardValue
	Suit  CardSuit
}

func NewCard(v CardValue, s CardSuit) *Card {
	return &Card{
		Value: v,
		Suit:  s,
	}
}

func (c *Card) Equals(other *Card) bool {
	return c.Value == other.Value && c.Suit == other.Suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value.String(), c.Suit.String())
}

func (c *Card) Runes() string {
	return fmt.Sprintf("[%c%c]", c.Value.Char(), c.Suit.Rune())
}
