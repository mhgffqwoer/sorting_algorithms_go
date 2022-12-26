package algorithms

import (
	"time"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
)

func quick(ar []int) int{
	if len(ar) <= 1 {
		return 1
	}

	split := partition(ar)

	return quick(ar[split:])+quick(ar[:split])
}

func partition(a []int) int {
	pivot := a[len(a)/2]

	left := 0
	right := len(a) - 1

	for {
		for ; a[left] < pivot; left++ {
		}

		for ; a[right] > pivot; right-- {
		}
        if a[left]==a[right]{
            left++
        }
		if left >= right {
			return right
		}

		a[left], a[right] = a[right], a[left]
	}
}

func Testimony_quick(a *model.Array){
	start := time.Now()

	a.Swapped = quick(a.Array)

	duration := time.Since(start)
	a.Time = duration.String()
}