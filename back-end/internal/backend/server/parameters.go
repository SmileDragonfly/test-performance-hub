package server

type ResponseCommon struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

func HttpError(status int, sErr string) (int, ResponseCommon) {
	return status, ResponseCommon{
		ErrorCode: status,
		ErrorDesc: sErr,
	}
}

type RequestCreateTestData struct {
	HubPrefix      string `json:"hub_prefix" binding:"required,len=4"`
	NumberOfHub    int    `json:"number_of_hub" binding:"required,min=1,max=9999"`
	TerminalPerHub int    `json:"terminal_per_hub" binding:"required,min=1,max=99"`
	ShopId         string `json:"shop_id" binding:"required,omitempty,uuid_rfc4122"`
	ModeId         string `json:"mode_id" binding:"required,omitempty,uuid_rfc4122"`
	OsVersionId    string `json:"os_version_id" binding:"required,omitempty,uuid_rfc4122"`
	ProcessorId    int    `json:"processor_id" binding:"required"`
}

type RequestRunTests struct {
	HubPrefix             string `json:"hub_prefix" binding:"required,len=4"`
	NumberOfHub           int    `json:"number_of_hub" binding:"required,min=1,max=9999"`
	DelayBetweenHubs      int    `json:"delay_between_hubs" binding:"min=0"`
	DelayBetweenTerminals int    `json:"delay_between_terminals"  binding:"min=0"`
	DelayBetweenRequests  int    `json:"delay_between_requests"  binding:"min=0"`
	Duration              int    `json:"duration"`
	Iteration             int    `json:"iteration"`
	HubUrl                string `json:"hub_url"`
}

type UriDeleteTestData struct {
	HubPrefix string `uri:"hub_prefix" binding:"required"`
}

type UriGetWorkingKey struct {
	SerialNumber string `uri:"serial_number"`
}

type ResponseGetWorkingKey struct {
	Key1    string `json:"key1"`
	Key2    string `json:"key2"`
	Key3    string `json:"key3"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	RawByte string `json:"rawByte"`
}

type RequestPayment struct {
	Total          int64  `json:"total"`
	Track2         string `json:"track2"`
	SN             string `json:"SN"`
	Pan            string `json:"pan"`
	Track1         string `json:"track1"`
	Emv            string `json:"emv"`
	Pin            string `json:"pin"`
	Track3         string `json:"track3"`
	OriginalAmount string `json:"originalAmount"`
	CashBack       string `json:"cashBack"`
	TipAmount      string `json:"tipAmount"`
}

type ResponsePayment struct {
	Code          int    `json:"code"`
	CustomMessage string `json:"customMessage"`
	SeqNum        int64  `json:"seqNum"`
	Tid           string `json:"tid"`
	Txnid         string `json:"txnid"`
}

type RequestBalanceInquiry struct {
	Track2 string `json:"track2"`
	SN     string `json:"SN"`
	Pan    string `json:"pan"`
	Track1 string `json:"track1"`
	Emv    string `json:"emv"`
	Pin    string `json:"pin"`
	Track3 string `json:"track3"`
	SeqNum string `json:"seqNum"`
}

type ResponseBalanceInquiry struct {
	Amount string `json:"amount"`
}
