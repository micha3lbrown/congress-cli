/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// financeCmd represents the finance command
var financeCmd = &cobra.Command{
	Use:   "finance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finance called")
		getFinanceData()
	},
}

func init() {
	rootCmd.AddCommand(financeCmd)
}

func getFinanceData() {
	fmt.Println("Make some call using pkg/campaignFinance")
}
