package db

import (
	"backend/internal/pkg/models"
	"context"
	mssql "github.com/microsoft/go-mssqldb"
)

// Get terminal hub id and terminal id by serie no
func (q Queries) GetTerminalFromSerieNo(ctx context.Context, sSerieNo string) (*models.Terminal, error) {
	query := `SELECT "Id", "ShopId", "TerminalID", "TerminalName", "SerieNo", "TerminalModelId", "OsVersionId", "Address1", "Address2", "City", "State", "Location", "NetworkType", "CreatedDate", "ActivedDate", "LastModifiedDate", "Deleted", "Status", "IsActive", "Zipcode", "ProcessorTID", "ProcessorId", "DeviceID", "TerminalHubId"
FROM dbo.tblTerminal WHERE SerieNo = @p1 AND Status = 1 AND IsActive = 1`
	var terminal Terminal
	err := q.db.QueryRowContext(ctx, query, sSerieNo).Scan(&terminal.Id,
		&terminal.ShopId,
		&terminal.TerminalID,
		&terminal.TerminalName,
		&terminal.SerieNo,
		&terminal.TerminalModelId,
		&terminal.OsVersionId,
		&terminal.Address1,
		&terminal.Address2,
		&terminal.City,
		&terminal.State,
		&terminal.Location,
		&terminal.NetworkType,
		&terminal.CreatedDate,
		&terminal.ActivedDate,
		&terminal.LastModifiedDate,
		&terminal.Deleted,
		&terminal.Status,
		&terminal.IsActive,
		&terminal.Zipcode,
		&terminal.ProcessorTID,
		&terminal.ProcessorId,
		&terminal.DeviceID,
		&terminal.TerminalHubId)
	if err != nil {
		return nil, err
	}
	termRet := models.Terminal{
		Id:               terminal.Id.String(),
		ShopId:           terminal.ShopId.String(),
		TerminalID:       terminal.TerminalID.String,
		TerminalName:     terminal.TerminalName.String,
		SerieNo:          terminal.SerieNo.String,
		TerminalModelId:  terminal.TerminalModelId.String(),
		OsVersionId:      terminal.OsVersionId.String(),
		Address1:         terminal.Address1.String,
		Address2:         terminal.Address2.String,
		City:             terminal.City.String,
		State:            terminal.State.String,
		Location:         terminal.Location.String,
		NetworkType:      terminal.NetworkType.Int64,
		CreatedDate:      terminal.CreatedDate.Time,
		ActivedDate:      terminal.ActivedDate.Time,
		LastModifiedDate: terminal.LastModifiedDate.Time,
		Deleted:          terminal.Deleted.Bool,
		Status:           terminal.Status.Int64,
		IsActive:         terminal.IsActive.Bool,
		Zipcode:          terminal.Zipcode.String,
		ProcessorTID:     terminal.ProcessorTID.String,
		ProcessorId:      terminal.ProcessorId.Int64,
		DeviceID:         terminal.DeviceID.String,
		TerminalHubId:    terminal.TerminalHubId.String(),
	}
	return &termRet, nil
}

func (s Queries) GetTokenByTerminalID(ctx context.Context, sID string) (*models.Token, error) {
	query := `SELECT "Id", "TerminalID", "UserID", "Token", "FirebaseToken", "IP", "CreateDate", "ExpiryDate" FROM dbo.tblAPIToken WHERE TerminalID= @p1 AND ExpiryDate > GETDATE()`
	var token Token
	var uuid mssql.UniqueIdentifier
	err := s.db.QueryRowContext(ctx, query, sID).Scan(&token.ID,
		&uuid,
		&token.UserID,
		&token.Token,
		&token.FirebaseToken,
		&token.IP,
		&token.CreateDate,
		&token.ExpiryDate,
	)
	if err != nil {
		return nil, err
	}
	model := models.Token{
		ID:            token.ID.Int64,
		TerminalID:    uuid.String(),
		UserID:        token.UserID.String,
		Token:         token.Token.String,
		FirebaseToken: token.FirebaseToken.String,
		IP:            token.IP.String,
		CreateDate:    token.CreateDate.Time,
		ExpiryDate:    token.ExpiryDate.Time,
	}
	return &model, nil
}
