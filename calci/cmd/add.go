/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// add represents the add command
var add = &cobra.Command{
	Use:   "add",
	Short: "Adds two numbers",
	Long:  "This command adds two numbers provided as arguments.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Error: Two numbers must be provided.")
			return
		}

		num1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid number '%s' for the first argument.\n", args[0])
			return
		}

		num2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: Invalid number '%s' for the second argument.\n", args[1])
			return
		}

		result := num1 + num2
		fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, result)
	},
}

func init() {
	rootCmd.AddCommand(add)
	//creating flag for addition commands



	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
