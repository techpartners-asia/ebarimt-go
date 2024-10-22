package pos3_0

import "encoding/json"

// Цахим төлбөрийн баримтын систем /PosApi/-ээс үүсгэж буй төлбөрийн баримтын үйл ажиллагаа явуулж буй байршлын мэдээллийг “districtCode” гэсэн баганад бөглөн илгээдэг ба дээрх талбарт бөглөн, илгээх байршлын татварын алба, дэд албаны кодын жагсаалтын мэдээллийг энэхүү сервисээс авах боломжтой.
//
// Жишээ нь: Номин холдинг ХХК-ийн Архангай аймаг дахь салбараас үүсгэсэн баримтын “districtCode”-г 01-гэж бөглөн илгээнэ.
//
// Ebarimt.mn API call: GET /api/info/check/getBranchInfo
//
// See https://developer.itc.gov.mn/docs/ebarimt-api/vxfs8o0ezfnsa-tatvaryn-alba-ded-albany-nerijn-kod-zhagsaaltyn-servis
func (p *pos3_0) GetBranchInfo() (GetBranchInfoResponse, error) {
	response, err := p.httpRequest(nil, GetBranchInfoAPI, "", nil)
	if err != nil {
		return GetBranchInfoResponse{}, err
	}
	var resp GetBranchInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetTinInfo(regNo string) (GetTinInfoResponse, error) {
	response, err := p.httpRequest(nil, GetTinInfoAPI, regNo, nil)
	if err != nil {
		return GetTinInfoResponse{}, err
	}
	var resp GetTinInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetInfo(regNo string) (GetInfoResponse, error) {
	response, err := p.httpRequest(nil, GetInfoAPI, regNo, nil)
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
	response, err := p.httpRequest(body, GetSalesTotalAPI, "", headers)
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
	response, err := p.httpRequest(body, GetSalesListERPAPI, "", headers)
	if err != nil {
		return GetSalesTotalDataResponse{}, err
	}
	var resp GetSalesTotalDataResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) SaveOprMerchants(body SaveOprMerchantsRequest) (SaveOprMerchantsResponse, error) {
	response, err := p.httpRequest(body, GetSalesListERPAPI, "", nil)
	if err != nil {
		return SaveOprMerchantsResponse{}, err
	}
	var resp SaveOprMerchantsResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}
