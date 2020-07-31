package termin

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/KlyuchnikovV/termin/low_level/raw_mode"
)

type Termin struct {
	out    chan rune

	file *os.File
}

func New() *Termin {
	return &Termin{
		out: make(chan rune, 1),
	}
}

func (t *Termin) GetChan() chan rune {
	return t.out
}

func (t *Termin) StartReading(async bool) {
	if t.out == nil {
		t.out = make(chan rune, 1)
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

	in := bufio.NewReader(t.file)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR: %#v", err)
		}
	}()

	for {
		r, _, err := in.ReadRune()
		if err != nil && err != io.EOF {
			// TODO: remake
			panic(err)
		}
		t.out <- r
	}
}
