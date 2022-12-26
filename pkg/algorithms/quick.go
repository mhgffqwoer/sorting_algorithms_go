package algorithms

import (
	"time"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
)

func quick(ar []int, sw *int){
	if len(ar) <= 1 {
		return
	}

	split,sw1 := partition(ar, 0 )
    *sw += sw1


	quick(ar[split:],sw)
    quick(ar[:split],sw)
}

func partition(a []int, sw int) (int, int){
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
			return right, sw
		}

        sw+=1
		a[left], a[right] = a[right], a[left]
	}
}

func Testimony_quick(a *model.Array){
	start := time.Now()

	quick(a.Array, &a.Swapped)

	duration := time.Since(start)
	a.Time = duration.String()
}