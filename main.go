package main
import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func randFloats(min, max float64, n int) []float64 {
    res := make([]float64, n)
    for i := range res {
        res[i] = min + rand.Float64() * (max - min)
    }
    return res
}

// swap - swap elements at i1 and i2 in arr
func swap (arr []float64, i1, i2 int) {
	swapF := reflect.Swapper(arr)
	swapF(i1, i2)
	return
} 

func partition(arr []float64, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high - 1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i + 1, high)
	return i + 1
}

// QuickSort sort `arr` with `low` (starting index) and `high` (Ending index) using quick sort algorithm
// I use quick sort to enhance worst case of time complexity
func QuickSort(arr []float64, low int, high int) {
	if (low < high) {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi - 1)
		QuickSort(arr, pi + 1, high)
	}
	return
}

// BucketSort sort `arr` with `s` (size) using bucket sort
func BucketSort(arr []float64, s int) {
	// initializing new buckets
	buckets := make([][]float64, s)
	// Assign array elements in different bucket
	for i := 0; i < s; i++ {
		bi := int(float64(s) * arr[i]) // Index in bucket
		buckets[bi] = append(buckets[bi], arr[i])
	}
	// sort all buckets
	for _, b := range buckets {
		QuickSort(b, 0, len(b) - 1)
	}
	// Concatenate all buckets into arr[]
	index := 0
	for _, b := range buckets {
		for _, el := range b {
			arr[index] = el
			index++
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := randFloats(0, 1, 1000000)
	start := time.Now()
	BucketSort(arr, len(arr))
	elapsed := time.Since(start)
	fmt.Printf("Bucket sort took: %d ms\n", (elapsed / time.Millisecond))
}