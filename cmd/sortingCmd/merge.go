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

var commandMerge string
var mergeCmd = &cobra.Command{
    Use:   "merge",
    Short:  "Merge sort",
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
        case commandMerge=="random":
            generator.Random(&a)
        case commandMerge=="reverse":
            generator.Reverse(&a)
        case commandMerge=="direct":
            generator.Direct(&a)
        default:
            log.Fatal()
        }


        if a.Length<=20{
            fmt.Printf("start array: %v\n",a.Array)
            algorithms.Testimony_merge(&a)
            fmt.Printf("finish array: %v\n",a.Array)
        }else{
            algorithms.Testimony_merge(&a)
        }

        fmt.Printf("Time: %v\nSwapped: cannot be determined\n",a.Time)
    },
}

func init() {
    mergeCmd.Flags().StringVarP(&commandMerge,"command","c","random","random/reverse/direct")
    rootCmd.AddCommand(mergeCmd)
}