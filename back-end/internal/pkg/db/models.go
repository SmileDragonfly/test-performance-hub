package db

import (
	"database/sql"
	mssql "github.com/microsoft/go-mssqldb"
)

// TerminalHub represents the structure of the "tblTerminalHub" table in SQL Server.
type TerminalHub struct {
	Id                      mssql.UniqueIdentifier `json:"Id" gorm:"column:Id"`
	TerminalHubID           sql.NullString         `json:"TerminalHubID" gorm:"column:TerminalHubID"`
	ProcessorId             sql.NullInt64          `json:"ProcessorId" gorm:"column:ProcessorId"`
	Cassette1Denom          sql.NullInt64          `json:"Cassette1Denom" gorm:"column:Cassette1Denom"`
	Cassette2Denom          sql.NullInt64          `json:"Cassette2Denom" gorm:"column:Cassette2Denom"`
	Cassette1Remain         sql.NullInt64          `json:"Cassette1Remain" gorm:"column:Cassette1Remain"`
	Cassette2Remain         sql.NullInt64          `json:"Cassette2Remain" gorm:"column:Cassette2Remain"`
	TotalLowAmount          sql.NullInt64          `json:"TotalLowAmount" gorm:"column:TotalLowAmount"`
	WorkingKey              sql.NullString         `json:"WorkingKey" gorm:"column:WorkingKey"`
	WorkingKeyInterval      sql.NullInt64          `json:"WorkingKeyInterval" gorm:"column:WorkingKeyInterval"`
	HealthCheckInterval     sql.NullInt64          `json:"HealthCheckInterval" gorm:"column:HealthCheckInterval"`
	TxnDelayMin             sql.NullInt64          `json:"TxnDelayMin" gorm:"column:TxnDelayMin"`
	TxnDelayMax             sql.NullInt64          `json:"TxnDelayMax" gorm:"column:TxnDelayMax"`
	Surcharge               sql.NullFloat64        `json:"Surcharge" gorm:"column:Surcharge"`
	Status                  sql.NullBool           `json:"Status" gorm:"column:Status"`
	Deleted                 sql.NullBool           `json:"Deleted" gorm:"column:Deleted"`
	CreatedDate             sql.NullTime           `json:"CreatedDate" gorm:"column:CreatedDate"`
	RequestTimeoutProcessor sql.NullInt64          `json:"RequestTimeoutProcessor" gorm:"column:RequestTimeoutProcessor"`
	RequestTimeoutTerminal  sql.NullInt64          `json:"RequestTimeoutTerminal" gorm:"column:RequestTimeoutTerminal"`
	CurrentTSQ              sql.NullInt64          `json:"CurrentTSQ" gorm:"column:CurrentTSQ"`
	LastRunWorkingKey       sql.NullTime           `json:"LastRunWorkingKey" gorm:"column:LastRunWorkingKey"`
	LastRunHealthCheck      sql.NullTime           `json:"LastRunHealthCheck" gorm:"column:LastRunHealthCheck"`
}

func (TerminalHub) TableName() string {
	return "tblTerminalHub"
}

// Terminal represents the structure of the "tblTerminal" table in SQL Server.
type Terminal struct {
	Id               mssql.UniqueIdentifier `json:"Id" gorm:"column:Id"`
	ShopId           mssql.UniqueIdentifier `json:"ShopId" gorm:"column:ShopId"`
	TerminalID       sql.NullString         `json:"TerminalID" gorm:"column:TerminalID;not null"`
	TerminalName     sql.NullString         `json:"TerminalName" gorm:"column:TerminalName"`
	SerieNo          sql.NullString         `json:"SerieNo" gorm:"column:SerieNo"`
	TerminalModelId  mssql.UniqueIdentifier `json:"TerminalModelId" gorm:"column:TerminalModelId;not null"`
	OsVersionId      mssql.UniqueIdentifier `json:"OsVersionId" gorm:"column:OsVersionId"`
	Address1         sql.NullString         `json:"Address1" gorm:"column:Address1"`
	Address2         sql.NullString         `json:"Address2" gorm:"column:Address2"`
	City             sql.NullString         `json:"City" gorm:"column:City"`
	State            sql.NullString         `json:"State" gorm:"column:State"`
	Location         sql.NullString         `json:"Location" gorm:"column:Location"`
	NetworkType      sql.NullInt64          `json:"NetworkType" gorm:"column:NetworkType;not null"`
	CreatedDate      sql.NullTime           `json:"CreatedDate" gorm:"column:CreatedDate;not null"`
	ActivedDate      sql.NullTime           `json:"ActivedDate" gorm:"column:ActivedDate"`
	LastModifiedDate sql.NullTime           `json:"LastModifiedDate" gorm:"column:LastModifiedDate"`
	Deleted          sql.NullBool           `json:"Deleted" gorm:"column:Deleted;not null"`
	Status           sql.NullInt64          `json:"Status" gorm:"column:Status;not null"`
	IsActive         sql.NullBool           `json:"IsActive" gorm:"column:IsActive;not null"`
	Zipcode          sql.NullString         `json:"Zipcode" gorm:"column:Zipcode"`
	ProcessorTID     sql.NullString         `json:"ProcessorTID" gorm:"column:ProcessorTID"`
	ProcessorId      sql.NullInt64          `json:"ProcessorId" gorm:"column:ProcessorId"`
	DeviceID         sql.NullString         `json:"DeviceID" gorm:"column:DeviceID"`
	TerminalHubId    mssql.UniqueIdentifier `json:"TerminalHubId" gorm:"column:TerminalHubId"`
}

