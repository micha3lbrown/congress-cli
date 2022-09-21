package campaignFinance

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var PropublicaCampaignFinanceURL = "https://api.propublica.org/campaign-finance/v1/2022/candidates/leaders/contribution-total.json"

func init() {
	PropublicaAPIKey := os.Getenv("PROPUBLICA_API_KEY")
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
