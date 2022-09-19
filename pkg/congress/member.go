package congress

type MembersAPIResponse struct {
	Members []Member
}

type Member struct {
	BioguideId string
	Depiction  Depiction
	District   int
	Name       string
	Party      string
	Served     Served
	State      string
	Url        string
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
	sendRequest(c, "member", &mar)
	return mar
}