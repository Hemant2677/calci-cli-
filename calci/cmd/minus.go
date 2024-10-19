/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// minusCmd represents the minus command
var sub = &cobra.Command{
	Use:   "sub [num1] [num2]",
	Short: "Subtracts second number from the first",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)!= 2 {
            fmt.Println("Error: Two numbers must be provided.")
            return
        }
		num1, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("Error: Invalid number '%s' for the first argument.\n", args[0])
			return
		}
		num2, err := strconv.ParseFloat(args[1], 64)
		if err!= nil {
            fmt.Printf("Error: Invalid number '%s' for the second argument.\n", args[1])
            return
        }
		result := num1 - num2
		fmt.Printf("Result: %f\n", result)
	},
}

func init() {
	rootCmd.AddCommand(sub)
	//creting flag for substraction functionality
	sub.Flags().BoolP("subtract", "s", false, "Subtract the second number from the first")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
