package main

import (
	"backend/internal/backend/config"
	"backend/internal/backend/server"
	"backend/internal/pkg/logger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	err := logger.NewLogger("./configs/logcfg.json")
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig("./configs")
	_ = cfg
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	dsn := "sqlserver://paygreen_dev:123@123A@192.168.66.51:1433?database=PayGreen_Dev_Hub"
	conn, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	srv := server.NewServer(cfg, conn)
	err = srv.Start()
	if err != nil {
		panic(err)
	}
	//terminal := db.Terminal{}
	//terminalHub := db.TerminalHub{}
	//token := db.Token{}
	//conn.Where("SerieNo = ?", "000141220404274").First(&terminal)
	//conn.Where("Id = ?", terminal.TerminalHubId.String()).First(&terminalHub)
	//tx := conn.Where("TerminalID = ?", terminal.Id.String()).First(&token)
	//fmt.Println("Error:", tx.Error)
	//fmt.Println("Terminal:", terminal)
	//fmt.Println("TerminalHub:", terminalHub)
	//fmt.Println("Token:", token)

}
