package db

import (
	"strings"

	"github.com/cnf/structhash"
)

// A Hashmap maps keys to values and enables O(1) lookup complexity.
type Hashmap struct {
	*itemDefaults

	// data is stored as a map from strings to Items. The string keys
	// are the digests of the actual keys.
	data map[string]Item

	// keys stores the original keys which generated the hashes.
	keys map[string]Item
}

// NewHashmap makes a new empty Hashmap
func NewHashmap() *Hashmap {
	return &Hashmap{
		data: make(map[string]Item),
		keys: make(map[string]Item),
	}
}

// Type returns the type of the Item
func (h *Hashmap) Type() string {
	return TypeHashmap
}

func (h *Hashmap) String() string {
	str := &strings.Builder{}

	str.WriteByte('<')

	i := 0
	for hash, val := range h.data {
		key := h.keys[hash]
		if i > 0 {
			str.WriteString(", ")
		}

		if i > 10 {
			str.WriteString("...")
			break
		}

		str.WriteString(key.String())
		str.WriteString(": ")
		str.WriteString(val.String())

		i++
	}

	str.WriteByte('>')

	return str.String()
}

// GetKey gets the given key from the hashmap
func (h *Hashmap) GetKey(key Item) (result Item, status string) {
	hash, err := structhash.Hash(key, 1)
	if err != nil {
		return nil, StatusError
	}

	val, ok := h.data[hash]
	if !ok {
		return nil, StatusIndex
	}

	return val, StatusOK
}

// SetKey sets the given key in the hashmap to a value
func (h *Hashmap) SetKey(key Item, to Item) (status string) {
	hash, err := structhash.Hash(key, 1)
	if err != nil {
		return StatusError
	}

	h.data[hash] = to
	h.keys[hash] = key

	return StatusOK
}
