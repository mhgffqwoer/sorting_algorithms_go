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