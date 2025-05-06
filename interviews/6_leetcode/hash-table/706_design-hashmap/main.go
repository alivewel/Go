package main

type MyHashMap struct {
    backets [][]Pair
}

type Pair struct {
    key, value int
}

func Constructor() MyHashMap {
    
}


func (this *MyHashMap) Put(key int, value int) {
    // определяем номер бакета
	index := key % len(this.backets)
	// определяем бакет
	backet := this.backets[index]
	// пытаемся найти ключ в бакете
	for _, pair := range backet {
		if pair.key == key {
			pair.value = value
			return
		}
	}
	this.backets[index] = append(this.backets[index], Pair{key: key, value: value})
}


func (this *MyHashMap) Get(key int) int {
    index := key % len(this.backets)
	// определяем бакет
	backet := this.backets[index]
	// пытаемся найти ключ в бакете
	for _, pair := range backet {
		if pair.key == key {
			return value
		}
	}
	return -1
}


func (this *MyHashMap) Remove(key int) {
    newBucket := []Pair{}
	for _, p := range bucket {
		if p.key != targetKey {
			newBucket = append(newBucket, p)
		}
	}
	bucket = newBucket
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */