package main

import (
	"strconv"
)

type hashmap struct {
	list     []*hashElem
	capacity int
	hashFn   func(interface{}) int
	size     int
}

type hashElem struct {
	key   interface{}
	value interface{}
}

func defaultHashFn(key interface{}) int {
	var strKey string
	// convert to string
	switch key.(type) {
	case string:
		strKey = key.(string)
	case int:
		strKey = strconv.Itoa(key.(int))
	case bool:
		strKey = strconv.FormatBool(key.(bool))
	case float64:
		strKey = strconv.FormatFloat(key.(float64), 'E', -1, 32)
	default:
	}

	bytesKey := []byte(strKey)

	hashSum := 0
	for i := 0; i < len(bytesKey); i++ {
		hashSum += int(bytesKey[i])
	}

	return hashSum
}

func NewHashmap(n int) *hashmap {
	h := &hashmap{
		list:     make([]*hashElem, n),
		capacity: n,
		hashFn:   defaultHashFn,
		size:     0,
	}

	return h
}

// SetHashFn ...
func (h *hashmap) SetHashFn(hashFn func(interface{}) int) {
	h.hashFn = hashFn
}

// Add ...
func (h *hashmap) Add(key interface{}, value interface{}) {

	e := &hashElem{
		key:   key,
		value: value,
	}

	hashCode := h.hashFn(key)
	pos := hashCode % h.capacity

	i := pos
	for ; i < len(h.list) && h.list[i] != nil; i++ {
		// 这没法处理冲突呀。
		if h.list[i].key == key {
			return
		}
	}

	if i == len(h.list) {
		h.list = append(h.list, e)
	} else {
		h.list[i] = e
	}

	h.size++
}

// Get ...
func (h *hashmap) Get(key interface{}) interface{} {

	hashCode := h.hashFn(key)
	pos := hashCode % h.capacity

	if pos >= len(h.list) {
		return nil
	}

	for i := pos; i < len(h.list); i++ {
		if h.list[i] != nil {
			if h.list[i].key == key {
				return h.list[i].value
			}
		}
	}

	return nil
}

// Remove ...
func (h *hashmap) Remove(key interface{}) interface{} {
	hashCode := h.hashFn(key)
	pos := hashCode % h.capacity

	if pos >= len(h.list) {
		return nil
	}

	for i := pos; i < len(h.list); i++ {
		if h.list[i].key == key {
			tmp := h.list[i].value
			h.list[i] = nil
			h.size--
			return tmp
		}
	}

	return nil
}

// RemoveAll ...
func (h *hashmap) RemoveAll() {
	h.list = make([]*hashElem, h.capacity)
}

// Size ...
func (h *hashmap) Size() int {
	return h.size
}
