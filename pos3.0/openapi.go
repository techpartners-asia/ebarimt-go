package pos3_0

import "encoding/json"

func (p *pos3_0) GetBranchInfo() (GetBranchInfoResponse, error) {
	response, err := p.httpRequest(nil, GetBranchInfoAPI, "", nil, false)
	if err != nil {
		return GetBranchInfoResponse{}, err
	}
	var resp GetBranchInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetTinInfo(regNo string) (GetTinInfoResponse, error) {
	response, err := p.httpRequest(nil, GetTinInfoAPI, regNo, nil, false)
	if err != nil {
		return GetTinInfoResponse{}, err
	}
	var resp GetTinInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetInfo(regNo string) (GetInfoResponse, error) {
	response, err := p.httpRequest(nil, GetInfoAPI, regNo, nil, false)
	if err != nil {
		return GetInfoResponse{}, err
	}
	var resp GetInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetSalesTotalData(body GetSalesTotalDataRequest) (GetSalesTotalDataResponse, error) {
	var headers []CustomHeader
	header := CustomHeader{
		Name:  "X-API-KEY",
		Value: p.apiKey,
	}
	headers = append(headers, header)
	response, err := p.httpRequest(body, GetSalesTotalAPI, "", headers, true)
	if err != nil {
		return GetSalesTotalDataResponse{}, err
	}
	var resp GetSalesTotalDataResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetSalesListERP(body GetSalesListERPRequest) (GetSalesTotalDataResponse, error) {
	var headers []CustomHeader
	header := CustomHeader{
		Name:  "X-API-KEY",
		Value: p.apiKey,
	}
	headers = append(headers, header)
	response, err := p.httpRequest(body, GetSalesListERPAPI, "", headers, true)
	if err != nil {
		return GetSalesTotalDataResponse{}, err
	}
	var resp GetSalesTotalDataResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) SaveOprMerchants(body SaveOprMerchantsRequest) (SaveOprMerchantsResponse, error) {
	response, err := p.httpRequest(body, GetSalesListERPAPI, "", nil, false)
	if err != nil {
		return SaveOprMerchantsResponse{}, err
	}
	var resp SaveOprMerchantsResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}
