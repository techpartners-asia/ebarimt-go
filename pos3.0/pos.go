package pos3_0

type pos3_0 struct {
	posEndpoint string
	apiKey      string
	token       *TokenResponse
}
type Pos3_0 interface {
	GetInfo(regNo string) (GetInfoResponse, error)
	GetTinInfo(regNo string) (GetTinInfoResponse, error)
	GetBranchInfo() (GetBranchInfoResponse, error)
	GetSalesTotalData(body GetSalesTotalDataRequest) (GetSalesTotalDataResponse, error)
	GetSalesListERP(body GetSalesListERPRequest) (GetSalesTotalDataResponse, error)
	SaveOprMerchants(body SaveOprMerchantsRequest) (SaveOprMerchantsResponse, error)
}

func New(endpoint, apiKey string) Pos3_0 {
	return &pos3_0{
		apiKey:      apiKey,
		posEndpoint: endpoint,
	}
}
