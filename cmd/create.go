/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create prometheus-operator resources",
	Long: `This command is used to create prometheus-operator resources, such as:
    - prometheus`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().Bool("dry-run", false, "If true, only try to create the Prometheus instance without actually creating it.")
	createCmd.PersistentFlags().String("output", "", "The output format to use, can be 'yaml'.")
}
