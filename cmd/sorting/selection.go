package sorting

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
            break
        case commandSelect=="reverse":
            generator.Reverse(&a)
            break
        case commandSelect=="direct":
            generator.Direct(&a)
            break
        default:
            log.Fatal()
        }


        algorithms.Selection(&a)

		fmt.Println(a.Array)
        fmt.Println(a.Time,a.Swapped)
    },
}

func init() {
    selectionCmd.Flags().StringVarP(&commandSelect,"command","c","random","random/reverse/direct")
    rootCmd.AddCommand(selectionCmd)
}