/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// cubeCmd represents the cube command
var cube = &cobra.Command{
	Use:   "cube",
	Short: "this command generates cube of any number",
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
		if err != nil {
			fmt.Printf("Error: %v is not a valid number.\n", args[0])
			return
		}

		result := math.Pow(num, 3)
		fmt.Printf("The cube of %f is: %f\n", num, result)
	},
}

func init() {
	rootCmd.AddCommand(cube)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cubeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
