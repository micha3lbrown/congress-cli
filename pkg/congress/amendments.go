package congress

type AmendmentsAPIResponse struct {
	Amendments []Amendment
}

type Amendment struct {
	Congress     string
	LatestAction LatestAction
	Number       string
	Purpose      string
	Type         string
	Url          string
}

func GetAmendments() AmendmentsAPIResponse {
	var aar = AmendmentsAPIResponse{}
	c := newClient()
	sendRequest(c, "amendment", &aar)
	return aar
}
