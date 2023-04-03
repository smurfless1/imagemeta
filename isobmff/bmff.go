package isobmff

import (
	"io"

	"github.com/smurfless1/imagemeta/meta"
)

type ExifReader func(r io.Reader, h meta.ExifHeader) error

const (
	optionSpeed uint8 = 1
)
