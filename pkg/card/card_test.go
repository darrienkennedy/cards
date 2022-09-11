package card

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Equals(t *testing.T) {
	tests := []struct {
		a      *Card
		b      *Card
		expect bool
	}{
		{
			a:      NewCard(Ace, Spades),
			b:      NewCard(Ace, Spades),
			expect: true,
		},
		{
			a:      NewCard(Ace, Diamonds),
			b:      NewCard(Ace, Hearts),
			expect: false,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test[%d]", i)
		t.Run(testname, func(t *testing.T) {
			result := tt.a.Equals(tt.b)
			assert.Equal(t, result, tt.expect)
		})
	}
}
