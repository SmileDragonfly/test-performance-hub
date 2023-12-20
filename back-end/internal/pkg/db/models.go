package db

import (
	"database/sql"
	mssql "github.com/denisenkom/go-mssqldb"
)

// TerminalHub represents the structure of the "tblTerminalHub" table in SQL Server.
type TerminalHub struct {
	ID                      mssql.UniqueIdentifier `json:"Id"`                      // Unique Identifier
	TerminalHubID           sql.NullString         `json:"TerminalHubID"`           // Terminal Hub ID
	ProcessorId             sql.NullInt64          `json:"ProcessorId"`             // Processor ID
	Cassette1Denom          sql.NullInt64          `json:"Cassette1Denom"`          // Denomination for Cassette 1
	Cassette2Denom          sql.NullInt64          `json:"Cassette2Denom"`          // Denomination for Cassette 2
	Cassette1Remain         sql.NullInt64          `json:"Cassette1Remain"`         // Remaining Amount in Cassette 1
	Cassette2Remain         sql.NullInt64          `json:"Cassette2Remain"`         // Remaining Amount in Cassette 2
	TotalLowAmount          sql.NullInt64          `json:"TotalLowAmount"`          // Total Low Amount
	WorkingKey              sql.NullString         `json:"WorkingKey"`              // Working Key
	WorkingKeyInterval      sql.NullInt64          `json:"WorkingKeyInterval"`      // Working Key Interval
	HealthCheckInterval     sql.NullInt64          `json:"HealthCheckInterval"`     // Health Check Interval
	TxnDelayMin             sql.NullInt64          `json:"TxnDelayMin"`             // Minimum Transaction Delay
	TxnDelayMax             sql.NullInt64          `json:"TxnDelayMax"`             // Maximum Transaction Delay
	Surcharge               sql.NullFloat64        `json:"Surcharge"`               // Surcharge Amount
	Status                  sql.NullBool           `json:"Status"`                  // Status (true - Active, false - Inactive)
	Deleted                 sql.NullBool           `json:"Deleted"`                 // Deleted (true - Deleted, false - Not Deleted)
	CreatedDate             sql.NullTime           `json:"CreatedDate"`             // Creation Date and Time
	RequestTimeoutProcessor sql.NullInt64          `json:"RequestTimeoutProcessor"` // Request Timeout for Processor
	RequestTimeoutTerminal  sql.NullInt64          `json:"RequestTimeoutTerminal"`  // Request Timeout for Terminal
	CurrentTSQ              sql.NullInt64          `json:"CurrentTSQ"`              // Current Transaction Sequence Number
	LastRunWorkingKey       sql.NullTime           `json:"LastRunWorkingKey"`
	LastRunHealthCheck      sql.NullTime           `json:"LastRunHealthCheck"`
}

// Terminal represents the structure of the "tblTerminal" table in SQL Server.
type Terminal struct {
	ID               mssql.UniqueIdentifier `json:"Id"`               // Unique Identifier
	ShopID           mssql.UniqueIdentifier `json:"ShopId"`           // Shop ID (nullable)
	TerminalID       sql.NullString         `json:"TerminalID"`       // Terminal ID (Not Null)
	TerminalName     sql.NullString         `json:"TerminalName"`     // Terminal Name
	SerieNo          sql.NullString         `json:"SerieNo"`          // Serial Number
	TerminalModelID  mssql.UniqueIdentifier `json:"TerminalModelId"`  // Terminal Model ID (Not Null)
	OsVersionID      mssql.UniqueIdentifier `json:"OsVersionId"`      // OS Version ID (nullable)
	Address1         sql.NullString         `json:"Address1"`         // Address Line 1
	Address2         sql.NullString         `json:"Address2"`         // Address Line 2
	City             sql.NullString         `json:"City"`             // City
	State            sql.NullString         `json:"State"`            // State
	Location         sql.NullString         `json:"Location"`         // Location
	NetworkType      sql.NullInt64          `json:"NetworkType"`      // Network Type (Not Null)
	CreatedDate      sql.NullTime           `json:"CreatedDate"`      // Creation Date and Time (Not Null)
	ActivedDate      sql.NullTime           `json:"ActivedDate"`      // Activation Date and Time
	LastModifiedDate sql.NullTime           `json:"LastModifiedDate"` // Last Modified Date and Time
	Deleted          sql.NullBool           `json:"Deleted"`          // Deleted (Not Null)
	Status           sql.NullInt64          `json:"Status"`           // Status (Not Null)
	IsActive         sql.NullBool           `json:"IsActive"`         // Is Active (Not Null)
	Zipcode          sql.NullString         `json:"Zipcode"`          // Zipcode
	ProcessorTID     sql.NullString         `json:"ProcessorTID"`     // Processor Terminal ID
	ProcessorID      sql.NullInt64          `json:"ProcessorId"`      // Processor ID (nullable)
	DeviceID         sql.NullString         `json:"DeviceID"`         // Device ID
	TerminalHubID    mssql.UniqueIdentifier `json:"TerminalHubId"`    // Terminal Hub ID (nullable)
}

