package fs

import (
	"bytes"
	"testing"
)

// tests for FileIO

func TestWriteText_ReadText(t *testing.T) {

	text0 := "hello,world"
	dir := prepareDirForTest(t)
	file := dir.GetChild("test.writeText")
	fio := file.GetIO()

	err := fio.WriteText(text0, nil)
	if err != nil {
		t.Error(err)
		return
	}

	text1, err := fio.ReadText()
	if err != nil {
		t.Error(err)
		return
	}

	if text0 == text1 {
		t.Log("FileIO.ReadText()  &  FileIO.WriteText() ok")
	} else {
		t.Error("FileIO.ReadText() != FileIO.WriteText() ")
	}
}

func TestWriteBinary(t *testing.T) {

	bin0 := []byte{0, 0x99, 0xcd, 0xff}
	dir := prepareDirForTest(t)
	file := dir.GetChild("test.writeBinary")
	fio := file.GetIO()

	err := fio.WriteBinary(bin0, nil)
	if err != nil {
		t.Error(err)
		return
	}

	bin1, err := fio.ReadBinary()
	if err != nil {
		t.Error(err)
		return
	}

	if bytes.Equal(bin0, bin1) {
		t.Log("FileIO.ReadBinary()  &  FileIO.WriteBinary() ok")
	} else {
		t.Error("FileIO.ReadBinary() != FileIO.WriteBinary()")
	}

}
