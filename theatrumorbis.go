package TheatrumOrbis
/*
TheatrumOrbis is a single file database like a very simple sqlLite but pure go with no external dependencies.
(this last bit is possibly negotiable but for now it would be nice) defiantly no cgo

The name TheatrumOrbis is one of the first atlas' This is my first time building a database. Always wanted to try so
here we go

Features:
* Single file
* Indexes
* Schema

Call `TheatrumOrbis.NewDB("path/to/your/data.db")` to create a database instance

All strings are UTF-8 encoded
All byte arrays are big endian
All offsets are from the start of the file and 64 bits.

Supported data types:

int64 (all ints are int64 for simplicity sake)
string (all ways utf-8 encoded, variable length)
bytearray (blob type)
double (floating point type always 64 bits)



 */

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/emilyselwood/TheatrumOrbis/schema"
	"github.com/emilyselwood/TheatrumOrbis/utils"
	"io"
	"os"
)

const VERSION int16 = 1

type DB struct {
	File io.ReadWriteSeeker
	Schema *schema.Schema
}


/*
NewDB will create a database from the given path. If it already exists it will call LoadDB otherwise CreateDB

Make sure you call `defer db.Close()` on the resulting object to close the file handle down.
 */
func NewDB(path string) (*DB, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("could not create file for db, %v", err)
		}
		return CreateDB(f)
	}

	return LoadDB(f)
}


func CreateDB(f io.ReadWriteSeeker) (*DB, error){
	var result DB

	// write header
	if err := writeHeader(f); err != nil {
		return nil, fmt.Errorf("could not create header, %v", err)
	}

	// write empty schema object


	result.File = f
	return &result, nil
}


func LoadDB(f io.ReadWriteSeeker) (*DB, error) {
	var result DB

	result.File = f
	if err := checkHeader(f); err != nil {
		return nil, fmt.Errorf("could not load db, %v", err)
	}
	return &result, nil
}

func (db *DB) Close() error {
	if db.File != nil {
		// if the file is closeable then close it.
		if f, ok := db.File.(io.Closer); ok {
			if err := f.Close(); err != nil {
				return fmt.Errorf("could not close database, %v", err)
			}
		}
		db.File = nil
	}
	return nil
}


func writeHeader(f io.WriteSeeker) error {
	if err := utils.WriteString(f, "TODB"); err != nil {
		return fmt.Errorf("could not write magic value, %v", err)
	}

	if err := utils.WriteInt16(f, VERSION); err != nil {
		return fmt.Errorf("could not write version number, %v", err)
	}

	// we will default the payload offset to just after the header
	if err := utils.WriteInt64(f, int64(14)); err != nil {
		return fmt.Errorf("could not write payload offset, %v", err)
	}



	return nil
}

func checkHeader(f io.ReadSeeker) error {
	const headerLength = 14
	header := make([]byte, headerLength)
	n, err := f.Read(header)
	if err != nil {
		return fmt.Errorf("could not read header, %v", err)
	}

	if n != headerLength {
		return fmt.Errorf("incorrect header length")
	}

	if bytes.Compare(header[0:4], []byte("TODB")) == 0 {
		return fmt.Errorf("incorrect magic value")
	}

	version := binary.BigEndian.Uint16(header[4:6])
	// if the file version is higher than our version we should not try and load it.
	// we do not want to damage the data.
	if int16(version) > VERSION {
		return fmt.Errorf("file version %v is incompatible with this version %v", version, VERSION)
	}

	payloadOffset := binary.BigEndian.Uint64(header[6:headerLength])
	// the payload must not start in the middle of the header.
	if payloadOffset < headerLength {
		return fmt.Errorf("invalid payload offset")
	}

	return nil
}
