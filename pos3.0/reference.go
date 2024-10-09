package pos3_0

import "encoding/json"

func (p *pos3_0) DistrictCode() (DistrictResponse, error) {
	response, err := p.httpReferncesRequest(DistrictCodeAPI)
	if err != nil {
		return DistrictResponse{}, err
	}
	var resp DistrictResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}