// HubTransaction represents the structure of the "tblHubTransaction" table in SQL Server.
type HubTransaction struct {
	ID                 sql.NullString `json:"Id"`                 // Unique Identifier
	TerminalHubID      sql.NullString `json:"TerminalHubID"`      // Terminal Hub ID
	TerminalID         sql.NullString `json:"TerminalID"`         // Terminal ID
	SerialNumber       sql.NullString `json:"SerialNumber"`       // Serial Number
	TranCode           sql.NullString `json:"TranCode"`           // Transaction Code
	ServerDate         sql.NullTime   `json:"ServerDate"`         // Server Date and Time
	TsqHub             sql.NullString `json:"TsqHub"`             // TSQ (Transaction Sequence Number) for the Hub
	CardMasked         sql.NullString `json:"CardMasked"`         // Masked Card Number
	EncryptedPAN       sql.NullString `json:"EncryptedPAN"`       // Encrypted PAN (Primary Account Number)
	WithdrawalAmount   sql.NullInt64  `json:"WithdrawalAmount"`   // Withdrawal Amount
	OriginalAmount     sql.NullInt64  `json:"OriginalAmount"`     // Original Transaction Amount
	CashbackAmount     sql.NullInt64  `json:"CashbackAmount"`     // Cashback Amount
	TipAmount          sql.NullInt64  `json:"TipAmount"`          // Tip Amount
	SurchargeAmount    sql.NullInt64  `json:"SurchargeAmount"`    // Surcharge Amount
	TranStatus         sql.NullInt64  `json:"TranStatus"`         // Transaction Status (1 - Success, 0 - Failed, 2 - Create)
	ResponseCode       sql.NullString `json:"ResponseCode"`       // Response Code
	Description        sql.NullString `json:"Description"`        // Description
	SettlementDate     sql.NullTime   `json:"SettlementDate"`     // Settlement Date and Time
	AccountBalance     sql.NullInt64  `json:"AccountBalance"`     // Account Balance
	AvailableBalance   sql.NullInt64  `json:"AvailableBalance"`   // Available Balance
	SeqNum             sql.NullString `json:"SeqNum"`             // Sequence Number
	TxnId              sql.NullString `json:"TxnId"`              // Transaction ID (Formatted as 'yyyyMMddHHmmss|TerminalHubID|TSQ')
	SystemTraceNumber  sql.NullString `json:"SystemTraceNumber"`  // System Trace Number
	TransactionCounter sql.NullString `json:"TransactionCounter"` // EMV: 9F36 Value
}

// HubTxnLock represents the structure of the "tblHubTxnLock" table in SQL Server.
type HubTxnLock struct {
	ID            sql.NullString `json:"Id"`            // Unique Identifier
	TerminalHubID sql.NullString `json:"TerminalHubId"` // Terminal Hub ID (nullable)
	HubTID        sql.NullString `json:"HubTID"`        // Hub Terminal ID
	TerminalID    sql.NullString `json:"TerminalID"`    // Terminal ID
	CreatedDate   sql.NullTime   `json:"CreatedDate"`   // Creation Date and Time
	ExpiredTime   sql.NullTime   `json:"ExpiredTime"`   // Expiration Time
}

// Processor represents the structure of the "tblProcessor" table in SQL Server.
type Processor struct {
	ID             sql.NullInt64  `json:"Id"`             // Unique Identifier
	ProcessorName  sql.NullString `json:"ProcessorName"`  // Processor Name
	BIN            sql.NullString `json:"BIN"`            // BIN (Bank Identification Number)
	IpAddress      sql.NullString `json:"IpAddress"`      // IP Address
	Port           sql.NullInt64  `json:"Port"`           // Port
	ProtocolTypeID sql.NullInt64  `json:"ProtocolTypeId"` // Protocol Type ID
	EnableTLS      sql.NullBool   `json:"EnableTLS"`      // Enable TLS (true - Enabled, false - Disabled)
	Description    sql.NullString `json:"Description"`    // Description
}

type Token struct {
	ID            sql.NullInt64
	TerminalID    sql.NullString
	UserID        sql.NullString
	Token         sql.NullString
	FirebaseToken sql.NullString
	IP            sql.NullString
	CreateDate    sql.NullTime
	ExpiryDate    sql.NullTime
}
