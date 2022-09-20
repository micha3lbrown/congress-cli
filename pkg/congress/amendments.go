package congress

type AmendmentsAPIResponse struct {
	Amendments []Amendment
	Pagination
}

type Amendment struct {
	Congress int
	LatestAction
	Number  string
	Purpose string
	Type    string
	Url     string
}

func GetAmendments() AmendmentsAPIResponse {
	var aar = AmendmentsAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "amendment", &aar, nil)
	sendRequest(rp)
	return aar
}
