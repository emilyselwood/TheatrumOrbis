package TheatrumOrbis

import "encoding/binary"

type Type int
const (
	Integer Type = 0
	Double  Type = 1
	String  Type = 2
	Bytes   Type = 3
)

type Value interface {
	/*
	T returns the type of this value
	 */
	T() Type
	/*
	Bytes returns the raw bytes for this value
	 */
	Bytes() []byte
	/*
	Decode will decode the content of the byte array into this Value object
	It will also return a reference to this value object.
	 */
	Decode([]byte) Value
}


type IntegerValue struct {
	v int64
}

func (i *IntegerValue) T() Type {
	return Integer
}

func (i *IntegerValue) Bytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i.v))
	return b
}

func (i *IntegerValue) Decode(b []byte) *IntegerValue {
	i.v = int64(binary.BigEndian.Uint64(b))
	return i
}

func (i *IntegerValue) Value() int64 {
	return i.v
}

