package congress

import (
	"fmt"
	"os"
)

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

func GetSummaries(start string, end string) SummariesAPIResponse {
	var sar = SummariesAPIResponse{}
	dateParams := make(map[string]string)
	dateParams["startDate"] = start
	dateParams["endDate"] = end
	fmt.Println(dateParams)
	c := newClient()
	// Time/Date Format: 2022-04-01T00:00:00Z
	// 									 YYYY-MM-DDThh:mm:ssTZD
	rp := newRequestParams(c, "summaries", &sar, dateParams)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return sar
}
