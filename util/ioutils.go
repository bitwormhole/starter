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
		n1, err1 := in.Read(buffer)
		if n1 > 0 {
			n2, err2 := out.Write(buffer[0:n1])
			if err2 != nil {
				return err2
			}
			if n1 != n2 {
				return errors.New("len(read) != len(write)")
			}
		}
		if err1 != nil {
			if "EOF" == strings.ToUpper(err1.Error()) {
				break
			}
			return err1
		}
	}

	return nil
}
