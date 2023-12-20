package db

import (
	"code/internal/pkg/models"
	"context"
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)

func (q Queries) GetWorkingKeyByTerminalHubID(ctx context.Context, sID string) (string, error) {
	query := `SELECT WorkingKey FROM dbo.tblTerminalHub WHERE Id = @p1`
	var sWorkingKey string
	var uuidTerminalHubID mssql.UniqueIdentifier
	err := uuidTerminalHubID.Scan(sID)
	if err != nil {
		return "", err
	}
	err = q.db.QueryRowContext(ctx, query, uuidTerminalHubID).Scan(&sWorkingKey)
	if err != nil {
		return "", err
	}
	return sWorkingKey, nil
}

func (q Queries) UpdateWorkingKeyForTerminalHub(ctx context.Context, sID string, sWorkingKey string) error {
	query := `UPDATE dbo.tblTerminalHub SET WorkingKey = @p1 WHERE Id = @p2`
	var uuidTerminalHubID mssql.UniqueIdentifier
	err := uuidTerminalHubID.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query, sWorkingKey, uuidTerminalHubID)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) GetTerminalHub(ctx context.Context, sID string) (*models.TerminalHub, error) {
	query := `SELECT "id", 
       "terminalhubid", 
       "processorid", 
       "cassette1denom", 
       "cassette2denom", 
       "cassette1remain", 
       "cassette2remain", 
       "totallowamount", 
       "workingkey", 
       "workingkeyinterval", 
       "healthcheckinterval", 
       "txndelaymin", 
       "txndelaymax", 
       "surcharge", 
       "status", 
       "deleted", 
       "createddate", 
       "requesttimeoutprocessor", 
       "requesttimeoutterminal", 
       "currenttsq",
       "LastRunWorkingKey",
       "LastRunHealthCheck"
FROM dbo.tblTerminalHub WHERE Id = @p1`
	var terminalHub TerminalHub
	var uuidTerminalHubID mssql.UniqueIdentifier
	err := uuidTerminalHubID.Scan(sID)
	if err != nil {
		return nil, err
	}
	err = q.db.QueryRowContext(ctx, query, uuidTerminalHubID).Scan(&terminalHub.ID,
		&terminalHub.TerminalHubID,
		&terminalHub.ProcessorId,
		&terminalHub.Cassette1Denom,
		&terminalHub.Cassette2Denom,
		&terminalHub.Cassette1Remain,
		&terminalHub.Cassette2Remain,
		&terminalHub.TotalLowAmount,
		&terminalHub.WorkingKey,
		&terminalHub.WorkingKeyInterval,
		&terminalHub.HealthCheckInterval,
		&terminalHub.TxnDelayMin,
		&terminalHub.TxnDelayMax,
		&terminalHub.Surcharge,
		&terminalHub.Status,
		&terminalHub.Deleted,
		&terminalHub.CreatedDate,
		&terminalHub.RequestTimeoutProcessor,
		&terminalHub.RequestTimeoutTerminal,
		&terminalHub.CurrentTSQ,
		&terminalHub.LastRunWorkingKey,
		&terminalHub.LastRunHealthCheck)
	if err != nil {
		return nil, err
	}
	// Map data
	modelHub := models.TerminalHub{
		ID:                      terminalHub.ID.String(),
		TerminalHubID:           terminalHub.TerminalHubID.String,
		ProcessorId:             terminalHub.ProcessorId.Int64,
		Cassette1Denom:          terminalHub.Cassette1Denom.Int64,
		Cassette2Denom:          terminalHub.Cassette2Denom.Int64,
		Cassette1Remain:         terminalHub.Cassette1Remain.Int64,
		Cassette2Remain:         terminalHub.Cassette2Remain.Int64,
		TotalLowAmount:          terminalHub.TotalLowAmount.Int64,
		WorkingKey:              terminalHub.WorkingKey.String,
		WorkingKeyInterval:      terminalHub.WorkingKeyInterval.Int64,
		HealthCheckInterval:     terminalHub.HealthCheckInterval.Int64,
		TxnDelayMin:             terminalHub.TxnDelayMin.Int64,
		TxnDelayMax:             terminalHub.TxnDelayMax.Int64,
		Surcharge:               terminalHub.Surcharge.Float64,
		Status:                  terminalHub.Status.Bool,
		Deleted:                 terminalHub.Deleted.Bool,
		CreatedDate:             terminalHub.CreatedDate.Time,
		RequestTimeoutProcessor: terminalHub.RequestTimeoutProcessor.Int64,
		RequestTimeoutTerminal:  terminalHub.RequestTimeoutTerminal.Int64,
		CurrentTSQ:              terminalHub.CurrentTSQ.Int64,
		LastRunWorkingKey:       terminalHub.LastRunWorkingKey.Time,
		LastRunHealthCheck:      terminalHub.LastRunHealthCheck.Time,
	}
	_, offset := time.Now().Zone()
	modelHub.LastRunWorkingKey = modelHub.LastRunWorkingKey.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	modelHub.LastRunHealthCheck = modelHub.LastRunHealthCheck.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
	return &modelHub, nil
}

