package sortingCmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
	"github.com/mhgffqwoer/sorting_algorithms_go/pkg/algorithms"
	"github.com/mhgffqwoer/sorting_algorithms_go/pkg/generator"
	"github.com/spf13/cobra"
)

var commandSelect string
var selectionCmd = &cobra.Command{
	Use:   "selection",
    Short:  "Selection sort",
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
        case commandSelect=="random":
            generator.Random(&a)
        case commandSelect=="reverse":
            generator.Reverse(&a)
        case commandSelect=="direct":
            generator.Direct(&a)
        default:
            log.Fatal()
        }


        if a.Length<=20{
            fmt.Printf("start array: %v\n",a.Array)
            algorithms.Selection(&a)
            fmt.Printf("finish array: %v\n",a.Array)
        }else{
            algorithms.Selection(&a)
        }

        fmt.Printf("Time: %v\nSwapped: %v\n",a.Time,a.Swapped)
    },
}

func init() {
    selectionCmd.Flags().StringVarP(&commandSelect,"command","c","random","random/reverse/direct")
    rootCmd.AddCommand(selectionCmd)
}