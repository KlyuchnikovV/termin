package keys

type KeyboardKey interface {
	// Does nothing but helps to disallow
	// interface implementation elsewhere
	distinguish()
}

func NewKeyboardKey(r rune) KeyboardKey {
	if isBetween(r, NullChar, UnitSeparator) {
		return controlChar(r)
	}
	if isBetween(r, Space, Delete) {
		return symbolChar(r)
	}
	return nil
}

func IsSymbolChar(k RuneKey) bool {
	return isBetween(k.Rune(), Space, Delete)
}

func IsControlChar(k RuneKey) bool {
	return isBetween(k.Rune(), NullChar, UnitSeparator)
}

func isBetween(r rune, left, right RuneKey) bool {
	if r >= left.Rune() && r <= right.Rune() {
		return true
	}
	return false
}

func ToEscapeSequence(k KeyboardKey) (EscapeSequence, bool) {
	result, ok := k.(EscapeSequence)
	return result, ok
}

func ToRuneKey(k KeyboardKey) (RuneKey, bool) {
	result, ok := k.(RuneKey)
	return result, ok
}

func RuneIsKey(r rune, k RuneKey) bool {
	if _, ok := ToEscapeSequence(k); ok {
		return false
	}
	return r == k.Rune()
}

func StringIsSequence(s string, e EscapeSequence) bool {
	return s == e.String()
}
