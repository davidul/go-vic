package linkedlist

import (
	"bytes"
	"encoding/binary"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMap_Add(t *testing.T) {
	log := NewLog()
	addFunc := func(key uint64, value []byte, log *AppendOnly) {
		log.Add(key, value)
		assert.Equal(t, log.Get(key), value)
	}

	t.Run("Add", func(t *testing.T) {
		addFunc(1, []byte("ABCDS"), log)
		addFunc(2, []byte("XYZ"), log)
	})
}

func TestMap_Add_1(t *testing.T) {
	now, _ := time.Now().MarshalBinary()
	log := NewLog()

	log.Add(1, now)
	log.Add(2, []byte("David"))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, math.Pi)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	log.Add(3, buf.Bytes())

	assert.Equal(t, now, log.Get(1))
	assert.Equal(t, []byte("David"), log.Get(2))

	buf2 := new(bytes.Buffer)
	err = binary.Write(buf2, binary.BigEndian, math.Pi)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, buf2.Bytes(), log.Get(3))
}

func TestMap_Types(t *testing.T) {
	log := NewLog()
	//bytes
	log.Add(1, []byte{1, 2, 3, 4})
	//ints
	bytes32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes32, 124)
	log.Add(2, bytes32)
	bytes64 := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes64, 123423234)
	log.Add(3, bytes64)
	// string/rune
	log.Add(4, []byte("Hello World"))

	assert.Equal(t, log.Get(1), []byte{1, 2, 3, 4})
	assert.Equal(t, log.Get(2), bytes32)
	assert.Equal(t, log.Get(3), bytes64)
	assert.Equal(t, log.Get(4), []byte("Hello World"))
}

// TestConcurrentWrites verifies that multiple goroutines can safely write to the log
func TestConcurrentWrites(t *testing.T) {
	log := NewLog()
	numGoroutines := 5
	writesPerGoroutine := 3

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Each goroutine writes multiple entries
	// Each entry: 8 (key) + 8 (timestamp) + 4 (length) + 2 (value) = 22 bytes
	// Total: 5 * 3 * 22 = 330 bytes (well under 1024)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < writesPerGoroutine; j++ {
				key := uint64(goroutineID*100 + j)
				value := []byte{byte(goroutineID), byte(j)}
				log.Add(key, value)
			}
		}(i)
	}

	wg.Wait()

	// Verify all entries were written correctly
	for i := 0; i < numGoroutines; i++ {
		for j := 0; j < writesPerGoroutine; j++ {
			key := uint64(i*100 + j)
			expected := []byte{byte(i), byte(j)}
			actual := log.Get(key)
			assert.Equal(t, expected, actual, "Mismatch for goroutine %d, write %d", i, j)
		}
	}
}

// TestConcurrentReads verifies that multiple goroutines can safely read from the log
func TestConcurrentReads(t *testing.T) {
	log := NewLog()

	// Prepopulate the log with small entries
	// Each entry: 22 bytes, 20 entries = 440 bytes
	numEntries := 20
	for i := uint64(0); i < uint64(numEntries); i++ {
		log.Add(i, []byte{byte(i)})
	}

	numGoroutines := 10
	readsPerGoroutine := 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Multiple goroutines reading concurrently
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < readsPerGoroutine; j++ {
				key := uint64(j % numEntries)
				value := log.Get(key)
				expected := []byte{byte(key)}
				assert.Equal(t, expected, value)
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentReadWrite verifies that reads and writes can happen concurrently
func TestConcurrentReadWrite(t *testing.T) {
	log := NewLog()

	// Prepopulate with some initial data (10 entries = 220 bytes)
	for i := uint64(0); i < 10; i++ {
		log.Add(i, []byte{byte(i)})
	}

	numReaders := 5
	numWriters := 3
	readsPerReader := 30
	writesPerWriter := 3

	var wg sync.WaitGroup
	wg.Add(numReaders + numWriters)

	// Start reader goroutines
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < readsPerReader; j++ {
				key := uint64(j % 10)
				log.Get(key) // Just verify no panic occurs
				time.Sleep(time.Microsecond)
			}
		}()
	}

	// Start writer goroutines (9 more entries = 198 bytes, total ~418 bytes)
	for i := 0; i < numWriters; i++ {
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < writesPerWriter; j++ {
				key := uint64(100 + writerID*10 + j)
				value := []byte{byte(writerID), byte(j)}
				log.Add(key, value)
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	wg.Wait()

	// Verify the written data
	for i := 0; i < numWriters; i++ {
		for j := 0; j < writesPerWriter; j++ {
			key := uint64(100 + i*10 + j)
			expected := []byte{byte(i), byte(j)}
			actual := log.Get(key)
			assert.Equal(t, expected, actual)
		}
	}
}

// TestConcurrentDuplicateKeys verifies that concurrent writes with duplicate keys are handled safely
func TestConcurrentDuplicateKeys(t *testing.T) {
	log := NewLog()
	numGoroutines := 5

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	keys := make([]uint64, numGoroutines)

	// All goroutines try to add with the same key (1)
	// 5 entries = 110 bytes
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			value := []byte{byte(goroutineID)}
			// All try to use key 1, but only the first should succeed
			// Others will get auto-generated keys
			key := log.Add(1, value)
			keys[goroutineID] = key
		}(i)
	}

	wg.Wait()

	// Verify that all keys are unique and all values are retrievable
	keySet := make(map[uint64]bool)
	for i := 0; i < numGoroutines; i++ {
		key := keys[i]
		assert.False(t, keySet[key], "Duplicate key generated: %d", key)
		keySet[key] = true

		value := log.Get(key)
		assert.NotNil(t, value)
	}
}
