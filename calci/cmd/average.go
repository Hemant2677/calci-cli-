/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// averageCmd represents the average command
var average = &cobra.Command{
	Use:   "average",
	Short: "this command gives the average of given list of values",
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

		sum := 0.0
		for _, arg := range args {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Error: Invalid number '%s'\n", arg)
				return
			}
			sum += num
		}

		average := sum / float64(len(args))
		fmt.Printf("The average of the numbers is: %.2f\n", average)
	},
}

func init() {
	rootCmd.AddCommand(average)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// averageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// averageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
