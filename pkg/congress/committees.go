package congress

import "fmt"

type CommitteesAPIResponse struct {
	Pagination Pagination
	Committees []Committee
}

type CommitteeReportAPIResponse struct {
	Pagination Pagination
	Reports    []Report
}

type Report struct {
	Citation   string
	UpdateDate string
	URL        string
}

type Committee struct {
	Chamber           string
	CommitteeTypeCode string
	Name              string
	Parent            Parent
	Subcommittees     []Subcommittee
	SystemCode        string
	URL               string
}

type Subcommittee struct {
	Name       string
	SystemCode string
	URL        string
}

type Parent struct {
	Name       string
	SystemCode string
	URL        string
}

func GetCommittees() CommitteesAPIResponse {
	var car = CommitteesAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "committee", &car, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return car
}

func GetCommitteReports() CommitteeReportAPIResponse {
	var crar = CommitteeReportAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "committeeReport", &crar, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return crar
}
