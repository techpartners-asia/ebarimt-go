package pos3_0

type ebarimtbilltype string

const (
	EBarimtPersonType       = ebarimtbilltype("1")
	EBarimtOrganizationType = ebarimtbilltype("3")
)

type GetSalesDataEnum int

const (
	SalesDataAll GetSalesDataEnum = iota
	SalesDataB2B
	SalesDataLottery
	SalesDataInvoice
	SalesDataBatch
)
