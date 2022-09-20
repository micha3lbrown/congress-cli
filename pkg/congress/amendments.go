package congress

type AmendmentsAPIResponse struct {
	Amendments []Amendment
}

type Amendment struct {
	Congress     int
	LatestAction LatestAction
	Number       string
	Purpose      string
	Type         string
	Url          string
}

func GetAmendments() AmendmentsAPIResponse {
	var aar = AmendmentsAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "amendment", &aar, nil)
	sendRequest(rp)
	return aar
}
