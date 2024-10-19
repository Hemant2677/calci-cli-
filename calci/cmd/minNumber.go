/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// minNumberCmd represents the minNumber command
var min = &cobra.Command{
	Use:   "min",
	Short: "this command provides minimum value from a list of",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
            fmt.Println("Error: At least two numbers must be provided.")
            return
        }

        min := args[0]
        for _, num := range args[1:] {
            if num < min {
                min = num
            }
        }

        fmt.Printf("The minimum number is: %s\n", min)
	},
}

func init() {
	rootCmd.AddCommand(min)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
