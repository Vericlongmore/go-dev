package main

import (
	"fmt"
	"go-dev/Week09/binary"
)

func main() {
	data := encoder("Hello")
	decoder(data)
}

func decoder(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data len < 16.")
		return
	}

	packetLen := binary.BigEndian.Int32(data[:4])
	fmt.Printf("packetLen:%v\n", packetLen)

	headerLen := binary.BigEndian.Int16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Int16(data[6:8])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Int32(data[8:12])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Int32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}

func encoder(body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)

	binary.BigEndian.PutInt32(ret[:4], int32(packetLen))
	binary.BigEndian.PutInt16(ret[4:6], int16(headerLen))

	version := 5
	binary.BigEndian.PutInt16(ret[6:8], int16(version))
	operation := 6
	binary.BigEndian.PutInt32(ret[8:12], int32(operation))
	sequence := 7
	binary.BigEndian.PutInt32(ret[12:16], int32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)

	return ret
}
