package cardlist

import (
	"fmt"
	"github.com/darrienkennedy/cards/pkg/card"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PopCard(t *testing.T) {
	tests := []struct {
		list   *CardList
		card   *card.Card
		expect *CardList
		pass   bool
		err    bool
	}{
		{
			list:   NewCardList(),
			card:   &card.Card{},
			expect: NewCardList(),
			pass:   false,
			err:    true,
		},
		{
			list: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.Nine, card.Clubs),
			}),
			card: card.NewCard(card.Jack, card.Diamonds),
			expect: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Nine, card.Clubs),
			}),
			pass: true,
			err:  false,
		},
		{
			list: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.Nine, card.Clubs),
			}),
			card: card.NewCard(card.Nine, card.Clubs),
			expect: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
			}),
			pass: false,
			err:  false,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("tests[%d]", i)
		t.Run(testname, func(t *testing.T) {
			initialN := tt.list.n
			c, e := tt.list.PopCard()

			if tt.err {
				assert.Error(t, e)
			} else if tt.pass {
				assert.NoError(t, e)
				assert.Equal(t, tt.list.n, initialN-1)
				assert.True(t, c.Equals(tt.card))
				assert.True(t, tt.list.equals(tt.expect))
			} else {
				assert.False(t, c.Equals(tt.card))
			}
		})
	}
}

func Test_PushCard(t *testing.T) {
	tests := []struct {
		list   *CardList
		card   *card.Card
		expect *CardList
		pass   bool
	}{
		{
			list: NewCardList(),
			card: card.NewCard(card.Ace, card.Hearts),
			expect: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Ace, card.Hearts),
			}),
			pass: true,
		},
		{
			list: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.Nine, card.Clubs),
			}),
			card: card.NewCard(card.King, card.Spades),
			expect: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.King, card.Spades),
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.Nine, card.Clubs),
			}),
			pass: true,
		},
		{
			list: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.Nine, card.Clubs),
			}),
			card: card.NewCard(card.King, card.Spades),
			expect: NewCardListFromSlice([]*card.Card{
				card.NewCard(card.Jack, card.Diamonds),
				card.NewCard(card.King, card.Spades),
				card.NewCard(card.Nine, card.Clubs),
			}),
			pass: false,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("tests[%d]", i)
		t.Run(testname, func(t *testing.T) {
			initialN := tt.list.n
			tt.list.PushCard(tt.card)

			if tt.pass {
				assert.Equal(t, tt.list.n, initialN+1)
				assert.True(t, tt.list.equals(tt.expect))
			} else {
				assert.False(t, tt.list.equals(tt.expect))
			}
		})
	}
}
