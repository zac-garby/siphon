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

	keyType Type
	valType Type
}

// NewHashmap makes a new empty Hashmap
func NewHashmap(keyType, valType Type) *Hashmap {
	return &Hashmap{
		data:    make(map[string]Item),
		keys:    make(map[string]Item),
		keyType: keyType,
		valType: valType,
	}
}

// Type returns the type of the Item
func (h *Hashmap) Type() Type {
	return &HashmapType{
		KeyType: h.keyType,
		ValType: h.valType,
	}
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
	if !key.Type().Equals(h.keyType) {
		return nil, StatusType
	}

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
	if !key.Type().Equals(h.keyType) {
		return StatusType
	}

	if !to.Type().Equals(h.valType) {
		return StatusType
	}

	hash, err := structhash.Hash(key, 1)
	if err != nil {
		return StatusError
	}

	h.data[hash] = to
	h.keys[hash] = key

	return StatusOK
}

// GetField gets the given field from the hashmap
func (h *Hashmap) GetField(key string) (result Item, status string) {
	return h.GetKey(NewString(key))
}

// SetField sets the given field in the hashmap to a value
func (h *Hashmap) SetField(key string, to Item) (status string) {
	return h.SetKey(NewString(key), to)
}
