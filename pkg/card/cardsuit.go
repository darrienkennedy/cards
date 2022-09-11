package card

var CardSuits = [4]CardSuit{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}

type CardSuit int

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
	NullSuit
)

func (cs *CardSuit) String() string {
	switch *cs {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	case NullSuit:
		return "null"
	default:
		return "Undefined"
	}
}

func (cs *CardSuit) Rune() rune {
	switch *cs {
	case Clubs:
		return '♣'
	case Diamonds:
		return '♦'
	case Hearts:
		return '♥'
	case Spades:
		return '♠'
	case NullSuit:
		return '_'
	default:
		return 'u'
	}
}
