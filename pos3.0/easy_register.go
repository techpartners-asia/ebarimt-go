package pos3_0

import "encoding/json"

func (p *pos3_0) ConsumerInfo(regNo string) (ConsumerInfoResponse, error) {
	response, err := p.httpRequest(nil, ConsumerInfoAPI, regNo, nil)
	if err != nil {
		return ConsumerInfoResponse{}, err
	}
	var resp ConsumerInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) GetProfile(body GetProfileRequest) (GetProfileResponse, error) {
	response, err := p.httpRequest(body, GetProfileAPI, "", nil)
	if err != nil {
		return GetProfileResponse{}, err
	}
	var resp GetProfileResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) ApproveQr(body ApproveQrRequest) (ApproveQrResponse, error) {
	response, err := p.httpRequest(body, ApproveQrAPI, "", nil)
	if err != nil {
		return ApproveQrResponse{}, err
	}
	var resp ApproveQrResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) ForiegnerPassportInfo(fNumber, passportNo string) (ForiegnerInfoResponse, error) {
	response, err := p.httpRequest(nil, ForiegnerPassportInfoAPI, passportNo+"/"+fNumber, nil)
	if err != nil {
		return ForiegnerInfoResponse{}, err
	}
	var resp ForiegnerInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) ForiegnerCustomerNoInfo(loginName string) (ForiegnerInfoResponse, error) {
	response, err := p.httpRequest(nil, ForiegnerCustomerNoInfoAPI, loginName, nil)
	if err != nil {
		return ForiegnerInfoResponse{}, err
	}
	var resp ForiegnerInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) ForiegnerInfoRegister(passportNo string, body ForiegnerInfoRequest) (ForiegnerInfoResponse, error) {
	response, err := p.httpRequest(body, ForiegnerCustomerNoInfoAPI, passportNo, nil)
	if err != nil {
		return ForiegnerInfoResponse{}, err
	}
	var resp ForiegnerInfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}
