
# Quick Sort in Golang

Quick Sort is a highly efficient sorting algorithm, known for its divide-and-conquer approach. Invented by Tony Hoare in 1959, the algorithm works by selecting a “pivot” element from the array and partitioning the other elements into two sub-arrays based on whether they are less than or greater than the pivot. It recursively applies this process until the array is fully sorted.

- Time complexity of O(n2)
- Space complexity of O(log n)

Please consider diving into ["quick_sort.go"](https://github.com/xoticdsign/QuickSort/blob/xoti%24/quick_sort.go) file for more explanation with examples and vizualization!

### Glance of the algorithm

Let's see how the algorithm looks:

```go

func QuickSort(array []int, low int, high int) []int {
	if low >= high {
		return array
	}

	p := QuickPartition(array, low, high)
	QuickSort(array, low, p-1)
	QuickSort(array, p+1, high)

	return array
}

func QuickPartition(array []int, low int, high int) int {
	pivot := array[high]
	x := low - 1

	for y := low; y < high; y++ {
		if array[y] < pivot {
			x++
			array[x], array[y] = array[y], array[x]
		}
	}

	array[x+1], array[high] = array[high], array[x+1]

	return x + 1
}

```

Key Concepts of Quick Sort:

1. Pivot Selection: A pivot element is chosen, often the last element in the current sub-array.

2. Partitioning: The array is rearranged so that elements smaller than the pivot are on the left, and elements larger than the pivot are on the right.

3. Recursion: Quick Sort is recursively applied to the sub-arrays on both sides of the pivot.

4. Base Case: The recursion stops when sub-arrays have zero or one element, which means they are sorted.

### License

[MIT](https://choosealicense.com/licenses/mit/)

