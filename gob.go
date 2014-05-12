// gob
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func encode(value []int32) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		panic("encode error!")
	}
	return buf.Bytes()
}

func decode(value []byte) (data []int32) {
	buf := bytes.NewReader(value)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&data)
	if err != nil {
		panic("decode error!")
	}
	return
}

func main() {
	fmt.Println("Hello World!")

	v1 := []int32{1, 2, 3, 4, 5, 6, 7}
	v2 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

	b1 := encode(v1)
	b2 := encode(v2)

	fmt.Println(b1, len(b1))
	fmt.Println(b2, len(b2))

	d1 := decode(b1)
	d2 := decode(b2)

	fmt.Println(d1)
	fmt.Println(d2)

	b3 := append(b1, b2...)
	d3 := decode(b3)
	fmt.Println(d3)
}
