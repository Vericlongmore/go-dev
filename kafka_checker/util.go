package kafka_checker

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func loadConfigFile(cfgPath string, ptr interface{}) error {
	if ptr == nil {
		return fmt.Errorf("ptr of type (%T) is nil", ptr)
	}
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s, err: %v", cfgPath, err)
	}
	if err := json.Unmarshal(data, ptr); err != nil {
		return fmt.Errorf("failed to unmarshal file %s, err: %v", cfgPath, err)
	}
	return nil
}

func LoadConfigFileOrPanic(cfgPath string, ptr interface{}) {
	if err := loadConfigFile(cfgPath, ptr); err != nil {
		log.WithError(err).Fatalf("failed to load config %s", cfgPath)
	}
}

func RemoveDuplicateElements(arrs []int64) []int64 {
	res := make([]int64, 0, len(arrs))
	lookup := make(map[int64]struct{})
	for _, item := range arrs {
		if _, ok := lookup[item]; !ok {
			lookup[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

func Uint8ToBytes(x uint8) []byte {
	b := make([]byte, 1)
	b[0] = x
	return b
}

func BytesToUint8(b []byte) uint8 {
	var x uint8
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &x)
	return x
}

func Uint16ToBytes(x uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, x)
	return b
}

func BytesToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func Uint32ToBytes(x uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, x)
	return b
}

func BytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func Uint64ToBytes(x uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, x)
	return b
}

func BytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}
