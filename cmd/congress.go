/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/micha3lbrown/congress-cli/pkg/congress"
	"github.com/spf13/cobra"
)

var CongressGovAPIKey string

var congressCmd = &cobra.Command{
	Use:   "congress",
	Short: "A list of congressional sessions including name, year and sessions",
	Long:  `A list of congressional sessions including name, year and sessions`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetCongress()
		for _, v := range resp.Congresses {
			fmt.Println(v.Name, v.EndYear)
		}
	},
}

var congressionalRecordsCmd = &cobra.Command{
	Use:   "congressional-records",
	Short: "A list of congressional records and associated issues and sessions",
	Long:  `A list of congressional records and associated issues and sessions`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetCongressionalRecords()
		fmt.Println(resp)
	},
}

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "A list of the members of congress",
	Long:  `A list of the members of congress`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetMembers()
		fmt.Println(resp)
	},
}

var billsCmd = &cobra.Command{
	Use:   "bills",
	Short: "A list of bills and the associated congress/session",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetBills()
		fmt.Println(resp)
	},
}

var amendmentsCmd = &cobra.Command{
	Use:   "amendments",
	Short: "A list of amendments and associated details",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetAmendments()
		fmt.Println(resp)
	},
}

var committeesCmd = &cobra.Command{
	Use:   "committees",
	Short: "A list of congressional committees",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetCommittees()
		fmt.Println(resp)
	},
}

var committeeReportsCmd = &cobra.Command{
	Use:   "committeeReports",
	Short: "A list of congressional committee reports",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetCommitteReports()
		fmt.Println(resp)
	},
}

var houseCommunicationsCmd = &cobra.Command{
	Use:   "house-communications",
	Short: "A list of congressional house communications",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetHouseCommunications()
		fmt.Println(resp)
	},
}

var nominationsCmd = &cobra.Command{
	Use:   "nominations",
	Short: "A list of congressional nominations",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetNominations()
		fmt.Println(resp)
	},
}

var treatiesCmd = &cobra.Command{
	Use:   "treaties",
	Short: "A list of congressional treaties",
	Run: func(cmd *cobra.Command, args []string) {
		resp := congress.GetTreaties()
		fmt.Println(resp)
	},
}

var summariesCmd = &cobra.Command{
	Use:   "summaries",
	Short: "A list of congressional summaries by date range (default: past 6 months)",
	Run: func(cmd *cobra.Command, args []string) {
		dateFormat := "2006-04-01T00:00:00Z"
		// dateRange, err := cmd.Flags().GetString("since")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		startDate := time.Now().UTC().Format(dateFormat)
		endDate := time.Now().UTC().AddDate(0, -6, 0).Format(dateFormat)

		resp := congress.GetSummaries(string(startDate), string(endDate))

		for _, v := range resp.Summaries {
			fmt.Printf("Bill #%s\n%s\n%s\n\n\n", v.Bill.Number, v.Bill.Title, v.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(congressCmd)
	congressCmd.AddCommand(congressionalRecordsCmd, membersCmd, billsCmd, amendmentsCmd, committeesCmd, committeeReportsCmd, summariesCmd, houseCommunicationsCmd, nominationsCmd, treatiesCmd)
	congressCmd.PersistentFlags().StringP("since", "s", "6M", "Expects int followed by date range Y M i.e. 6M")
}
