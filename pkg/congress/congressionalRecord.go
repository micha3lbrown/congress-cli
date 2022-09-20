package congress

import "fmt"

type CongressionalRecordAPIResponse struct {
	Results    Result
	SetSize    int
	TotalCount int
}

type Result struct {
	IndexStart int
	Issues     []Issue
}

type Issue struct {
	Congress    string
	Id          int
	Issue       string
	Links       map[string]Link
	PublishDate string
	Session     string
	Volume      string
}

type Link struct {
	Label    string
	Ordinal  int
	PDFLinks []PDF
}

type PDF struct {
	Part string
	URL  string
}

func GetCongressionalRecords() CongressionalRecordAPIResponse {
	var crar = CongressionalRecordAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "congressional-record", &crar, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return crar
}
