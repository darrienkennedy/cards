package card

var CardValues = [13]CardValue{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

type CardValue int

const (
	Ace CardValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	NullValue
)

func (cv *CardValue) String() string {
	switch *cv {
	case Ace:
		return "Ace"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case NullValue:
		return "null"
	default:
		return "undefined"
	}
}

func (cv *CardValue) Char() uint8 {
	switch *cv {
	case Ace:
		return 'A'
	case Two:
		return '2'
	case Three:
		return '3'
	case Four:
		return '4'
	case Five:
		return '5'
	case Six:
		return '6'
	case Seven:
		return '7'
	case Eight:
		return '8'
	case Nine:
		return '9'
	case Ten:
		return 't'
	case Jack:
		return 'J'
	case Queen:
		return 'Q'
	case King:
		return 'K'
	case NullValue:
		return '_'
	default:
		return 'u'
	}
}
