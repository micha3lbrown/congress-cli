package congress

import "fmt"

type SummariesAPIResponse struct {
	Summaries []Summary
}

type Summary struct {
	ActionDate         string
	ActionDesc         string
	Bill               Bill
	CurrentChamber     string
	CurrentChamberCode string
	LastSummaryUpdate  string
	Text               string
}

func GetSummaries() SummariesAPIResponse {
	var sar = SummariesAPIResponse{}
	c := newClient()
	err := sendRequest(c, "summaries", &sar)
	if err != nil {
		fmt.Println(err)
	}
	return sar
}
