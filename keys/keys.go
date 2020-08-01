package keys

func NewKeyboardKey(r rune) (KeyboardKey, bool) {
	if IsControlChar(r) {
		return controlChar(r), true
	}
	if IsSymbolChar(r) {
		return symbolChar(r), true
	}
	return nil, false
}

func IsSymbolChar(r rune) bool {
	return isBetween(r, Space, Delete)
}

func IsControlChar(r rune) bool {
	return isBetween(r, NullChar, UnitSeparator)
}

func isBetween(r rune, left, right KeyboardKey) bool {
	if r >= left.getRune() && r <= right.getRune() {
		return true
	}
	return false
}

