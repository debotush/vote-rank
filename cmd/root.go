package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "vote-rank",
    Short: "A CLI tool for ranking candidates based on voting results",
    Long:  `vote-rank processes voting data and runs through a 4-phase ranking algorithm.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
