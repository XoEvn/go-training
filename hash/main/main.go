package main

var elem = []int{}

var count = int(0)

var size = int(12)

var nullKey = 0

func sum(n int) int {
	if n <= 1 {
		return 1
	}
	return sum(n-1) + n
}

func HashTable() {
	elem = make([]int, size)
	for i := 0; i < size; i++ {
		elem[i] = nullKey
	}
}

func searchHash(key int) string {
	index := key % size
	for elem[index] != key {
		index = (index + 1) % size
		if elem[index] == nullKey || index == key%size {
			return "不存在"
		}
	}
	return "存在，索引为" + string(rune(index))
}

func insertHash(key int) {
	if count == size {
		return
	}
	index := key % size
	for elem[index] != nullKey {
		index = (index + 1) % size
	}
	elem[index] = key

	count++

}
