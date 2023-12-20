package db

import (
	"code/internal/pkg/models"
	"context"
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)

func (q Queries) InsertTransactionLock(ctx context.Context, lock *models.HubTxnLock) error {
	query := `INSERT INTO dbo.tblHubTxnLock ("Id", "TerminalHubId", "HubTID", "TerminalID", "CreatedDate", "ExpiredTime") 
VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`
	_, err := q.db.ExecContext(ctx, query,
		lock.ID,
		lock.TerminalHubID,
		lock.HubTID,
		lock.TerminalID,
		lock.CreatedDate,
		lock.ExpiredTime)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) GetTransactionLockByHubID(ctx context.Context, sHubID string) (*models.HubTxnLock, error) {
	query := `SELECT "Id", "TerminalHubId", "HubTID", "TerminalID", "CreatedDate", "ExpiredTime" FROM dbo.tblHubTxnLock WHERE TerminalHubId = @p1`
	var lock models.HubTxnLock
	var lockUuid mssql.UniqueIdentifier
	var hubUuid mssql.UniqueIdentifier
	err := q.db.QueryRowContext(ctx, query, sHubID).Scan(&lockUuid,
		&hubUuid,
		&lock.HubTID,
		&lock.TerminalID,
		&lock.CreatedDate,
		&lock.ExpiredTime)
	if err != nil {
		return nil, err
	}
	lock.ID = lockUuid.String()
	lock.TerminalHubID = hubUuid.String()
	_, offset := time.Now().Zone()
	lock.CreatedDate = lock.CreatedDate.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	lock.ExpiredTime = lock.ExpiredTime.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	return &lock, nil
}

func (q Queries) DeleteTransactionLockByHubID(ctx context.Context, sHubID string) error {
	query := `DELETE FROM  dbo.tblHubTxnLock WHERE TerminalHubId = @p1`
	_, err := q.db.ExecContext(ctx, query, sHubID)
	if err != nil {
		return err
	}
	return nil
}
