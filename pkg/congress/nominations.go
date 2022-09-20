package congress

import "fmt"

type NominationsAPIResponse struct {
	Nominations []Nomination
	Pagination  Pagination
}

type Nomination struct {
	Citation       string
	Congress       int
	Description    string
	LatestAction   []LatestAction
	NominationType NominationType
	Number         int
	PartNumber     string
	ReceivedDate   string
	UpdateDate     string
	URL            string
}

type NominationType struct {
	IsCivilian bool
	IsMilitary bool
	Name       string
}

func GetNominations() NominationsAPIResponse {
	var responseType = NominationsAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "nomination", &responseType, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return responseType
}
