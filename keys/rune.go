package keys

type RuneKey interface {
	KeyboardKey
	Rune() rune
	Is(r RuneKey) bool
}

type symbolChar rune

func (s symbolChar) Rune() rune {
	return rune(s)
}

func (s symbolChar) distinguish() {}

func (s symbolChar) Is(r RuneKey) bool {
	return s.Rune() == r.Rune()
}

type controlChar rune

func (c controlChar) Rune() rune {
	return rune(c)
}

func (c controlChar) distinguish() {}

func (c controlChar) Is(r RuneKey) bool {
	return c.Rune() == r.Rune()
}
