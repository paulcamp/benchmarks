package maps

func makeUnAllocatedMap(size int) {
	myMap := make(map[int]int)

	for i := 0; i < size; i++ {
		myMap[i] = 1
	}

}

func makePreAllocatedMap(size int) {

	myMap := make(map[int]int, size)

	for i := 0; i < size; i++ {
		myMap[i] = i
	}
}
