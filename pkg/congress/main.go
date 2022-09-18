package congress

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetCongress() {
	var CongressAPIVersion = "v3"
	var CongressGovBaseURL = fmt.Sprintf("https://api.congress.gov/%s/congress", CongressAPIVersion)
	var CongressGovAPIKey = os.Getenv("CONGRESS_GOV_API_KEY")

	client := &http.Client{}
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

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
