package linkedlist

import (
	"encoding/binary"
	"fmt"
	"time"
)

// key -> offset
var m map[string]int = make(map[string]int)

//type Item struct {
//	key         uint8
//	keyLength   uint8
//	timestamp   time.Time
//	valueLength uint32
//	value       []byte
//}

// m map of offsets
// offset last offset
// memory the actual values
type AppendOnly struct {
	m      map[uint64]int
	offset int
	memory []uint8
}

var offset int = 0

// offsets
var memory []uint8 = make([]uint8, 1024)

func NewLog() *AppendOnly {
	only := &AppendOnly{
		m:      make(map[uint64]int),
		offset: 0,
		memory: make([]uint8, 1024),
	}

	return only
}

func (A *AppendOnly) Add(key uint64, v []byte) uint64 {
	if _, ok := A.m[key]; ok {
		key = A.GenerateKey()
	}
	// |key|timestamp|valueLength|value|
	A.m[key] = A.offset
	binary.BigEndian.PutUint64(A.memory[A.offset:], key)
	A.offset += 8
	binary.BigEndian.PutUint64(A.memory[A.offset:], uint64(time.Now().UnixNano()))
	A.offset += 8
	valLength := len(v)
	binary.BigEndian.PutUint32(A.memory[A.offset:], uint32(valLength))
	A.offset += 4
	copy(A.memory[A.offset:], v)
	A.offset += valLength

	return key
}

// |key|timestamp|valueLength|value|
func (A *AppendOnly) Get(key uint64) []uint8 {
	end := A.m[key]
	end += 8
	//key
	k := binary.BigEndian.Uint64(A.memory[:end])
	fmt.Println(k)
	start := end
	end += 8
	//timestamp
	ts := binary.BigEndian.Uint64(A.memory[start:end])
	fmt.Println(ts)
	start = end
	end += 4
	//value length
	vLength := binary.BigEndian.Uint32(A.memory[start:end])
	e := end + int(vLength)

	//value
	value := A.memory[end:e]
	return value
}

func (A *AppendOnly) GetOffset(key uint64) int {
	return A.m[key]
}

func (A *AppendOnly) GetMemory() []uint8 {
	return A.memory
}

func (A *AppendOnly) GenerateKey() uint64 {
	return uint64(time.Now().UnixNano())
}
