package congress

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type requestParams struct {
	client   *http.Client
	Path     string
	Params   map[string]string
	respType interface{}
}

func newClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func newRequestParams(client *http.Client, path string, respType interface{}, params map[string]string) requestParams {
	return requestParams{
		client:   client,
		Path:     path,
		Params:   params,
		respType: respType,
	}
}

func sendRequest(rp requestParams) error {
	var CongressAPIVersion = "v3"
	var CongressGovBaseURL = fmt.Sprintf("https://api.congress.gov/%s/%s", CongressAPIVersion, rp.Path)
	var CongressGovAPIKey = os.Getenv("CONGRESS_GOV_API_KEY")

	req, err := http.NewRequest("GET", CongressGovBaseURL, nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("api_key", CongressGovAPIKey)
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()
	if rp.Params["startDate"] != "" {
		q.Add("fromDateTime", rp.Params["startDate"])
	}
	if rp.Params["startDate"] != "" {
		q.Add("toDateTime", rp.Params["endDate"])
	}

	resp, err := rp.client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("Requests Remaining: " + resp.Header.Get("x-ratelimit-remaining"))
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("[API Failure] Status Code: %d", resp.StatusCode)
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&rp.respType)
	if err != nil {
		return err
	}
	return nil
}
