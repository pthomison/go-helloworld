package main

import (
	"fmt"
	"github.com/spf13/cobra"
	utils "github.com/pthomison/golang-utils"
)

var (
	message string
	name string

	rootCmd = &cobra.Command{
		Use:   "go-helloworld",
		Short: "go-helloworld",
		Run: run,
	}

)

func main() {

	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

	err := rootCmd.Execute()

	utils.Check(err)
}

func run(cmd *cobra.Command, args []string) {
	// nameSt, err := cmd.Flags().GetBool("float")

	fmt.Println(message + ", " + name)
}
