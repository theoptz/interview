package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
)

const defaultCap = 20

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	cache    map[string]*list.Element
	list     *list.List
	capacity int

	mu sync.Mutex
}

func NewLRU(capacity int) *LRU {
	if capacity <= 0 {
		capacity = defaultCap
	}

	return &LRU{
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		capacity: capacity,
		mu:       sync.Mutex{},
	}
}

func (l *LRU) Set(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	el, exists := l.cache[key]
	if exists {
		l.list.MoveToFront(el)
		el.Value.(*Item).Value = value
		return
	}

	if l.list.Len() >= l.capacity {
		l.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	newEl := l.list.PushFront(item)
	l.cache[key] = newEl
}

func (l *LRU) Get(key string) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	el, exists := l.cache[key]
	if !exists {
		return nil
	}

	l.list.MoveToFront(el)
	return el.Value.(*Item).Value
}

func (l *LRU) Print() {
	// for debug purpose
	l.mu.Lock()
	defer l.mu.Unlock()

	next := l.list.Front()
	for next != nil {
		el := next.Value.(*Item)
		fmt.Println(el.Key, el.Value)
		next = next.Next()
	}
}

func (l *LRU) purge() {
	for l.list.Len() >= l.capacity {
		el := l.list.Back()
		if el != nil {
			key := el.Value.(*Item).Key
			l.list.Remove(el)
			delete(l.cache, key)
		}
	}
}

func main() {
	cache := NewLRU(5)

	for i := 0; i < 30; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	cache.Print() // from 29 to 25
	fmt.Println()

	cache.Set("foo", 100)
	fmt.Println(cache.Get("bar")) // nil (not exist)
	fmt.Println(cache.Get("25"))  // nil (purged)
	fmt.Println(cache.Get("26"))  // 26
	fmt.Println()

	cache.Print() // 26 26, foo 100, 29 29, 28 28, 27 27
}
