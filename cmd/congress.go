/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/micha3lbrown/congress-cli/pkg/congress"
	"github.com/spf13/cobra"
)

var CongressGovAPIKey string

var congressCmd = &cobra.Command{
	Use:   "congress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetCongress()
		for _, v := range resp.Congresses {
			fmt.Println(v.Name, v.EndYear)
			fmt.Println("================================")
			for _, v := range v.Sessions {

				fmt.Println(v.Number)
				fmt.Println(v.Type)
				fmt.Println(v.Chamber)
				fmt.Println(v.StartDate, v.EndDate)

				fmt.Println()
			}
		}
	},
}

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetMembers()
		fmt.Println(resp)
	},
}

var billsCmd = &cobra.Command{
	Use:   "bills",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bills called")
	},
}

func init() {
	rootCmd.AddCommand(congressCmd)
	congressCmd.AddCommand(membersCmd)
	congressCmd.AddCommand(billsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// congressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// congressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