func (Terminal) TableName() string {
	return "tblTerminal"
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
	ID             sql.NullInt64  `json:"Id" gorm:"column:Id"`                         // Unique Identifier
	ProcessorName  sql.NullString `json:"ProcessorName" gorm:"column:ProcessorName"`   // Processor Name
	BIN            sql.NullString `json:"BIN" gorm:"column:BIN"`                       // BIN (Bank Identification Number)
	IpAddress      sql.NullString `json:"IpAddress" gorm:"column:IpAddress"`           // IP Address
	Port           sql.NullInt64  `json:"Port" gorm:"column:Port"`                     // Port
	ProtocolTypeID sql.NullInt64  `json:"ProtocolTypeId" gorm:"column:ProtocolTypeId"` // Protocol Type ID
	EnableTLS      sql.NullBool   `json:"EnableTLS" gorm:"column:EnableTLS"`           // Enable TLS (true - Enabled, false - Disabled)
	Description    sql.NullString `json:"Description" gorm:"column:Description"`       // Description
}

func (Processor) TableName() string {
	return "tblProcessor"
}

type Token struct {
	ID            sql.NullInt64          `json:"ID" gorm:"column:ID"`
	TerminalID    mssql.UniqueIdentifier `json:"TerminalID" gorm:"column:TerminalID"`
	UserID        sql.NullString         `json:"UserID" gorm:"column:UserID"`
	Token         sql.NullString         `json:"Token" gorm:"column:Token"`
	FirebaseToken sql.NullString         `json:"FirebaseToken" gorm:"column:FirebaseToken"`
	IP            sql.NullString         `json:"IP" gorm:"column:IP"`
	CreateDate    sql.NullTime           `json:"CreateDate" gorm:"column:CreateDate"`
	ExpiryDate    sql.NullTime           `json:"ExpiryDate" gorm:"column:ExpiryDate"`
}

func (Token) TableName() string {
	return "tblAPIToken"
}

type Shop struct {
	ID           mssql.UniqueIdentifier `json:"Id" gorm:"column:Id;type:uniqueidentifier"`
	MerchantID   mssql.UniqueIdentifier `json:"MerchantId" gorm:"column:MerchantId;type:uniqueidentifier"`
	Name         sql.NullString         `json:"Name" gorm:"column:Name;type:nvarchar(100)"`
	Address1     sql.NullString         `json:"Address1" gorm:"column:Address1;type:nvarchar(500)"`
	Address2     sql.NullString         `json:"Address2" gorm:"column:Address2;type:nvarchar(200)"`
	City         sql.NullString         `json:"City" gorm:"column:City;type:nvarchar(100)"`
	State        sql.NullString         `json:"State" gorm:"column:State;type:nvarchar(100)"`
	Phone        sql.NullString         `json:"Phone" gorm:"column:Phone;type:varchar(50)"`
	Zipcode      sql.NullString         `json:"Zipcode" gorm:"column:Zipcode;type:varchar(20)"`
	Email        sql.NullString         `json:"Email" gorm:"column:Email;type:nvarchar(500)"`
	TimeZone     sql.NullString         `json:"TimeZone" gorm:"column:TimeZone;type:nvarchar(100)"`
	TimezoneName sql.NullString         `json:"TimezoneName" gorm:"column:TimezoneName;type:nvarchar(100)"`
	EmailTime    sql.NullString         `json:"EmailTime" gorm:"column:EmailTime;type:varchar(8)"`
	CreatedAt    sql.NullTime           `json:"CreatedAt" gorm:"column:CreatedAt"`
	UpdatedAt    sql.NullTime           `json:"UpdatedAt" gorm:"column:UpdatedAt"`
	DeletedAt    sql.NullTime           `json:"DeletedAt" gorm:"column:DeletedAt"`
}

// TableName specifies the table name for the YourStructName model.
func (Shop) TableName() string {
	return "tblShop"
}

type TerminalModel struct {
	ID       mssql.UniqueIdentifier `json:"Id" gorm:"column:Id;type:uniqueidentifier"`
	VendorID mssql.UniqueIdentifier `json:"VendorId" gorm:"column:VendorId;type:uniqueidentifier"`
	Code     sql.NullString         `json:"Code" gorm:"column:Code;type:varchar(50)"`
	Name     sql.NullString         `json:"Name" gorm:"column:Name;type:nvarchar(100)"`
}

// TableName specifies the table name for the YourStructName model.
func (TerminalModel) TableName() string {
	return "tblTerminalModel"
}

type OsVersion struct {
	ID        mssql.UniqueIdentifier `json:"Id" gorm:"column:Id;type:uniqueidentifier"`
	Name      sql.NullString         `json:"Name" gorm:"column:Name;type:varchar(100)"`
	VersionNo sql.NullString         `json:"VersionNo" gorm:"column:VersionNo;type:varchar(50)"`
}

// TableName specifies the table name for the YourStructName model.
func (OsVersion) TableName() string {
	return "tblOsVersion"
}
