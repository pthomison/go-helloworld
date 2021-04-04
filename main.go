package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

func main() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

var rootCmd = &cobra.Command{
  Use:   "go-helloworld",
  Short: "go-helloworld",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("hello world")
  },
}