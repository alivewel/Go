package main

type MyHashMap struct {
    buckets [][]Pair
}

type Pair struct {
    key, value int
}

func Constructor() MyHashMap {
    size := 8
	return MyHashMap{
		buckets: make([][]Pair, size),
	}
}


func (this *MyHashMap) Put(key int, value int) {
    // определяем номер бакета
	index := key % len(this.buckets)
	// определяем бакет
	bucket := this.buckets[index]
	// пытаемся найти ключ в бакете
	for i, pair := range bucket {
		if pair.key == key {
			this.buckets[index][i].value = value
			return
		}
	}
	this.buckets[index] = append(this.buckets[index], Pair{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
    index := key % len(this.buckets)
	// определяем бакет
	bucket := this.buckets[index]
	// пытаемся найти ключ в бакете
	for _, pair := range bucket {
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    newBucket := []Pair{}
    index := key % len(this.buckets) // определяем бакет
	bucket := this.buckets[index] // пытаемся найти ключ в бакете
	for _, p := range bucket {
		if p.key != key {
			newBucket = append(newBucket, p)
		}
	}
	this.buckets[index] = newBucket
}


// имеется: backet должно быть: bucket

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */