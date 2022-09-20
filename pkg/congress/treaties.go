package congress

import "fmt"

type TreatiesAPIResponse struct {
	Treaties   []Treaty
	Pagination Pagination
}

type Treaty struct {
	Congress        int
	EndCongressId   int
	Parts           map[string]string
	ResolutionText  string
	TransmittedDate string
	TreatyNumber    int
	TreatySubject   string
	TreatySuffix    string
	UpdateDate      string
	URL             string
}

func GetTreaties() TreatiesAPIResponse {
	var responseType = TreatiesAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "treaty", &responseType, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return responseType
}
