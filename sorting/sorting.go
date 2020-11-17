package sorting

/*
	BubbleSort algorith simply swaps nearest
	elements if they at at wrong position
*/
func BubbleSort(numbers *[]int) {
	for i := 0; i < len(*numbers); i++ {
		for j := 1; j < len(*numbers)-i-1; j++ {
			if (*numbers)[j] > (*numbers)[j+1] {
				(*numbers)[j], (*numbers)[j+1] = (*numbers)[j+1], (*numbers)[j]
			}
		}
	}
}

/*
	CocktailSort algorith it is a variation of
	BubbleSort algorith, but despite BubbleSort
	slice traversing appears in both directions
*/
func CocktailSort(numbers *[]int) {
	for {
		swapped := false
		for i := 0; i < len(*numbers)-1; i++ {
			if (*numbers)[i] > (*numbers)[i+1] {
				(*numbers)[i], (*numbers)[i+1] = (*numbers)[i+1], (*numbers)[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
		swapped = false
		for j := len(*numbers) - 2; j >= 0; j-- {
			if (*numbers)[j] > (*numbers)[j+1] {
				(*numbers)[j], (*numbers)[j+1] = (*numbers)[j+1], (*numbers)[j]
			}
		}
		if !swapped {
			return
		}
	}
}

/*
	InsertionSort algorith is is algorithm that
	virtually splits slice into 2 parts, sorted and
	unsorted and place value from unsorted part by
	comparing with values from sorted part
*/
func InsertionSort(numbers *[]int) {
	for i := 1; i < len(*numbers); i++ {
		j := i
		for j > 0 {
			if (*numbers)[j] < (*numbers)[j-1] {
				(*numbers)[j], (*numbers)[j-1] = (*numbers)[j-1], (*numbers)[j]
			}
			j--
		}
	}
}

/*
	SelectionSort algorith finds the minimum
	element over each iteration and then place it
	at the beginning. There are 2 parts in given
	slice: sorted and unsurted. Over each iteration
	element from unsorted array is picked and moved
	to sorted part
*/
func SelectionSort(numbers *[]int) {
	for i := 0; i < len(*numbers)-1; i++ {
		for j := i + 1; j < len(*numbers); j++ {
			if (*numbers)[j] < (*numbers)[i] {
				(*numbers)[j], (*numbers)[i] = (*numbers)[i], (*numbers)[j]
			}
		}
	}
}

/*
	MergeSort is a Divide adn Conquer algorith,
	which recursevile divides slice into two parts,
	until one element will left in slice, sort them
	and then merges sorted halves. Implementation includes
	2 functions: Merge and MergeSort
*/
func MergeSort(numbers *[]int) {
	RMergeSort(numbers, 0, len(*numbers)-1)
}

/*
	RMergeSort function call itself for left and
	right hals and then merge those halfs
*/
func RMergeSort(numbers *[]int, start int, end int) {
	if start < end {
		middle := (start + end) / 2
		RMergeSort(numbers, start, middle)
		RMergeSort(numbers, middle+1, end)
		Merge(numbers, start, middle, end)
	}
}

/*
	Merge fucntion merges halfs by comparing them and
	putting at the right place at sorting slice
*/
func Merge(numbers *[]int, start int, middle int, end int) {
	sizeLeft := middle - start + 1
	sizeRight := end - middle
	var leftSlice, rightSlice []int
	for i := 0; i < sizeLeft; i++ {
		leftSlice = append(leftSlice, (*numbers)[start+i])
	}
	for i := 0; i < sizeRight; i++ {
		rightSlice = append(rightSlice, (*numbers)[middle+i+1])
	}
	i := 0
	j := 0
	k := start
	for i < sizeLeft && j < sizeRight {
		if leftSlice[i] <= rightSlice[j] {
			(*numbers)[k] = leftSlice[i]
			i++
		} else {
			(*numbers)[k] = rightSlice[j]
			j++
		}
		k++
	}
	for i < sizeLeft {
		(*numbers)[k] = leftSlice[i]
		i++
		k++
	}
	for j < sizeRight {
		(*numbers)[k] = rightSlice[j]
		j++
		k++
	}
}

/*
	Quick sort is a Divide and Conquer algorith, that
	pick pivot element, make a partion over pivot, put
	elements smaller that pivot to the left side from
	pivot, and elements greated than pivot to the
	right side, and then calls itself to the left and
	right halfs
*/
func QuickSort(numbers *[]int) {
	Sort(&(*numbers), 0, len(*numbers)-1)
}

/*
	For this function pivot element is the last element of
	slice
*/
func Sort(numbers *[]int, start, end int) {
	if (end - start) < 1 {
		return
	}
	pivot := (*numbers)[end]
	breakIndex := start

	for i := start; i < end; i++ {
		if (*numbers)[i] < pivot {
			(*numbers)[i], (*numbers)[breakIndex] = (*numbers)[breakIndex], (*numbers)[i]
			breakIndex++
		}
	}
	(*numbers)[end] = (*numbers)[breakIndex]
	(*numbers)[breakIndex] = pivot

	Sort(&(*numbers), start, breakIndex-1)
	Sort(&(*numbers), breakIndex+1, end)
}