func (q Queries) UpdateTerminalHubCassette(ctx context.Context, sID string, nCassette1Remain int64, nCassette2Remain int64) error {
	query := `UPDATE dbo.tblTerminalHub SET Cassette1Remain = @p1,Cassette2Remain = @p2 WHERE Id=@p3`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query, nCassette1Remain, nCassette2Remain, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) UpdateTerminalHubTSQ(ctx context.Context, sID string, nTSQ int64) error {
	query := `UPDATE dbo.tblTerminalHub SET CurrentTSQ = @p1 WHERE Id=@p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query, nTSQ, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) GetAllTerminalHub(ctx context.Context) ([]models.TerminalHub, error) {
	query := `SELECT "Id", "TerminalHubID", "ProcessorId", "Cassette1Denom", "Cassette2Denom", "Cassette1Remain", "Cassette2Remain", "TotalLowAmount", "WorkingKey", "WorkingKeyInterval", "HealthCheckInterval", "TxnDelayMin", "TxnDelayMax", "Surcharge", "Status", "Deleted", "CreatedDate", "RequestTimeoutProcessor", "RequestTimeoutTerminal", "CurrentTSQ", "LastRunWorkingKey", "LastRunHealthCheck"
				FROM dbo.tblTerminalHub WHERE Status = 1 AND Deleted = 0`
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var terminalHubs []models.TerminalHub
	for rows.Next() {
		var terminalHub TerminalHub
		err = rows.Scan(&terminalHub.ID,
			&terminalHub.TerminalHubID,
			&terminalHub.ProcessorId,
			&terminalHub.Cassette1Denom,
			&terminalHub.Cassette2Denom,
			&terminalHub.Cassette1Remain,
			&terminalHub.Cassette2Remain,
			&terminalHub.TotalLowAmount,
			&terminalHub.WorkingKey,
			&terminalHub.WorkingKeyInterval,
			&terminalHub.HealthCheckInterval,
			&terminalHub.TxnDelayMin,
			&terminalHub.TxnDelayMax,
			&terminalHub.Surcharge,
			&terminalHub.Status,
			&terminalHub.Deleted,
			&terminalHub.CreatedDate,
			&terminalHub.RequestTimeoutProcessor,
			&terminalHub.RequestTimeoutTerminal,
			&terminalHub.CurrentTSQ,
			&terminalHub.LastRunWorkingKey,
			&terminalHub.LastRunHealthCheck)
		if err != nil {
			continue
		}
		// Map data
		modelHub := models.TerminalHub{
			ID:                      terminalHub.ID.String(),
			TerminalHubID:           terminalHub.TerminalHubID.String,
			ProcessorId:             terminalHub.ProcessorId.Int64,
			Cassette1Denom:          terminalHub.Cassette1Denom.Int64,
			Cassette2Denom:          terminalHub.Cassette2Denom.Int64,
			Cassette1Remain:         terminalHub.Cassette1Remain.Int64,
			Cassette2Remain:         terminalHub.Cassette2Remain.Int64,
			TotalLowAmount:          terminalHub.TotalLowAmount.Int64,
			WorkingKey:              terminalHub.WorkingKey.String,
			WorkingKeyInterval:      terminalHub.WorkingKeyInterval.Int64,
			HealthCheckInterval:     terminalHub.HealthCheckInterval.Int64,
			TxnDelayMin:             terminalHub.TxnDelayMin.Int64,
			TxnDelayMax:             terminalHub.TxnDelayMax.Int64,
			Surcharge:               terminalHub.Surcharge.Float64,
			Status:                  terminalHub.Status.Bool,
			Deleted:                 terminalHub.Deleted.Bool,
			CreatedDate:             terminalHub.CreatedDate.Time,
			RequestTimeoutProcessor: terminalHub.RequestTimeoutProcessor.Int64,
			RequestTimeoutTerminal:  terminalHub.RequestTimeoutTerminal.Int64,
			CurrentTSQ:              terminalHub.CurrentTSQ.Int64,
			LastRunWorkingKey:       terminalHub.LastRunWorkingKey.Time,
			LastRunHealthCheck:      terminalHub.LastRunHealthCheck.Time,
		}
		_, offset := time.Now().Zone()
		modelHub.LastRunWorkingKey = modelHub.LastRunWorkingKey.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
		modelHub.LastRunHealthCheck = modelHub.LastRunHealthCheck.Add(-time.Duration(offset) * time.Second).In(time.Now().Location())
		terminalHubs = append(terminalHubs, modelHub)
	}
	return terminalHubs, nil
}

func (q Queries) UpdateLastRunWorkingKey(ctx context.Context, sID string, sLastRun time.Time) error {
	query := `UPDATE dbo.tblTerminalHub SET 
                                 LastRunWorkingKey = @p1
                             WHERE Id = @p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		sLastRun,
		uuid)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) UpdateLastRunHealthCheck(ctx context.Context, sID string, sLastRun time.Time) error {
	query := `UPDATE dbo.tblTerminalHub SET 
                                 LastRunHealthCheck = @p1
                             WHERE Id = @p2`
	var uuid mssql.UniqueIdentifier
	err := uuid.Scan(sID)
	if err != nil {
		return err
	}
	_, err = q.db.ExecContext(ctx, query,
		sLastRun,
		uuid)
	if err != nil {
		return err
	}
	return nil
}
