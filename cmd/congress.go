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
		resp := congress.GetBills()
		fmt.Println(resp)
	},
}

var amendmentsCmd = &cobra.Command{
	Use:   "amendments",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetAmendments()
		fmt.Println(resp)
	},
}

var summariesCmd = &cobra.Command{
	Use:   "summaries",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetSummaries()
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(congressCmd)
	congressCmd.AddCommand(membersCmd)
	congressCmd.AddCommand(billsCmd)
	congressCmd.AddCommand(amendmentsCmd)
	congressCmd.AddCommand(summariesCmd)
}
