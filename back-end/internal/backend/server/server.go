package server

import (
	"backend/internal/backend/config"
	"backend/internal/pkg/db"
	"backend/internal/pkg/logger"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	route  *gin.Engine
	config *config.Config
	store  *gorm.DB
}

func NewServer(cfg *config.Config, store *gorm.DB) *Server {
	return &Server{
		route:  gin.Default(),
		config: cfg,
		store:  store,
	}
}

func (s *Server) Start() error {
	route := s.route.Use(s.Logger)
	route.GET("/", s.Index)
	route.POST("/create-test-data", s.CreateTestData)
	route.POST("/run-tests", s.RunTests)
	route.DELETE("/delete-data-test/:hub_prefix", s.DeleteTestData)
	return s.route.Run(s.config.ServerAddress)
}

var MajorVer = "1.0.0"
var MinorVer = ""

func (s *Server) Index(ctx *gin.Context) {
	version := strings.Join([]string{MajorVer, MinorVer}, ".")
	sRet := fmt.Sprintf(`Test performance backend
Version: %s`, version)
	ctx.Header("Content-Type", "text/plain")
	// Return the raw string
	ctx.String(http.StatusOK, sRet)
	return
}

func (s *Server) CreateTestData(ctx *gin.Context) {
	var req RequestCreateTestData
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.ErrorFuncf("%s", err)
		ctx.JSON(HttpError(http.StatusBadRequest, err.Error()))
		return
	}
	// Check hub prefix is existed
	hub := db.TerminalHub{}
	tx := s.store.Where("TerminalHubID LIKE ?", req.HubPrefix+"%").First(&hub)
	if errors.Is(tx.Error, nil) {
		logger.ErrorFuncf("Prefix existed")
		ctx.JSON(HttpError(http.StatusBadRequest, "Prefix existed"))
		return
	}
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.ErrorFuncf("%s", tx.Error)
		ctx.JSON(HttpError(http.StatusBadRequest, tx.Error.Error()))
		return
	}
	// Verify shopId, modelId, osVersionId, processorId
	shop := db.Shop{}
	tx = s.store.Where("Id = ?", req.ShopId).First(&shop)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.ErrorFuncf("Invalid shop id: %s", tx.Error)
		ctx.JSON(HttpError(http.StatusBadRequest, fmt.Sprintf("Invalid shop id: %s", tx.Error)))
		return
	}
	model := db.TerminalModel{}
	tx = s.store.Where("Id = ?", req.ModeId).First(&model)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.ErrorFuncf("Invalid model id: %s", tx.Error)
		ctx.JSON(HttpError(http.StatusBadRequest, fmt.Sprintf("Invalid model id: %s", tx.Error)))
		return
	}
	osVer := db.OsVersion{}
	tx = s.store.Where("Id = ?", req.OsVersionId).First(&osVer)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.ErrorFuncf("Invalid os version id: %s", tx.Error)
		ctx.JSON(HttpError(http.StatusBadRequest, fmt.Sprintf("Invalid os version id: %s", tx.Error)))
		return
	}
	processor := db.Processor{}
	tx = s.store.Where("Id = ?", req.ProcessorId).First(&processor)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.ErrorFuncf("Invalid prcessor id: %s", tx.Error)
		ctx.JSON(HttpError(http.StatusBadRequest, fmt.Sprintf("Invalid prcessor id:: %s", tx.Error)))
		return
	}
	// Create terminal hub, terminal and api token
	for i := 0; i < req.NumberOfHub; i++ {
		id := uuid.New().String()
		hubId := mssql.UniqueIdentifier{}
		_ = hubId.Scan(id)
		// Create hub
		hub = db.TerminalHub{
			Id: hubId,
			TerminalHubID: sql.NullString{
				String: req.HubPrefix + fmt.Sprintf("%04d", i),
				Valid:  true,
			},
			ProcessorId: sql.NullInt64{
				Int64: int64(req.ProcessorId),
				Valid: true,
			},
			Cassette1Denom: sql.NullInt64{
				Int64: 20,
				Valid: true,
			},
			Cassette2Denom: sql.NullInt64{
				Int64: 5,
				Valid: true,
			},
			Cassette1Remain: sql.NullInt64{
				Int64: 1800,
				Valid: true,
			},
			Cassette2Remain: sql.NullInt64{
				Int64: 1800,
				Valid: true,
			},
			TotalLowAmount: sql.NullInt64{
				Int64: 1000,
				Valid: true,
			},
			WorkingKey: sql.NullString{},
			WorkingKeyInterval: sql.NullInt64{
				Int64: 4,
				Valid: true,
			},
			HealthCheckInterval: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			TxnDelayMin: sql.NullInt64{
				Int64: 30,
				Valid: true,
			},
			TxnDelayMax: sql.NullInt64{
				Int64: 40,
				Valid: true,
			},
			Surcharge: sql.NullFloat64{
				Float64: 2.5,
				Valid:   true,
			},
			Status: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			Deleted: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
			CreatedDate: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			RequestTimeoutProcessor: sql.NullInt64{
				Int64: 90,
				Valid: true,
			},
			RequestTimeoutTerminal: sql.NullInt64{
				Int64: 90,
				Valid: true,
			},
			CurrentTSQ: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			LastRunWorkingKey: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			LastRunHealthCheck: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		}
		tx = s.store.Create(&hub)
		if tx.Error != nil {
			logger.ErrorFuncf("[%s]Create hub failed", hub.TerminalHubID.String)
			continue
		}
		// Create terminals linked to hub
		for j := 0; j < req.TerminalPerHub; j++ {
			id = uuid.New().String()
			termId := mssql.UniqueIdentifier{}
			shopId := mssql.UniqueIdentifier{}
			modelId := mssql.UniqueIdentifier{}
			osId := mssql.UniqueIdentifier{}
			_ = termId.Scan(id)
			_ = shopId.Scan(req.ShopId)
			_ = modelId.Scan(req.ModeId)
			_ = osId.Scan(req.OsVersionId)
			terminalId := fmt.Sprintf("%s%04d", hub.TerminalHubID.String, j)
			terminalName := fmt.Sprintf("N%s%04d", hub.TerminalHubID.String, j)
			serieNo := fmt.Sprintf("S%s%04d", hub.TerminalHubID.String, j)
			term := db.Terminal{
				Id:     termId,
				ShopId: shopId,
				TerminalID: sql.NullString{
					String: terminalId,
					Valid:  true,
				},
				TerminalName: sql.NullString{
					String: terminalName,
					Valid:  true,
				},
				SerieNo: sql.NullString{
					String: serieNo,
					Valid:  true,
				},
				TerminalModelId: modelId,
				OsVersionId:     osId,
				Address1:        sql.NullString{},
				Address2:        sql.NullString{},
				City:            sql.NullString{},
				State:           sql.NullString{},
				Location:        sql.NullString{},
				NetworkType: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				CreatedDate: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				ActivedDate: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				LastModifiedDate: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				Deleted: sql.NullBool{
					Bool:  false,
					Valid: true,
				},
				Status: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				IsActive: sql.NullBool{
					Bool:  true,
					Valid: true,
				},
				Zipcode:       sql.NullString{},
				ProcessorTID:  sql.NullString{},
				ProcessorId:   sql.NullInt64{},
				DeviceID:      sql.NullString{},
				TerminalHubId: hubId,
			}
			tx = s.store.Create(&term)
			if tx.Error != nil {
				logger.ErrorFuncf("[%s]Create terminal failed: %s", hub.TerminalHubID.String, term.TerminalID.String)
				continue
			}
			// Create token for this terminal
			token := db.Token{
				ID:         sql.NullInt64{},
				TerminalID: term.Id,
				UserID:     sql.NullString{},
				Token: sql.NullString{
					String: "AAAAAAFmMGQtYTVlZC00M2EyLWIyOGUtZTg4ODJmMzdlZThh",
					Valid:  true,
				},
				FirebaseToken: sql.NullString{},
				IP: sql.NullString{
					String: "localhost",
					Valid:  true,
				},
				CreateDate: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				ExpiryDate: sql.NullTime{
					Time:  time.Now().AddDate(10, 0, 0),
					Valid: true,
				},
			}
			tx = s.store.Create(&token)
			if tx.Error != nil {
				logger.ErrorFuncf("[%s]Create token failed: %s", hub.TerminalHubID.String, term.TerminalID.String)
				continue
			}
		}
	}
	ctx.JSON(HttpError(http.StatusOK, "Success"))
	return
}

