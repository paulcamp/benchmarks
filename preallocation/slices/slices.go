package slices

func makeUnAllocatedSlice(size int) {
	var mySlice []int

	for i := 0; i < size; i++ {
		mySlice = append(mySlice, i)
	}
}

func makeSliceWithPreAllocatedCapacity(size int) {
	mySlice := make([]int, 0, size) //type, len, cap

	for i := 0; i < size; i++ {
		mySlice = append(mySlice, i)
	}
}

func makeSlicewithPreAllocatedLength(size int) {

	mySlice := make([]int, size) //type,len (cap will be inferred from len)

	for i := 0; i < size; i++ {
		mySlice[i] = i
	}
}
