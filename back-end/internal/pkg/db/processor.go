package db

import (
	"code/internal/pkg/models"
	"context"
	"database/sql"
)

func (q Queries) GetProcessorByID(ctx context.Context, nID int64) (*models.Processor, error) {
	query := `SELECT "Id", "ProcessorName", "BIN", "IpAddress", "Port", "ProtocolTypeId", "EnableTLS", "Description" FROM dbo.tblProcessor WHERE  Id = @p1`
	var processor Processor
	err := q.db.QueryRowContext(ctx, query, nID).Scan(&processor.ID,
		&processor.ProcessorName,
		&processor.BIN,
		&processor.IpAddress,
		&processor.Port,
		&processor.ProtocolTypeID,
		&processor.EnableTLS,
		&processor.Description)
	if err != nil {
		return nil, err
	}
	// Map data
	modelProcessor := models.Processor{
		ID:             processor.ID.Int64,
		ProcessorName:  processor.ProcessorName.String,
		BIN:            processor.BIN.String,
		IpAddress:      processor.IpAddress.String,
		Port:           processor.Port.Int64,
		ProtocolTypeID: processor.ProtocolTypeID.Int64,
		EnableTLS:      processor.EnableTLS.Bool,
		Description:    processor.Description.String,
	}
	return &modelProcessor, nil
}

func (q Queries) GetResponseCodeDesc(ctx context.Context, sCode string, nProID int64) (string, error) {
	query := `SELECT Name FROM dbo.tblResponseCode WHERE Code=@p1 AND ProcessorId=@p2`
	var sName sql.NullString
	err := q.db.QueryRowContext(ctx, query, sCode, nProID).Scan(&sName)
	if err != nil {
		return "", err
	}
	return sName.String, nil
}
