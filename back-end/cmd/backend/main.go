package main

import (
	"backend/internal/pkg/db"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := "sqlserver://paygreen_dev:123@123A@192.168.66.51:1433?database=PayGreen_Dev_Hub"
	conn, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	terminal := db.Terminal{}
	terminalHub := db.TerminalHub{}
	conn.Where("SerieNo = ?", "000141220404274").First(&terminal)
	tx := conn.Where("Id = ?", terminal.TerminalHubId.String()).First(&terminalHub)
	fmt.Println("Error:", tx.Error)
	fmt.Println("Terminal:", terminal)
	fmt.Println("TerminalHub:", terminalHub)
}
