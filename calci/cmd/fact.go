/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// factCmd represents the fact command
var fact = &cobra.Command{
	Use:   "fact",
	Short: "this command gives factorial of any number",
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

		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if num < 0 {
			fmt.Println("Error: Factorial is not defined for negative numbers.")
			return
		}

		result := 1
		for i := 1; i <= num; i++ {
			result *= i
		}

		fmt.Printf("The factorial of %d is: %d\n", num, result)
	},
}

func init() {
	rootCmd.AddCommand(fact)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
