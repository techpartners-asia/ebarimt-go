package pos3_0

type (
	TokenRequest struct {
		GrantType string `json:"grant_type"` // Example: password
		Username  string `json:"username"`   // Example: easy-register-test
		Password  string `json:"password"`   // Example: Test@123
		ClientID  string `json:"client_id"`  // Example: vatps
	}
	TokenResponse struct {
		AccessToken      string `json:"access_token"` // Use this token
		ExpiresIn        string `json:"expires_in"`
		RefreshToken     string `json:"refresh_token"`
		RefreshExpiresIn string `json:"refresh_expires_in"`
		TokenType        string `json:"token_type"`
		NotBeforePolicy  string `json:"not-before-policy"`
		SessionState     string `json:"session_state"`
		Scope            string `json:"scope"`
	}
	ErrorResponse struct {
		Path             string `json:"path"`
		Status           int    `json:"status"`
		Message          string `json:"message"`
		Timestamp        string `json:"timestamp"`
		RequestID        string `json:"requestId"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
)

// public apis
type (
	GetBranchInfoResponse struct {
		Msg    string              `json:"msg"`    // Сервисийн хариу
		Status int                 `json:"status"` // Service ийн төлөв
		Code   string              `json:"code"`
		Data   []GetBranchInfoData `json:"data"`
	}
	GetBranchInfoData struct {
		BranchCode    string `json:"branchCode"`
		BranchName    string `json:"branchName"`
		SubBranchCode string `json:"subBranchCode"`
		SubBranchName string `json:"subBranchName"`
	}

	GetTinInfoResponse struct {
		Msg    string `json:"msg"`    // Сервисийн хариу
		Status int    `json:"status"` // Service ийн төлөв
		Code   string `json:"code"`
		Data   string `json:"data"` // ТИН дугаар
	}

	GetInfoResponse struct {
		Msg    string      `json:"msg"`    // Сервисийн хариу
		Status int         `json:"status"` // Service ийн төлөв
		Code   string      `json:"code"`
		Data   GetInfoData `json:"data"` // ТИН дугаар
	}
	GetInfoData struct {
		Name                   string `json:"name"`                   // Нэр
		FreeProject            bool   `json:"freeProject"`            // НӨАТ-аас чөлөөлөгдөх төсөл хөтөлбөр эсэх
		CityPayer              bool   `json:"cityPayer"`              // НХАТ суутгагч эсэх
		VatPayer               bool   `json:"vatPayer"`               // НӨАТ суутган төлөгч эсэх
		Found                  bool   `json:"found"`                  // Татвар төлөгчөөр бүртгэлтэй эсэх
		VatPayerRegisteredDate string `json:"vatpayerRegisteredDate"` // НӨАТ суутган төлөгчөөр бүртгүүлсэн огноо
	}
	GetSalesTotalDataRequest struct {
		Year       string           `json:"year"`       // Задаргаа авч буй баримтын он
		Month      string           `json:"month"`      // Задаргаа авч буй баримтын сар
		Day        string           `json:"day"`        // Задаргаа авч буй баримтын өдөр /Заавал биш/
		Status     GetSalesDataEnum `json:"status"`     // Задаргаа авч буй баримтын төрөл
		StartCount int              `json:"startCount"` // Тухайн баримтын эхлэх тоо
		EndCount   int              `json:"endCount"`   // Тухайн баримтын дуусах тоо
	}
	GetSalesTotalDataResponse struct {
		Msg    string      `json:"msg"`    // Сервисийн хариу
		Status int         `json:"status"` // Service ийн төлөв
		Code   string      `json:"code"`
		Data   GetInfoData `json:"data"`
	}
	GetSalesTotalData struct {
		List      []GetSalesData `json:"list"`
		PageModel PageModel      `json:"pageModel"`
	}
	GetSalesData struct {
		PosSid       string  `json:"posSid"`       // Посын системийн дугаар
		PosRno       string  `json:"posRno"`       // Төлбөрийн баримтын ДДТД
		PosRdate     string  `json:"posRdate"`     // Төлбөрийн баримтын огноо
		PosRamt      float64 `json:"posRamt"`      // Төлбөрийн баримтын нийт дүн
		CityTax      float64 `json:"cityTax"`      // Нийслэл хотын албан татвар
		PosVamt      float64 `json:"posVamt"`      // Нэмэгдсэн өртгийн албан татвар
		NetAmt       float64 `json:"netAmt"`       // Төлбөрийн баримтын цэвэр дүн
		FromType     string  `json:"fromType"`     // Хаанаас үүсгэсэн /POS, ebarimt/
		CsmrRegNo    string  `json:"csmrRegNo"`    // Худалдан авагчийн регистр
		CsmrName     string  `json:"csmrName"`     // Худалдан авагчийн нэр
		PosNo        string  `json:"posNo"`        // Посын дугаар
		OperatorName string  `json:"operatorName"` // Хэрэглэгчийн систем нийлүүлэгчийн нэр
		DistrictCode string  `json:"districtCode"` // Төлбөрийн баримтын байршлын нэр
	}
	PageModel struct {
		TotalElements int `json:"totalElements"` // Нийт баримтын тоо
	}

	GetSalesListERPRequest struct {
		Pin       string   `json:"Pin"`       // Толгой компанийн регистрийн дугаар
		StartDate string   `json:"StartDate"` // Баримт татах эхлэх огноо
		SubPin    []string `json:"subPin"`    // Охин компанийн регистрийн дугаар. Хэрэв регистр оруулахгүй бол толгой компанийн мэдээлэл татагдах
		EndDate   int      `json:"EndDate"`   // Баримт татах дуусах огноо
	}

	GEtSalesListERPResponse struct {
		Msg    string              `json:"msg"`    // Сервисийн хариу
		Status int                 `json:"status"` // Service ийн төлөв
		Code   string              `json:"code"`
		Data   GEtSalesListERPData `json:"data"`
	}
	GEtSalesListERPData struct {
		StartDate           string            `json:"startDate"`           // эхлэх огноо
		EndDate             int               `json:"endDate"`             // дуусах огноо
		RegNo               string            `json:"regNo"`               // Толгой компанийн регистр
		RecieptBuyModelList []RecieptBuyModel `json:"receiptBuyModelList"` // Худалдан авалтын баримтын жагсаалт
	}
	RecieptBuyModel struct {
		PrPosRno      string  `json:"prPosRno"`      // Төлбөрийн баримтын ДДТД
		Name          string  `json:"name"`          // Борлуулагчийн нэр
		RegNo         string  `json:"regNo"`         // Борлуулагчийн регистр
		BuyerRegNo    string  `json:"buyerRegNo"`    // Худалдан авагчийн регистр
		Date          string  `json:"date"`          // Баримт хэвлэсэн огноо
		AmountVAT     float64 `json:"amountVat"`     // НӨАТ
		AmountCityTax float64 `json:"amountCityTax"` // НХАТ
		AmountTotal   float64 `json:"amountTotal"`   // Төлбөрийн баримтын нийт дүн
		AmountNet     float64 `json:"amountNet"`     // Төлбөрийн баримтын цэвэр дүн
		FromType      string  `json:"fromType"`      // Хаанаас баримт хэвлэсэн /INVOICE, POS API/
	}

	SaveOprMerchantsRequest struct {
		OprRegNo string   `json:"oprRegNo"` // Оператор компани буюу Хэрэглэгчийн систем нийлүүлэгчийн регистр
		PosNo    string   `json:"posNo"`    // Пос-ын дугаар
		List     []string `json:"list"`     // Бүртгэх мерчантын жагсаалт
	}

	SaveOprMerchantsResponse struct {
		Msg    string        `json:"msg"`    // Сервисийн хариу
		Status int           `json:"status"` // Service ийн төлөв
		Code   string        `json:"code"`
		Data   []interface{} `json:"data"`
	}
)
