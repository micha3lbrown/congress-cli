package congress

import "fmt"

type CongressAPIResponse struct {
	Congresses []Congresses
	Pagination Pagination
	Request    CongressRequest
}

type Congresses struct {
	EndYear  string
	Name     string
	Sessions []Session
}

type Pagination struct {
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
	rp := newRequestParams(c, "congress", &car, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return car
}
