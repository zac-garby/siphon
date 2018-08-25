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

// JSON returns a JSON representation of an item
func (h *Hashmap) JSON() string {
	str := &strings.Builder{}

	str.WriteByte('{')

	i := 0
	for hash, val := range h.data {
		key := h.keys[hash]
		if i > 0 {
			str.WriteString(", ")
		}

		if key.Type().Equals(&StringType{}) {
			str.WriteString(key.JSON())
		} else {
			str.WriteString("\"" + key.String() + "\"")
		}

		str.WriteString(": ")
		str.WriteString(val.JSON())

		i++
	}

	str.WriteByte('}')

	return str.String()
}

// Set sets the value of the item to the given value
func (h *Hashmap) Set(val interface{}) (err error) {
	if !h.keyType.Equals(&StringType{}) {
		return newError(ErrNOOP, "set can only be used on <string:any> hashmaps, due to JSON syntax")
	}

	hval, ok := val.(map[string]interface{})
	if !ok {
		return newError(ErrType, "expected a hashmap value")
	}

	h.data = make(map[string]Item, len(hval))
	h.keys = make(map[string]Item, len(hval))

	for k, v := range hval {
		newVal := MakeZeroValue(h.valType)
		if err := newVal.Set(v); err != nil {
			return err
		}

		if err := h.SetKey(&String{value: k}, newVal); err != nil {
			return err
		}
	}

	return nil
}

// GetKey gets the given key from the hashmap
func (h *Hashmap) GetKey(key Item) (result Item, err error) {
	if !key.Type().Equals(h.keyType) {
		return nil, newError(
			ErrType, "hashmap key type is %s, but a key of type %s was requested",
			h.keyType,
			key.Type(),
		)
	}

	hash, err := structhash.Hash(key, 1)
	if err != nil {
		return nil, newError(ErrUnknown, "could not hash a value for some reason")
	}

	val, ok := h.data[hash]
	if !ok {
		return nil, newError(ErrIndex, "key %s does not exist", key)
	}

	return val, nil
}

// SetKey sets the given key in the hashmap to a value
func (h *Hashmap) SetKey(key Item, to Item) (err error) {
	if !key.Type().Equals(h.keyType) {
		return newError(
			ErrType, "hashmap key type is %s, so a key of type %s cannot be assigned",
			h.keyType,
			key.Type(),
		)
	}

	if !to.Type().Equals(h.valType) {
		return newError(
			ErrType, "hashmap value type is %s, so a value of type %s cannot be assigned",
			h.valType,
			to.Type(),
		)
	}

	hash, err := structhash.Hash(key, 1)
	if err != nil {
		return newError(ErrIndex, "key %s does not exist", key)
	}

	h.data[hash] = to
	h.keys[hash] = key

	return nil
}

// SetKeyJSON sets the given key in the hashmap to a value, where the key and value
// are encoded in JSON
func (h *Hashmap) SetKeyJSON(keyJSON interface{}, toJSON interface{}) (err error) {
	var (
		key = MakeZeroValue(h.keyType)
		val = MakeZeroValue(h.valType)
	)

	if err := key.Set(keyJSON); err != nil {
		return err
	}

	if err := val.Set(toJSON); err != nil {
		return err
	}

	return h.SetKey(key, val)
}

// GetField gets the given field from the hashmap
func (h *Hashmap) GetField(key string) (result Item, err error) {
	return h.GetKey(NewString(key))
}

// SetField sets the given field in the hashmap to a value
func (h *Hashmap) SetField(key string, to Item) (err error) {
	return h.SetKey(NewString(key), to)
}

// Filter filters the hashmap, returning a new hashmap where only the
// filtered key:val pairs are present.
func (h *Hashmap) Filter(field string, kind Comparison, other Item) (result Item, err error) {
	result = &Hashmap{
		keyType: h.keyType,
		valType: h.valType,
		data:    make(map[string]Item, len(h.data)/2), // initialise with capacity as len()/2
		keys:    make(map[string]Item, len(h.keys)/2),
	}

	for hash, val := range h.data {
		var (
			key       = h.keys[hash]
			predicate bool
		)

		if field == "" {
			pred, err := val.Compare(kind, other)
			if err != nil {
				return nil, err
			}

			predicate = pred
		} else {
			fval, err := val.GetField(field)
			if err != nil {
				return nil, err
			}

			pred, err := fval.Compare(kind, other)
			if err != nil {
				return nil, err
			}

			predicate = pred
		}

		if predicate {
			result.SetKey(key, val)
		}
	}

	return result, nil
}
