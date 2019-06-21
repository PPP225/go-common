package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// ToGob - go binary encoder
func ToGob(m interface{}) []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
}

// FromGob - go binary decoder
func FromGob(by []byte, m interface{}) {
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err := d.Decode(m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
}
