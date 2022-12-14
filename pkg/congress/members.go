package congress

import "fmt"

type MembersAPIResponse struct {
	Members []Member
	Pagination
}

type Member struct {
	BioguideId string
	Depiction
	District int
	Name     string
	Party    string
	Served   Served
	State    string
	Url      string
}

type Depiction struct {
	Attribution string
	ImageUrl    string
}

type Served struct {
	Senate []Term
	House  []Term
}

type Term struct {
	End   int
	Start int
}

func GetMembers() MembersAPIResponse {
	var mar = MembersAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "member", &mar, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return mar
}
