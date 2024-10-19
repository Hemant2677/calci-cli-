/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// multiCmd represents the multi command
var multi = &cobra.Command{
	Use:   "multi",
	Short: "multiplication of two or more numbers",
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

		result := 1.0
		for _, numStr := range args {
			num, _ := strconv.ParseFloat(numStr, 64)
			result *= num
		}

		fmt.Printf("The product of %s is: %.2f\n", strings.Join(args, " * "), result)
	},
}

func init() {
	rootCmd.AddCommand(multi)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
