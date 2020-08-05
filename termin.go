package termin

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/KlyuchnikovV/termin/keys"
	"github.com/KlyuchnikovV/termin/low_level/raw_mode"
)

type Termin struct {
	out chan keys.KeyboardKey

	in   *bufio.Reader
	file *os.File
}

func New() *Termin {
	return &Termin{
		out: make(chan keys.KeyboardKey, 1),
	}
}

func (t *Termin) GetChan() chan keys.KeyboardKey {
	return t.out
}

func (t *Termin) StartReading(async bool) {
	if t.out == nil {
		t.out = make(chan keys.KeyboardKey, 1)
	}
	if async {
		go t.catch()
	} else {
		t.catch()
	}
}

func (t *Termin) Stop() {
	close(t.out)
	t.file.Close()
}

func (t *Termin) catch() {
	raw_mode.EnableRawMode()
	defer raw_mode.DisableRawMode()

	var err error
	t.file, err = os.Open("/dev/tty")
	if err != nil {
		t.file = os.Stdin
	}

	t.in = bufio.NewReader(t.file)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR: %#v", err)
		}
	}()

	for {
		r, _, err := t.in.ReadRune()
		if err != nil {
			// TODO: remake
			panic(err)
		}

		key := t.processRune(r)
		if key == nil {
			panic(fmt.Errorf("rune %c is not valid (%#v)", r, r))
		}

		t.out <- key
	}
}

func (t *Termin) processRune(r rune) keys.KeyboardKey {
	var key keys.KeyboardKey
	if keys.RuneIsKey(r, keys.Escape) {
		key = t.processEscape()
	} else {
		key = keys.NewKeyboardKey(r)
	}
	return key
}

func (t *Termin) processEscape() keys.KeyboardKey {
	var bufferSize = t.in.Buffered()
	if bufferSize == 0 {
		return keys.Escape
	}

	var ok bool
	var runes = make([]rune, bufferSize+1)
	runes[0] = rune(keys.Escape)
	runes[1], ok = t.readOpeningBracket()
	if !ok {
		return keys.Escape
	}

	for i := 2; i < len(runes); i++ {
		r, _, err := t.in.ReadRune()
		if err != nil {
			panic(err)
		}
		runes[i] = r
	}

	return keys.NewEscapeSequence(runes)
}

func (t *Termin) readOpeningBracket() (rune, bool) {
	runes, err := t.in.Peek(1)
	if err != nil {
		panic(err)
	}
	var r = rune(runes[0])
	if r != rune(keys.OpeningBracket) {
		return r, false
	}
	_, _, err = t.in.ReadRune()
	if err != nil {
		panic(err)
	}

	return r, true
}
