package termin

import "os"

type Termin struct {
	out chan []byte
}

func New() *Termin {
	return &Termin{
		out: make(chan []byte, 1),
	}
}

func (t *Termin) GetChan() chan []byte {
	return t.out
}

func (t *Termin) StartReading(async bool) {
	if async {
		go t.catch()
	} else {
		t.catch()
	}
}

func (t *Termin) catch() {
	for {
		var bytes = make([]byte, 6)
		n, err := os.Stdin.Read(bytes)
		if err != nil {
			panic(err)
		}
		if n < 1 {
			continue
		}
		t.out <- bytes
	}
}
