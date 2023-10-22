package main

import "fmt"

type MyHashMap struct {
	data [100001][10]int
}

func Constructor() MyHashMap {
	var data [100001][10]int
	for i := 0; i < 100001; i++ {
		for j := 0; j < 10; j++ {
			data[i][j] = -1
		}
	}
	return MyHashMap{data: data}
}

func (this *MyHashMap) Put(key int, value int) {
	hash, index := this.hash(key)
	this.data[hash][index] = value
}

func (this *MyHashMap) Get(key int) int {
	hash, index := this.hash(key)
	return this.data[hash][index]
}

func (this *MyHashMap) Remove(key int) {
	hash, index := this.hash(key)
	this.data[hash][index] = -1
}

func (this *MyHashMap) hash(key int) (hash int, index int) {
	hash = key / 10
	index = key % 10

	return
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	data := Constructor()

	fmt.Println(data.Get(1))
	data.Put(8888, 100)
	fmt.Println(data.Get(8888))
}
