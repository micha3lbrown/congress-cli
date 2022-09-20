package congress

import "fmt"

type HouseCommincationsAPIResponse struct {
	HouseCommunications []HouseCommunication
	Pagination          Pagination
}

type HouseCommunication struct {
	Chamber             string
	CommunicationNumber int
	CommunicationType   map[string]string
	CongressNumber      int
	URL                 string
}

func GetHouseCommunications() HouseCommincationsAPIResponse {
	var responseType = HouseCommincationsAPIResponse{}
	c := newClient()
	rp := newRequestParams(c, "house-communication", &responseType, nil)
	err := sendRequest(rp)
	if err != nil {
		fmt.Println(err)
	}
	return responseType
}
