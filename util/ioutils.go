package util

import (
	"errors"
	"io"
	"strings"
)

// PumpStream 从in读数据，并写到out，直到读到EOF
func PumpStream(in io.Reader, out io.Writer, buffer []byte) error {

	if buffer == nil {
		buffer = make([]byte, 1024)
	}

	for {
		n1, err := in.Read(buffer)
		if err != nil {
			if "EOF" == strings.ToUpper(err.Error()) {
				break
			}
			return err
		}
		n2, err := out.Write(buffer[0:n1])
		if err != nil {
			return err
		}
		if n1 != n2 {
			return errors.New("len(read) != len(write)")
		}
	}

	return nil
}
