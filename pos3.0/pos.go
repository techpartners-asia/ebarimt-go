package pos3_0

import "encoding/json"

func (p *pos3_0) ReceiptSend(body ReceiptRequest) (ReceiptResponse, error) {
	response, err := p.httpPosRequest(body, PosReceiptSendAPI, "", nil)
	if err != nil {
		return ReceiptResponse{}, err
	}
	var resp ReceiptResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) ReceiptDelete(body ReceiptDeleteRequest) (SuccessResponse, error) {
	_, err := p.httpPosRequest(body, PosReceiptDeleteAPI, "", nil)
	if err != nil {
		return SuccessResponse{}, err
	}
	return SuccessResponse{Success: true}, nil
}

func (p *pos3_0) SendData() (SuccessResponse, error) {
	_, err := p.httpPosRequest(nil, PosSendAPI, "", nil)
	if err != nil {
		return SuccessResponse{}, err
	}
	return SuccessResponse{Success: true}, nil
}

func (p *pos3_0) Info() (InfoResponse, error) {
	response, err := p.httpPosRequest(nil, PosInfoAPI, "", nil)
	if err != nil {
		return InfoResponse{}, err
	}
	var resp InfoResponse
	json.Unmarshal(response, &resp)
	return resp, nil
}

func (p *pos3_0) BankAccounts(tin string) ([]BankAccountData, error) {
	response, err := p.httpPosRequest(nil, PosBankAccAPI, "tin="+tin, nil)
	if err != nil {
		return nil, err
	}
	var resp []BankAccountData
	json.Unmarshal(response, &resp)
	return resp, nil
}
