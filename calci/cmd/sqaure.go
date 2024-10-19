/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// sqaureCmd represents the sqaure command
var sq = &cobra.Command{
	Use:   "sq",
	Short: "this command gives square of any value",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: One number must be provided.")
			return
		}
		num, err := strconv.ParseFloat(args[0], 64)
		if err!= nil {
            fmt.Printf("Error: Invalid number '%s'.\n", args[0])
            return
        }
		result := num * num
		fmt.Printf("The square of %f is: %f\n", num, result)
	},
}

func init() {
	rootCmd.AddCommand(sq)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqaureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqaureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
