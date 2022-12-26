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

var commandQuick string
var quickCmd = &cobra.Command{
    Use:   "quick",
    Short:  "Quick sort",
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
        case commandQuick=="random":
            generator.Random(&a)
        case commandQuick=="reverse":
            generator.Reverse(&a)
        case commandQuick=="direct":
            generator.Direct(&a)
        default:
            log.Fatal()
        }


        if a.Length<=20{
            fmt.Printf("start array: %v\n",a.Array)
            algorithms.Testimony_quick(&a)
            fmt.Printf("finish array: %v\n",a.Array)
        }else{
            algorithms.Testimony_quick(&a)
        }

        fmt.Printf("Time: %v\nSwapped: %v\n",a.Time,a.Swapped)
    },
}

func init() {
    quickCmd.Flags().StringVarP(&commandQuick,"command","c","random","random/reverse/direct")
    rootCmd.AddCommand(quickCmd)
}