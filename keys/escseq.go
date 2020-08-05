package keys

type EscapeSequence interface {
	KeyboardKey
	String() string
}

func NewEscapeSequence(runes []rune) EscapeSequence {
	if runes[0] != Escape.Rune() {
		return nil
	}
	if runes[1] != OpeningBracket.Rune() {
		return nil
	}
	var result string
	for _, r := range runes {
		result += string(r)
	}
	return escapeSequence(result)
}

type escapeSequence string

func (e escapeSequence) Rune() rune {
	return rune(e[len(e)-1])
}

func (e escapeSequence) String() string {
	return string(e)
}

func (e escapeSequence) distinguish() {}
