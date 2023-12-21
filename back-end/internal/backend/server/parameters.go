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
	DelayBetweenHubs      int    `json:"delay_between_hubs" binding:"required,min=0"`
	DelayBetweenTerminals int    `json:"delay_between_terminals" binding:"required,min=0"`
	Duration              int    `json:"duration" binding:"required,min=1"`
}

type UriDeleteTestData struct {
	HubPrefix string `uri:"hub_prefix" binding:"required"`
}
