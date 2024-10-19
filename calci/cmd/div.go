/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var div = &cobra.Command{
	Use:   "div",
	Short: "this command is use for division of numbers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Error: Two numbers must be provided.")
			return
		}

		num1, _ := strconv.ParseFloat(args[0], 64)
		num2, _ := strconv.ParseFloat(args[1], 64)

		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}

		result := num1 / num2
		fmt.Printf("The division of %f and %f is: %f\n", num1, num2, result)
	},
}

func init() {
	rootCmd.AddCommand(div)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
