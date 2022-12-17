package generator

import ("math/rand"

		"github.com/mhgffqwoer/sorting_algorithms_go/model"	
)

func  Random(a *model.Array){
	for i:=0; i<a.Length; i++{
		a.Array = append(a.Array, rand.Intn(a.Length))
	}
}

func Reverse(a *model.Array){
	for i:=a.Length-1; i>=0; i--{
		a.Array = append(a.Array, i)
	}
}

func Direct(a *model.Array){
	for i:=0; i<a.Length; i++{
		a.Array = append(a.Array, i)
	}
}
