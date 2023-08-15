package cache

import (
	"errors"
)

var (
	errSmallCapacity = errors.New("минимальная емкость кэша - 1")
	errNotFound      = errors.New("значение не найдено")
	errEmptyCache    = errors.New("кэш пуст")
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, bool)
}

type LRUCache struct {
	Values map[string]string
	Times  TimeQueue
	Cap    int
	Len    int
}

type TimeQueue struct {
	start *Node
	end   *Node
}

type Node struct {
	val  string
	next *Node
}

func CreateLRUCache(cap int) (*LRUCache, error) {
	if cap < 1 {
		return nil, errSmallCapacity
	}

	node := &Node{}
	return &LRUCache{
		Values: make(map[string]string),
		Times:  TimeQueue{node, node},
		Cap:    cap,
		Len:    0,
	}, nil
}

func (c *LRUCache) Set(key string, val string) error {
	if c.Len == c.Cap {

		lastKey, err := c.Times.RemoveLastTime()

		if err != nil {
			return err
		}

		delete(c.Values, lastKey)

	} else {
		c.Len++
	}

	c.Values[key] = val
	c.Times.Add(key)
	return nil
}

func (t *TimeQueue) Add(key string) {
	newNode := &Node{val: key}
	t.end.next = newNode
	t.end = t.end.next
}

func (c *LRUCache) Get(key string) (string, bool) {
	if val, ok := c.Values[key]; ok {

		if err := c.Times.Update(key); err != nil {
			c.Times.Add(key)
		}

		return val, true
	}
	return "", false
}

func (t *TimeQueue) Update(key string) error {
	prevNode := t.start

	for prevNode.next != nil {

		if prevNode.next.val == key {
			selectedNode := prevNode.next
			prevNode.next = selectedNode.next
			t.Add(key)
			return nil
		}

		prevNode = prevNode.next
	}

	return errNotFound
}

func (t *TimeQueue) RemoveLastTime() (string, error) {
	if t.start.next != nil {
		key := t.start.next.val
		t.start.next = t.start.next.next
		return key, nil
	}
	return "", errEmptyCache
}
