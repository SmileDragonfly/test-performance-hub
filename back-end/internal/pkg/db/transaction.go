package db

import (
	"code/internal/pkg/models"
	"context"
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)

func (q Queries) InsertTransaction(ctx context.Context, transaction *models.HubTransaction) error {
	query := `INSERT INTO dbo.tblHubTransaction ("Id", "TerminalHubID", "TerminalID", "SerialNumber", "TranCode", "ServerDate", "TsqHub", "CardMasked", "EncryptedPAN", "WithdrawalAmount", "OriginalAmount", "DebitAmount", "CashbackAmount", "TipAmount", "SurchargeAmount", "TranStatus", "ResponseCode", "Description", "SettlementDate", "AccountBalance", "AvailableBalance", "SeqNum", "TxnId", "SystemTraceNumber", "TransactionCounter", "ProcessorId", "ProcessorName", "RawRequest", "RawResponse") 
VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29)`
	_, err := q.db.ExecContext(ctx, query,
		transaction.ID,
		transaction.TerminalHubID,
		transaction.TerminalID,
		transaction.SerialNumber,
		transaction.TranCode,
		transaction.ServerDate,
		transaction.TsqHub,
		transaction.CardMasked,
		transaction.EncryptedPAN,
		transaction.WithdrawalAmount,
		transaction.OriginalAmount,
		transaction.DebitAmount,
		transaction.CashbackAmount,
		transaction.TipAmount,
		transaction.SurchargeAmount,
		transaction.TranStatus,
		transaction.ResponseCode,
		transaction.Description,
		transaction.SettlementDate,
		transaction.AccountBalance,
		transaction.AvailableBalance,
		transaction.SeqNum,
		transaction.TxnId,
		transaction.SystemTraceNumber,
		transaction.TransactionCounter,
		transaction.ProcessorId,
		transaction.ProcessorName,
		transaction.RawRequest,
		transaction.RawResponse)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) UpdateTransaction(ctx context.Context, transaction *models.HubTransaction) error {
	query := `UPDATE dbo.tblHubTransaction SET 
                                 ServerDate = @p1,
                                 TranStatus = @p2, 
                                 ResponseCode = @p3, 
                                 Description = @p4, 
                                 SettlementDate = @p5,
                                 AccountBalance = @p6,
                                 AvailableBalance = @p7,
                                 SeqNum = @p8,
                                 SystemTraceNumber = @p9 ,
                                 RawRequest = @p10,
                                 RawResponse = @p11,
                                 DebitAmount = @p12
                             WHERE Id = @p13`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(transaction.ID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		transaction.ServerDate,
		transaction.TranStatus,
		transaction.ResponseCode,
		transaction.Description,
		transaction.SettlementDate,
		transaction.AccountBalance,
		transaction.AvailableBalance,
		transaction.SeqNum,
		transaction.SystemTraceNumber,
		transaction.RawRequest,
		transaction.RawResponse,
		transaction.DebitAmount,
		uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) UpdateTransactionStatus(ctx context.Context, sID string, nStatus int64) error {
	query := `UPDATE dbo.tblHubTransaction SET 
                                 TranStatus = @p1 WHERE Id = @p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		nStatus,
		uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) GetTransactionForCheck(ctx context.Context, sSerialNumber string, sEncryptedPan string, fWithdrawAmount float64, sTransactionCounter string) (*models.HubTransaction, error) {
	query := `SELECT TOP 1 "Id", "TerminalHubID", "TerminalID", "SerialNumber", "TranCode", "ServerDate", "TsqHub", "CardMasked", "EncryptedPAN", "WithdrawalAmount", "OriginalAmount", "DebitAmount", "CashbackAmount", "TipAmount", "SurchargeAmount", "TranStatus", "ResponseCode", "Description", "SettlementDate", "AccountBalance", "AvailableBalance", "SeqNum", "TxnId", "SystemTraceNumber", "TransactionCounter", "ProcessorId", "ProcessorName", "RawRequest", "RawResponse" FROM dbo.tblHubTransaction 
								WHERE (SerialNumber = @p1) AND
                                    (EncryptedPAN = @p2) AND
                                    (WithdrawalAmount = @p3) AND
                                    (TransactionCounter = @p4) ORDER BY ServerDate DESC`
	var trxn models.HubTransaction
	var uuid mssql.UniqueIdentifier
	err := q.db.QueryRowContext(ctx, query,
		sSerialNumber,
		sEncryptedPan,
		fWithdrawAmount,
		sTransactionCounter).Scan(&uuid,
		&trxn.TerminalHubID,
		&trxn.TerminalID,
		&trxn.SerialNumber,
		&trxn.TranCode,
		&trxn.ServerDate,
		&trxn.TsqHub,
		&trxn.CardMasked,
		&trxn.EncryptedPAN,
		&trxn.WithdrawalAmount,
		&trxn.OriginalAmount,
		&trxn.DebitAmount,
		&trxn.CashbackAmount,
		&trxn.TipAmount,
		&trxn.SurchargeAmount,
		&trxn.TranStatus,
		&trxn.ResponseCode,
		&trxn.Description,
		&trxn.SettlementDate,
		&trxn.AccountBalance,
		&trxn.AvailableBalance,
		&trxn.SeqNum,
		&trxn.TxnId,
		&trxn.SystemTraceNumber,
		&trxn.TransactionCounter,
		&trxn.ProcessorId,
		&trxn.ProcessorName,
		&trxn.RawRequest,
		&trxn.RawResponse)
	if err != nil {
		return nil, err
	}
	trxn.ID = uuid.String()
	return &trxn, nil
}

func (q Queries) GetTransaction(ctx context.Context, sID string) (*models.HubTransaction, error) {
	query := `SELECT "Id", "TerminalHubID", "TerminalID", "SerialNumber", "TranCode", "ServerDate", "TsqHub", "CardMasked", "EncryptedPAN", "WithdrawalAmount", "OriginalAmount", "DebitAmount", "CashbackAmount", "TipAmount", "SurchargeAmount", "TranStatus", "ResponseCode", "Description", "SettlementDate", "AccountBalance", "AvailableBalance", "SeqNum", "TxnId", "SystemTraceNumber", "TransactionCounter", "ProcessorId", "ProcessorName", "RawRequest", "RawResponse" FROM dbo.tblHubTransaction WHERE Id=@p1`
	var transaction models.HubTransaction
	var uuid mssql.UniqueIdentifier
	err := q.db.QueryRowContext(ctx, query, sID).Scan(&uuid,
		&transaction.TerminalHubID,
		&transaction.TerminalID,
		&transaction.SerialNumber,
		&transaction.TranCode,
		&transaction.ServerDate,
		&transaction.TsqHub,
		&transaction.CardMasked,
		&transaction.EncryptedPAN,
		&transaction.WithdrawalAmount,
		&transaction.OriginalAmount,
		&transaction.DebitAmount,
		&transaction.CashbackAmount,
		&transaction.TipAmount,
		&transaction.SurchargeAmount,
		&transaction.TranStatus,
		&transaction.ResponseCode,
		&transaction.Description,
		&transaction.SettlementDate,
		&transaction.AccountBalance,
		&transaction.AvailableBalance,
		&transaction.SeqNum,
		&transaction.TxnId,
		&transaction.SystemTraceNumber,
		&transaction.TransactionCounter,
		&transaction.ProcessorId,
		&transaction.ProcessorName,
		&transaction.RawRequest,
		&transaction.RawResponse)
	if err != nil {
		return nil, err
	}
	transaction.ID = uuid.String()
	_, offset := time.Now().Zone()
	transaction.ServerDate = transaction.ServerDate.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	return &transaction, nil
}

func (q Queries) GetLastestTransaction(ctx context.Context, sHubID string, lastestTime time.Time) (*models.HubTransaction, error) {
	query := `SELECT TOP 1 "Id", "TerminalHubID", "TerminalID", "SerialNumber", "TranCode", "ServerDate", "TsqHub", "CardMasked", "EncryptedPAN", "WithdrawalAmount", "OriginalAmount", "DebitAmount", "CashbackAmount", "TipAmount", "SurchargeAmount", "TranStatus", "ResponseCode", "Description", "SettlementDate", "AccountBalance", "AvailableBalance", "SeqNum", "TxnId", "SystemTraceNumber", "TransactionCounter" , "ProcessorId", "ProcessorName", "RawRequest", "RawResponse"
FROM dbo.tblHubTransaction 
WHERE (TerminalHubID=@p1) AND (ServerDate > @p2) AND (TranStatus = 1) ORDER BY ServerDate desc`
	var transaction models.HubTransaction
	var uuid mssql.UniqueIdentifier
	sTime := lastestTime.Format("2006-01-02 15:04:05.000")
	err := q.db.QueryRowContext(ctx, query, sHubID, sTime).Scan(&uuid,
		&transaction.TerminalHubID,
		&transaction.TerminalID,
		&transaction.SerialNumber,
		&transaction.TranCode,
		&transaction.ServerDate,
		&transaction.TsqHub,
		&transaction.CardMasked,
		&transaction.EncryptedPAN,
		&transaction.WithdrawalAmount,
		&transaction.OriginalAmount,
		&transaction.DebitAmount,
		&transaction.CashbackAmount,
		&transaction.TipAmount,
		&transaction.SurchargeAmount,
		&transaction.TranStatus,
		&transaction.ResponseCode,
		&transaction.Description,
		&transaction.SettlementDate,
		&transaction.AccountBalance,
		&transaction.AvailableBalance,
		&transaction.SeqNum,
		&transaction.TxnId,
		&transaction.SystemTraceNumber,
		&transaction.TransactionCounter,
		&transaction.ProcessorId,
		&transaction.ProcessorName,
		&transaction.RawRequest,
		&transaction.RawResponse)
	if err != nil {
		return nil, err
	}
	transaction.ID = uuid.String()
	_, offset := time.Now().Zone()
	transaction.ServerDate = transaction.ServerDate.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	return &transaction, nil
}

func (q Queries) UpdateTransactionApiRawResponse(ctx context.Context, sID string, sRawResponse string) error {
	query := `UPDATE dbo.tblHubTransaction SET 
                                 ApiRawResponse = @p1 WHERE Id = @p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		sRawResponse,
		uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) UpdateTransactionRawRequest(ctx context.Context, sID string, sRawRequest string) error {
	query := `UPDATE dbo.tblHubTransaction SET 
                                 RawRequest = @p1 WHERE Id = @p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		sRawRequest,
		uuid)
	if err != nil {
		return err
	}
	return nil
}