func (s *Server) RunTests(ctx *gin.Context) {
	var req RequestRunTests
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.ErrorFuncf("%s", err)
		ctx.JSON(HttpError(http.StatusBadRequest, err.Error()))
		return
	}
	// Check mode run: duration or iteration
	// Duration: run for period of time
	// Iteration: send number of requests then stop
	if req.Duration < 1 && req.Iteration < 1 {
		logger.ErrorFuncf("Need duration > 1 or iteration > 1")
		ctx.JSON(HttpError(http.StatusBadRequest, "Need duration > 1 or iteration > 1"))
		return
	}
	// Get all hub
	hubs := []db.TerminalHub{}
	tx := s.store.Where("TerminalHubID LIKE ?", req.HubPrefix+"%").Find(&hubs)
	if len(hubs) == 0 {
		logger.ErrorFuncf("Prefix doesn't exist")
		ctx.JSON(HttpError(http.StatusBadRequest, "Prefix doesn't exist"))
		return
	}
	if !errors.Is(tx.Error, nil) {
		logger.ErrorFuncf("%s", err.Error())
		ctx.JSON(HttpError(http.StatusBadRequest, err.Error()))
		return
	}
	// Loop all hubs to run
	for _, hub := range hubs {
		go func(terminalHub db.TerminalHub) {
			// Get all terminal linked to hub
			logger.InfoFuncf("[%s]Start goroutine for hub", terminalHub.TerminalHubID.String)
			terminals := []db.Terminal{}
			tx = s.store.Where("TerminalHubId = ?", terminalHub.Id.String()).Find(&terminals)
			if tx.Error != nil {
				logger.ErrorFuncf("[%s]Find terminal error: %s", terminalHub.TerminalHubID.String, err)
				return
			}
			for _, terminal := range terminals {
				go func(hub db.TerminalHub, terminal db.Terminal) {
					logger.InfoFuncf("[%s]Start goroutine for terminal %s", terminalHub.TerminalHubID.String, terminal.TerminalID.String)
					startTime := time.Now()
					count := 0
					// 1. Call get working key
					for {
						err := s.GetWorkingKey(terminalHub, terminal, req)
						if err != nil {
							logger.ErrorFuncf("[%s]Get working key for %s failed: %s", terminalHub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
						} else {
							// 2. Call process transaction and process transaction check if need
							err = s.ProcessTransaction(terminalHub, terminal, req)
							if err != nil {
								logger.ErrorFuncf("[%s]Process transaction for %s failed: %s", terminalHub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
							}
						}
						count++
						if req.DelayBetweenRequests > 0 {
							time.Sleep(time.Duration(req.DelayBetweenRequests) * time.Second)
						}
						if req.Duration > 0 {
							if int(time.Now().Sub(startTime).Minutes()) > req.Duration {
								logger.InfoFuncf("[%s]Duration stop goroutine for terminal %s", terminalHub.TerminalHubID.String, terminal.TerminalID.String)
								break
							}
						}
						if req.Iteration == count {
							logger.InfoFuncf("[%s]Iteration stop goroutine for terminal %s", terminalHub.TerminalHubID.String, terminal.TerminalID.String)
							break
						}
					}
				}(terminalHub, terminal)
				if req.DelayBetweenTerminals > 0 {
					time.Sleep(time.Duration(req.DelayBetweenTerminals) * time.Second)
				}
			}
		}(hub)
		if req.DelayBetweenHubs > 0 {
			time.Sleep(time.Duration(req.DelayBetweenHubs) * time.Second)
		}
	}
	ctx.JSON(HttpError(http.StatusOK, "Success"))
	return
}

func (s *Server) DeleteTestData(ctx *gin.Context) {
	// Get serial number
	var uri UriDeleteTestData
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		logger.ErrorFuncf("%s", err)
		ctx.JSON(HttpError(http.StatusBadRequest, err.Error()))
		return
	}
	// Get all hub
	hubs := []db.TerminalHub{}
	tx := s.store.Where("TerminalHubID LIKE ?", uri.HubPrefix+"%").Find(&hubs)
	if len(hubs) == 0 {
		logger.ErrorFuncf("Prefix doesn't exist")
		ctx.JSON(HttpError(http.StatusBadRequest, "Prefix doesn't exist"))
		return
	}
	if !errors.Is(tx.Error, nil) {
		logger.ErrorFuncf("%s", err.Error())
		ctx.JSON(HttpError(http.StatusBadRequest, err.Error()))
		return
	}
	// Loop all hubs to delete
	for _, hub := range hubs {
		// Get all terminal linked to hub
		terminals := []db.Terminal{}
		tx = s.store.Where("TerminalHubId = ?", hub.Id.String()).Find(&terminals)
		if tx.Error != nil {
			logger.ErrorFuncf("[%s]Find terminal error: %s", hub.TerminalHubID, err)
			continue
		}
		for _, terminal := range terminals {
			tx = s.store.Where("TerminalID = ?", terminal.Id.String()).Delete(&db.Token{})
			if tx.Error != nil {
				logger.ErrorFuncf("[%s]Delete token for %s error: %s", hub.TerminalHubID, terminal.TerminalID, err)
			}
			tx = s.store.Where("Id = ?", terminal.Id.String()).Delete(&db.Terminal{})
			if tx.Error != nil {
				logger.ErrorFuncf("[%s]Delete terminal %s error: %s", hub.TerminalHubID, terminal.TerminalID, err)
			}
		}
		// Delete hub
		tx = s.store.Where("Id = ?", hub.Id.String()).Delete(&db.TerminalHub{})
		if tx.Error != nil {
			logger.ErrorFuncf("[%s]Delete hub error: %s", hub.TerminalHubID, err)
		}
	}
	ctx.JSON(HttpError(http.StatusOK, "Success"))
	return
}
