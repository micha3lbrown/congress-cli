package congress

type BillsAPIResponse struct {
	Bills []Bill
}

type Bill struct {
	Congress          int
	LatestAction      LatestAction
	Number            string
	OriginChamber     string
	OriginChamberCode string
	Title             string
	Type              string
	UpdateDate        string
	Url               string
}

type LatestAction struct {
	ActionDate string
	Text       string
}

func GetBills() BillsAPIResponse {
	var bar = BillsAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "bill", &bar, nil)
	sendRequest(rp)
	return bar
}
