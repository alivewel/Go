package main

type Node struct {
	key int
    value int
	next *Node
}

func NodeConstructor(key int, value int, next *Node) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  next,
	}
}

type MyHashMap struct {
	buckets []*Node
	size int
}

func Constructor() MyHashMap {
    size := 10001
	return MyHashMap {
		buckets: make([]*Node, size), 
		size:    size,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	current := this.buckets[index]

	for current != nil {
		if current.key == key {
			current.value = value
			return
		}
		current = current.next
	}
    
    newNode := NodeConstructor(key, value, this.buckets[index])
	this.buckets[index] = newNode
}

func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	current := this.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    index := this.hash(key)
	current := this.buckets[index]
    var previous *Node
    
    for current != nil {
        if current.key == key {
            if previous == nil {
                this.buckets[index] = current.next
            } else {
                previous.next = current.next
            }
            return 
        }
        previous = current
        current = current.next
    }
}