package deck

type Color int8
const (
	Unknown = iota - 1
	Red
	Black
)
func (c Color) String() string{
	switch c {
	case Red: return "Red"
	case Black: return "Black"
	default: return "Unknown"
	}
}

type Suit int8
const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
		case Diamonds: return "Diamonds"
		case Clubs: return "Clubs"
		case Hearts: return "Hearts"
		case Spades: return "Spades"
		default: return "Unknown"
	}
}
func (s Suit) Color() Color {
	switch s {
		case Diamonds: return Red
		case Clubs: return Black
		case Hearts: return Red
		case Spades: return Black
		default: return Unknown
	}
}

type Value int8
const (
	Ace Value = iota + 1
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
)

func (v Value) String() string {
	switch v {
		case Ace: return "Ace"
		case Two: return "Two"
		case Three: return "Three"
		case Four: return "Four"
		case Five: return "Five"
		case Six: return "Six"
		case Seven: return "Seven"
		case Eight: return "Eight"
		case Nine: return "Nine"
		case Ten: return "Ten"
		case Jack: return "Jack"
		case Queen: return "Queen"
		case King: return "King"
		default: return "Unknown"
	}
}

type Card struct {
	Suit Suit
	Value Value
}

func (c Card) String() string {
	return c.Value.String() + " of " + c.Suit.String()
}

func (c Card) Color() Color {return c.Suit.Color()}

type Deck struct {
	Cards []Card
}