package synctool

import (
	"fmt"
	"sync"
)

func Map() {
	var m sync.Map

	// write
	m.Store("super", "27")
	m.Store("cat", "28")

	// read
	v, ok := m.Load("super")
	fmt.Printf("Load: v,ok = %v, %v\n", v, ok)

	// delete
	m.Delete("asd")

	// read or write
	v, ok = m.LoadOrStore("super", "18")
	fmt.Printf("LoadOrStore: v, ok = %v,%v\n", v, ok)

	// range
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Range: k, v = %v, %v\n", key, value)
		return true
	})
}
