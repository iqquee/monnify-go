package transaction

type AcceptPaymentModel struct {
	PaymentReference    string `json:"paymentReference"`
	Amount              int    `json:"amount"`
	CurrencyCode        string `json:"currencyCode"`
	ContractCode        string `json:"contractCode"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerName        string `json:"customerName"`
	CustomerPhoneNumber string `json:"customerPhoneNumber"`
	RedirectUrl         string `json:"redirectUrl"`
	PaymentDescription  string `json:"paymentDescription"`
}

type InitTransacStatus struct {
	RequestSuccessful bool
	ResponseMessage   string
	ResponseCode      string
	ResponseBody      InitTransacStatusBody
}
type InitTransacStatusBody struct {
	TransactionReference string
	PaymentReference     string
	AmountPaid           string
	TotalPayable         string
	SettlementAmount     string
	PaidOn               string
	PaymentStatus        string
	PaymentDescription   string
	Currency             string
	PaymentMethod        string
}

func AcceptPayment(amount int, paymentReference, paymentDesc, currencyCode, contractCode, customerName, customerEmail, customerPhoneNumber, redirectUrl string) {
	// var
}
