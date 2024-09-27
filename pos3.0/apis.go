package pos3_0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/techpartners-asia/ebarimt-go/utils"
)

var (
	TokenAPI = utils.API{
		Url:    "https://st.auth.itc.gov.mn/auth/realms/Staging/protocol/openid-connect/token",
		Method: http.MethodPost,
	}
	GetBranchInfoAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/info/check/getBranchInfo",
		Method: http.MethodGet,
	}
	GetTinInfoAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/info/check/getTinInfo?regNo=",
		Method: http.MethodGet,
	}

	GetInfoAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/info/check/getInfo?tin=",
		Method: http.MethodGet,
	}

	GetSalesTotalAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/tpi/receipt/getSalesTotalData",
		Method: http.MethodPost,
	}
	GetSalesListERPAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/tpi/receipt/getSaleListERP",
		Method: http.MethodPost,
	}
	SaveOprMerchantsAPI = utils.API{
		Url:    "https://api.ebarimt.mn/api/tpi/receipt/%20saveOprMerchants",
		Method: http.MethodPost,
	}
)

type CustomHeader struct {
	Name  string
	Value string
}

func (p *pos3_0) httpRequest(body interface{}, api utils.API, ext string, headers []CustomHeader, isAuth bool) ([]byte, error) {

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, api.Url+ext, requestBody)
	req.Header.Add("Accept", utils.HttpAcceptPublic)
	for _, header := range headers {
		req.Header.Add(header.Name, header.Value)
	}
	if isAuth {
		token, err := p.auth()
		if err != nil {
			return nil, err
		}
		p.token = &token
		req.Header.Add("Authorization", "Bearer "+p.token.AccessToken)
	}
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}
	return response, nil
}

func (q *pos3_0) auth() (authRes TokenResponse, err error) {
	if q.token != nil {
		expireInA, _ := time.Parse(time.RFC3339, q.token.ExpiresIn)
		expireInB := expireInA.Add(time.Duration(-12) * time.Hour)
		now := time.Now()
		if now.Before(expireInB) {
			authRes = *q.token
			err = nil
			return
		}
	}
	body := TokenRequest{
		GrantType: "",
		Username:  "",
		Password:  "",
		ClientID:  "",
	}

	requestByte, _ := json.Marshal(body)
	requestBody := bytes.NewReader(requestByte)

	req, err := http.NewRequest(TokenAPI.Method, TokenAPI.Url, requestBody)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", utils.HttpAcceptPrivate)
	req.Header.Add("Content-Type", utils.HttpContentType)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return authRes, fmt.Errorf("%s- Ebarimt POS 3.0 openid connect error response: %s", time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS), res.Status)
	}
	resp, _ := io.ReadAll(res.Body)
	json.Unmarshal(resp, &authRes)
	return authRes, nil
}
