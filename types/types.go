package types

import "io"

type FileReader struct {
	src io.Reader
	out io.ReadCloser
}



