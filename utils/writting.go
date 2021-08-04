package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func WriteString(w io.Writer, value string) error {
	_, err := w.Write([]byte(value))
	return err
}

func WriteInt16(w io.Writer, value int16) error {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(value))
	_, err := w.Write(buf)
	return err
}

func WriteInt64(w io.Writer, value int64) error {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(value))
	_, err := w.Write(buf)
	return err
}

func WriteEmpty(w io.Writer, amount int) error {
	for amount > 10240 {
		if err := writeBatch(w, 10240); err != nil {
			return err
		}
		amount = amount - 10240
	}

	return writeBatch(w, amount)
}

func writeBatch(w io.Writer, amount int) error {
	payload := bytes.Repeat([]byte(""), amount)
	_, err := w.Write(payload)
	if err != nil {
		return fmt.Errorf("could not write zero bytes, %v", err)
	}

	return nil
}