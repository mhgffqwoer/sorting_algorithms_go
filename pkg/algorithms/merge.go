package algorithms

import (
	"time"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
)

func mergeSort(items []int, sw *int) []int {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2],sw)
    second := mergeSort(items[len(items)/2:],sw)
	*sw+=1
    return merge(first, second)
}

func merge(a []int, b []int) []int {
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}

func Testimony_merge(a *model.Array){
	start := time.Now()

	a.Array = mergeSort(a.Array, &a.Swapped)

	duration := time.Since(start)
	a.Time = duration.String()
}

