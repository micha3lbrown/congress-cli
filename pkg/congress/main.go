package congress

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func newClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func sendRequest(client *http.Client, path string, respType interface{}) {
	var CongressAPIVersion = "v3"
	var CongressGovBaseURL = fmt.Sprintf("https://api.congress.gov/%s/%s", CongressAPIVersion, path)
	var CongressGovAPIKey = os.Getenv("CONGRESS_GOV_API_KEY")

	req, err := http.NewRequest("GET", CongressGovBaseURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("api_key", CongressGovAPIKey)
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&respType)
	if err != nil {
		fmt.Println(err)
	}
}
