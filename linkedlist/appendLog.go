package linkedlist

import (
	"encoding/binary"
	"sync"
	"time"
)

// m map of offsets
// offset last offset
// memory the actual values
type AppendOnly struct {
	mu     sync.RWMutex   // Protects concurrent access to m, offset, and memory
	m      map[uint64]int // Maps keys to their offset positions in memory
	offset int            // Current write position in the memory buffer
	memory []uint8        // Byte array that stores all the data
}

func NewLog() *AppendOnly {
	only := &AppendOnly{
		m:      make(map[uint64]int),
		offset: 0,
		memory: make([]uint8, 1024),
	}

	return only
}

// |key (8 bytes)|timestamp (8 bytes)|valueLength (4 bytes)|value (variable)|
// Key: 8-byte unique identifier (uint64)
// Timestamp: 8-byte Unix timestamp in nanoseconds
// Value Length: 4-byte integer indicating how many bytes the value occupies
// Value: Variable-length byte array containing the actual data
func (A *AppendOnly) Add(key uint64, v []byte) uint64 {
	A.mu.Lock()
	defer A.mu.Unlock()

	// key must be unique, if it already exists, generate a new one
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
	A.mu.RLock()
	defer A.mu.RUnlock()

	start := A.m[key]
	end := start + 8
	//key
	//k := binary.BigEndian.Uint64(A.memory[start:end])
	start = end
	end += 8
	//timestamp
	//ts := binary.BigEndian.Uint64(A.memory[start:end])
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
	A.mu.RLock()
	defer A.mu.RUnlock()

	return A.m[key]
}

func (A *AppendOnly) GetMemory() []uint8 {
	A.mu.RLock()
	defer A.mu.RUnlock()

	return A.memory
}

func (A *AppendOnly) GenerateKey() uint64 {
	return uint64(time.Now().UnixNano())
}
