package algorithms

import (
	"time"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
)

func Insertion(a *model.Array){
	start := time.Now()

	for i := range a.Array{
		x := a.Array[i]
		j := i
		for ; j >= 1 && a.Array[j-1] > x; j-- {
			a.Swapped+=1
			a.Array[j] = a.Array[j-1]
		}
		a.Array[j] = x
	}

	duration := time.Since(start)
	a.Time = duration.String()
}