package sortingCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sorting",
	Version: "0.0.1",
	Short: "short",
	Long: "long",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute(){
	err := rootCmd.Execute()
	if err!=nil{
		fmt.Printf("!!! ERROR: '%s' !!!",err)
		os.Exit(1)
	}
}