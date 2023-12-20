package db

import (
	"code/internal/pkg/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

type Querier interface {
	// Terminal
	GetTerminalFromSerieNo(ctx context.Context, sSerieNo string) (*models.Terminal, error)
	// Terminal Hub
	// Return: workingkey (format: key1;key2;key3)
	GetWorkingKeyByTerminalHubID(ctx context.Context, sID string) (string, error)
	UpdateWorkingKeyForTerminalHub(ctx context.Context, sID string, sWorkingKey string) error
	GetTerminalHub(ctx context.Context, sID string) (*models.TerminalHub, error)
	UpdateTerminalHubCassette(ctx context.Context, sID string, nCassette1Remain int64, nCassette2Remain int64) error
	UpdateTerminalHubTSQ(ctx context.Context, sID string, nTSQ int64) error
	GetAllTerminalHub(ctx context.Context) ([]models.TerminalHub, error)
	UpdateLastRunWorkingKey(ctx context.Context, sID string, sLastRun time.Time) error
	UpdateLastRunHealthCheck(ctx context.Context, sID string, sLastRun time.Time) error
	// Processor
	GetProcessorByID(ctx context.Context, nID int64) (*models.Processor, error)
	GetResponseCodeDesc(ctx context.Context, sCode string, nProID int64) (string, error)
	// Transaction
	InsertTransaction(ctx context.Context, transaction *models.HubTransaction) error
	UpdateTransaction(ctx context.Context, transaction *models.HubTransaction) error
	UpdateTransactionStatus(ctx context.Context, sID string, nStatus int64) error
	GetTransaction(ctx context.Context, sID string) (*models.HubTransaction, error)
	GetTransactionForCheck(ctx context.Context, sSerialNumber string, sEncryptedPan string, fWithdrawAmount float64, sTransactionCounter string) (*models.HubTransaction, error)
	GetLastestTransaction(ctx context.Context, sHubID string, lastestTime time.Time) (*models.HubTransaction, error)
	UpdateTransactionApiRawResponse(ctx context.Context, sID string, sRawResponse string) error
	UpdateTransactionRawRequest(ctx context.Context, sID string, sRawRequest string) error
	// Transaction Lock
	InsertTransactionLock(ctx context.Context, lock *models.HubTxnLock) error
	GetTransactionLockByHubID(ctx context.Context, sHubID string) (*models.HubTxnLock, error)
	DeleteTransactionLockByHubID(ctx context.Context, sHubID string) error
	// Token
	GetTokenByTerminalID(ctx context.Context, sID string) (*models.Token, error)
}

type Store interface {
	Querier
	// Error check
	IsNoRow(err error) bool
}

type sqlStore struct {
	Querier
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return sqlStore{
		Querier: new(db),
		db:      db,
	}
}

func (s sqlStore) IsNoRow(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	return false
}
