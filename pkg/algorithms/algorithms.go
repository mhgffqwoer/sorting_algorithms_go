package algorithms

import (
	"time"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
)

func Bubble(a *model.Array){
	start := time.Now()

	for run := range a.Array {
        for i:=1; i<len(a.Array)-run; i++{
            if a.Array[i] < a.Array[i-1]{
                a.Array[i-1], a.Array[i] = a.Array[i], a.Array[i-1]
                a.Swapped += 1
			}
		}
	}

	duration := time.Since(start)
	a.Time = duration.String()
}

func Selection(a *model.Array){
	start := time.Now()

	for i:=0; i<len(a.Array)-1; i++{
		m := i
		for j:=i+1; j<len(a.Array);j++{
			if a.Array[j]<a.Array[m]{
				m = j
			}
		}
		a.Swapped += 1
		a.Array[i], a.Array[m] = a.Array[m], a.Array[i]
	}

	duration := time.Since(start)
	a.Time = duration.String()
}