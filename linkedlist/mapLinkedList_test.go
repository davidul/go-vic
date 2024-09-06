package linkedlist

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
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

	pi := log.Get(3)
	fmt.Println(pi)
	b := new(bytes.Buffer)
	n, _ := b.Read(pi)
	fmt.Println(n)
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
}
