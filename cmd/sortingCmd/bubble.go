package sortingCmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
    "github.com/mhgffqwoer/sorting_algorithms_go/pkg/generator"
    "github.com/mhgffqwoer/sorting_algorithms_go/pkg/algorithms"
	"github.com/spf13/cobra"
)

var commandBubble string
var bubbleCmd = &cobra.Command{
    Use:   "bubble",
    Short:  "Bubble sort",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {

        length, err := strconv.Atoi(args[0])
	    if err != nil {
		    log.Fatal(err)
	    }

        a := model.Array{
            Array: []int{},
            Length: length,
            Time: "",
            Swapped: 0,
        }

        switch {
        case commandBubble=="random":
            generator.Random(&a)
        case commandBubble=="reverse":
            generator.Reverse(&a)
        case commandBubble=="direct":
            generator.Direct(&a)
        default:
            log.Fatal()
        }

        
        if a.Length<=20{
            fmt.Printf("start array: %v\n",a.Array)
            algorithms.Bubble(&a)
            fmt.Printf("finish array: %v\n",a.Array)
        }else{
            algorithms.Bubble(&a)
        }

        fmt.Printf("Time: %v\nSwapped: %v\n",a.Time,a.Swapped)
    },
}

func init() {
    bubbleCmd.Flags().StringVarP(&commandBubble,"command","c","random","random/reverse/direct")
    rootCmd.AddCommand(bubbleCmd)
}