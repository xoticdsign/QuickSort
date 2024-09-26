// Quick Sort in Golang

/*

- What is Quick Sort?

Quick Sort is a highly efficient sorting algorithm, known for its divide-and-conquer approach. Invented by Tony Hoare in 1959, the algorithm works by selecting
a “pivot” element from the array and partitioning the other elements into two sub-arrays based on whether they are less than or greater than the pivot. It
recursively applies this process until the array is fully sorted.

Quick Sort worst-case complexity:
	•   Time complexity of O(n2)
	•   Space complexity of O(log n)

- Key Concepts of Quick Sort:

	1.	Pivot Selection: A pivot element is chosen, often the last element in the current sub-array.
	2.	Partitioning: The array is rearranged so that elements smaller than the pivot are on the left, and elements larger than the pivot are on the right.
	3.	Recursion: Quick Sort is recursively applied to the sub-arrays on both sides of the pivot.
	4.	Base Case: The recursion stops when sub-arrays have zero or one element, which means they are sorted.

- Steps in the Algorithm

	1.	Partitioning: The array is partitioned using a pivot. Elements smaller than the pivot are moved to its left, and elements larger are moved to the right.
	2.	Recursion: The sub-arrays on either side of the pivot are recursively sorted.
	3.	Final Array: After all partitions and recursive calls, the array is fully sorted.

*/

// How to implement Quick Sort algorithm step by step

/*

    1.	QuickSort Function:
	•	The QuickSort function is responsible for the recursive partitioning of the array.
	•	If low is less than high (meaning there are still elements to sort), it calls the QuickPartition function to perform the partitioning.
	•	After partitioning, it recursively applies QuickSort to the sub-arrays left and right of the pivot.
	2.	QuickPartition Function:
	•	This function rearranges the elements of the array such that all elements smaller than the pivot are moved to its left, and all elements greater than
	    the pivot are moved to its right.
	•	The pivot element is selected as the last element (array[high]), but you could choose a different strategy (such as the first element, random element,
	    etc.).
	•	The function uses two indices (i and j). The i index keeps track of the boundary between elements smaller than the pivot and elements not yet compared.
	    The j index traverses the array to find elements smaller than the pivot, and these elements are swapped to the left of the pivot.
	3.	Swapping:
	•	During partitioning, elements are swapped to ensure that elements smaller than the pivot end up on the left, while elements larger than the pivot end
	    up on the right.

Example Walkthrough:

Let’s take the array: [45, 34, -34, -1, 21, 21, 8, 21, 10]

	1.	Initial Call:
	•	We call QuickSort(array, 0, len(array)-1), which sorts the entire array.
	2.	First Partition:
	•	Pivot: 10
	•	After the first partition, the array becomes:
        [-34, -1, 8, 10, 21, 21, 21, 45, 34]
	•	The pivot is now at index 3.
	3.	Recursive Sorting:
	•	The algorithm recursively sorts the left part [-34, -1, 8] and the right part [21, 21, 21, 45, 34].
	4.	Further Partitioning:
	•	The left part [-34, -1, 8] is already sorted.
	•	The right part is partitioned again, and the process continues until the entire array is sorted.
	5.	Final Sorted Array:
	•	After all partitions and recursive calls, the final sorted array is:
        [-34, -1, 8, 10, 21, 21, 21, 34, 45]

	•	Quick Sort is an efficient sorting algorithm that uses a pivot to partition the array and recursively sorts the sub-arrays.
	•	It has an average time complexity of O(n log n), which makes it faster than some other sorting algorithms, like Bubble Sort or Selection Sort, especially for large datasets.
	•	The key operations are selecting a pivot, partitioning the array, and recursively sorting the left and right sub-arrays.

*/

// Raw Quick Sort algorithm

package main

import "fmt"

func main() {
	array := []int{45, 34, -34, -1, 21, 21, 8, 21, 10}

	sortedArray := QuickSort(array, 0, len(array)-1)
	fmt.Println(sortedArray)
}

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

// Now let's see how the process looks in live

/*

package main

import (
	"fmt"
)

func main() {
	array := []int{45, 34, -34, -1, 21, 21, 8, 21, 10}

	fmt.Printf("\n1. Unsorted array: %v\n", array)
	fmt.Printf("\n2. Sort is started\n")

	sortedArray := QuickSort(array, 0, len(array)-1)

	fmt.Printf("\n3. Sort completed\n")
	fmt.Printf("\n4. Sorted array: %v\n", sortedArray)
}

func QuickSort(array []int, low int, high int) []int {
	fmt.Printf("\nChecking, if %v is available for sorting", array)

	if low >= high {
		fmt.Printf("\nNot available\n")

		return array
	} else {
		fmt.Printf("\nStarting partition\n")

		p := QuickPartition(array, low, high)
		QuickSort(array, low, p-1)
		QuickSort(array, p+1, high)

		return array
	}
}

func QuickPartition(array []int, low int, high int) int {
	pivot := array[high]
	x := low - 1

	for y := low; y < high; y++ {
		fmt.Printf("\npivot = %v\nx = %v\nlow = %v\nhigh = %v\n", pivot, x, low, high)
		fmt.Printf("\nfor y := low (%v); y < high; y++ {", y)
		fmt.Printf("\n   Comparing array[%v] = %v and pivot = %v {", y, array[y], pivot)

		if array[y] < pivot {
			fmt.Printf("\n      array[%v] < pivot", y)
			fmt.Printf("\n      Increasing x by 1")

			x++
			array[x], array[y] = array[y], array[x]

			fmt.Printf("\n      Replacing array[%v] with array[%v]", x, y)
			fmt.Printf("\n      Result: %v", array)
			fmt.Printf("\n   }")
			fmt.Printf("\n}\n")
		} else {
			fmt.Printf("\n      array[%v] > pivot", y)
			fmt.Printf("\n      Number is in the right place")
			fmt.Printf("\n      Result: %v", array)
			fmt.Printf("\n   }")
			fmt.Printf("\n}\n")
		}
	}

	fmt.Printf("\nReplacing array[x+1] with array[%v]\n", high)

	array[x+1], array[high] = array[high], array[x+1]

	return x + 1
}

*/
