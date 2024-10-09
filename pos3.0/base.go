package pos3_0

type pos3_0 struct {
	posEndpoint string
	apiKey      string
	token       *TokenResponse
}

func New(endpoint, apiKey string) Pos3_0 {
	return &pos3_0{
		apiKey:      apiKey,
		posEndpoint: endpoint,
	}
}

type Pos3_0 interface {
	// Цахим төлбөрийн баримт
	GetInfo(regNo string) (GetInfoResponse, error)
	GetTinInfo(regNo string) (GetTinInfoResponse, error)
	GetBranchInfo() (GetBranchInfoResponse, error)
	GetSalesTotalData(body GetSalesTotalDataRequest) (GetSalesTotalDataResponse, error)
	GetSalesListERP(body GetSalesListERPRequest) (GetSalesTotalDataResponse, error)
	SaveOprMerchants(body SaveOprMerchantsRequest) (SaveOprMerchantsResponse, error)
	// хялбар бүртгэл
	ConsumerInfo(regNo string) (ConsumerInfoResponse, error)
	GetProfile(body GetProfileRequest) (GetProfileResponse, error)
	ApproveQr(body ApproveQrRequest) (ApproveQrResponse, error)
	ForiegnerPassportInfo(fNumber, passportNo string) (ForiegnerInfoResponse, error)
	ForiegnerCustomerNoInfo(loginName string) (ForiegnerInfoResponse, error)
	ForiegnerInfoRegister(passportNo string, body ForiegnerInfoRequest) (ForiegnerInfoResponse, error)

	// POS
	ReceiptSend(body ReceiptRequest) (ReceiptResponse, error)
	ReceiptDelete(body ReceiptDeleteRequest) (SuccessResponse, error)
	SendData() (SuccessResponse, error)
	Info() (InfoResponse, error)
	BankAccounts(tin string) ([]BankAccountData, error)
}
