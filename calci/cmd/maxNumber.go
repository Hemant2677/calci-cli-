/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// maxNumberCmd represents the maxNumber command
var max = &cobra.Command{
	Use:   "max",
	Short: "this command provides the maximum number of given list",
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

        max := args[0]
        for _, num := range args[1:] {
            if num > max {
                max = num
            }
        }

        fmt.Printf("The maximum number among %s is: %s\n", args, max)
	},
}

func init() {
	rootCmd.AddCommand(max)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// maxNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// maxNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
