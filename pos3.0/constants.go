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

type PosStatusEnum string

const (
	PosStatusSuccess PosStatusEnum = "SUCCESS" // Төлбөрийн баримтын мэдээлэл амжилттай үүссэн.
	PosStatusError   PosStatusEnum = "ERROR"   // Төлбөрийн баримтын мэдээлэл үүсгэхэд алдаа гарсан.
	PosStatusPayment PosStatusEnum = "PAYMENT" // Төлбөрийн баримтын мэдээлэл үүсгэхэд төлбөрийн мэдээлэл дутуу.
)

type TaxTypeEnum string

const (
	TaxTypeVatAble TaxTypeEnum = "VAT_ABLE" // НӨАТ тооцох бүтээгдэхүүн, үйлчилгээ
	TaxTypeVatFree TaxTypeEnum = "VAT_FREE" // НӨАТ-аас чөлөөлөгдөх бүтээгдэхүүн, үйлчилгээ
	TaxTypeVatZero TaxTypeEnum = "VAT_ZERO" // НӨАТ-н 0 хувь тооцох бүтээгдэхүүн, үйлчилгээ
	TaxTypeNoVat   TaxTypeEnum = "NO_VAT"   // Монгол улсын хилийн гадна борлуулсан бүтээгдэхүүн үйлчилгээ
)

type ReceiptTypeEnum string

const (
	TypeB2CReceipt ReceiptTypeEnum = "B2C_RECEIPT"
	TypeB2BReceipt ReceiptTypeEnum = "B2B_RECEIPT"
	TypeB2CInvoice ReceiptTypeEnum = "B2C_INVOICE"
	TypeB2BInvoice ReceiptTypeEnum = "B2B_INVOICE"
)

type PaymentCodeEnum string

const (
	PaymentCodeCash PaymentCodeEnum = "CASH"         // Бэлнээр
	PaymentCodeCard PaymentCodeEnum = "PAYMENT_CARD" // Төлбөрийн карт
)

type PaymentStatusEnum string

const (
	PaymentStatusPaid     PaymentStatusEnum = "PAID"     // Төлбөр амжилттай хийгдсэнийг тодорхойлоно
	PaymentStatusPay      PaymentStatusEnum = "PAY"      // Төлбөрийн мэдээллийг “Баримтын мэдээлэл солилцох сервис”-г ашиглан гүйцэтгэнэ.
	PaymentStatusReversed PaymentStatusEnum = "REVERSED" // Төлбөр буцаагдсан
	PaymentStatusError    PaymentStatusEnum = "ERROR"    // Төлөлт амжилтгүй болсон
)

type BarcodeTypeEnum string

const (
	BarcodeTypeUndefined BarcodeTypeEnum = "UNDEFINED"
	BarcodeTypeGS1       BarcodeTypeEnum = "GS1"
	BarcodeTypeISBN      BarcodeTypeEnum = "ISBN"
)
