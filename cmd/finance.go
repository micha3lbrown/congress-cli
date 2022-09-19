/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// financeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// financeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getFinanceData() {
	PropublicaAPIKey := os.Getenv("PROPUBLICA_API_KEY")
	var PropublicaCampaignFinanceURL = "https://api.propublica.org/campaign-finance/v1/2022/candidates/leaders/contribution-total.json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", PropublicaCampaignFinanceURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("X-API-Key", PropublicaAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
