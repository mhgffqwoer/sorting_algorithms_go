package sorting

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mhgffqwoer/sorting_algorithms_go/model"
    "github.com/mhgffqwoer/sorting_algorithms_go/pkg/generator"
	"github.com/spf13/cobra"
)

var command string
var bubbleCmd = &cobra.Command{
    Use:   "bubble",
    Aliases: []string{"bubble"},
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
            Time: 0.0,
            Swapped: 0,
        }

        switch {
        case command=="random":
            generator.Random(&a)
            fmt.Print(a.Array)
            break
        case command=="reverse":
            generator.Reverse(&a)
            fmt.Print(a.Array)
            break
        case command=="direct":
            generator.Direct(&a)
            fmt.Print(a.Array)
            break
        default:
            log.Fatal()
        }

    },
}

func init() {
    bubbleCmd.Flags().StringVarP(&command,"command","c","random","random/reverse/direct")
    
    rootCmd.AddCommand(bubbleCmd)
}