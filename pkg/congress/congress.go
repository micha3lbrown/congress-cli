package congress

type CongressAPIResponse struct {
	Congresses []Congresses
	Pagination CongressPagination `json:"pagination"`
	Request    CongressRequest    `json:"request"`
}

type Congresses struct {
	EndYear  string
	Name     string
	Sessions []Session
}

type CongressPagination struct {
	Count int
	Next  string
}

type CongressRequest struct {
	ContentType string
	Format      string
}

type Session struct {
	Chamber   string
	EndDate   string
	Number    int
	StartDate string
	Type      string
}

func GetCongress() CongressAPIResponse {
	var car = CongressAPIResponse{}
	c := newClient()
	sendRequest(c, "congress", &car)
	return car
}
