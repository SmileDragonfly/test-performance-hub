package server

import (
	"backend/internal/pkg/db"
	"backend/internal/pkg/logger"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func (s Server) GetWorkingKey(hub db.TerminalHub, terminal db.Terminal, req RequestRunTests) error {
	url := req.HubUrl + "/GetWorkingKeys/" + terminal.SerieNo.String
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		logger.ErrorFuncf("[%s]Get working key for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	defer resp.Body.Close()
	respBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorFuncf("[%s]io.ReadAll(resp.Body) for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	// Parse response
	var respJson ResponseGetWorkingKey
	err = json.Unmarshal(respBuf, &respJson)
	if err != nil {
		logger.ErrorFuncf("[%s]json.Unmarshal(respBuf, &respJson) for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	if !respJson.Success {
		logger.ErrorFuncf("[%s]respJson.Success for %s failed", hub.TerminalHubID.String, terminal.TerminalID.String)
		return errors.New("Get working key failed")
	}
	return nil
}

func (s Server) ProcessTransaction(hub db.TerminalHub, terminal db.Terminal, req RequestRunTests) error {
	url := req.HubUrl + "/ProcessTransaction"
	org, tip, cashback, total := randomAmountWithTipAndCashback()
	reqHub := RequestPayment{
		Total:          int64(total) * 100,
		Track2:         terminal.TerminalName.String,
		SN:             terminal.SerieNo.String,
		Pan:            terminal.TerminalID.String,
		Track1:         terminal.ShopId.String(),
		Emv:            "",
		Pin:            terminal.TerminalModelId.String(),
		Track3:         terminal.OsVersionId.String(),
		OriginalAmount: strconv.Itoa(org * 100),
		CashBack:       strconv.Itoa(cashback * 100),
		TipAmount:      strconv.Itoa(tip * 100),
	}
	reqHubBuf, _ := json.Marshal(&reqHub)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(reqHubBuf))
	if err != nil {
		logger.ErrorFuncf("[%s]Process transaction for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	defer resp.Body.Close()
	respBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorFuncf("[%s]io.ReadAll(resp.Body) for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	// Parse response
	var respJson ResponsePayment
	err = json.Unmarshal(respBuf, &respJson)
	if err != nil {
		logger.ErrorFuncf("[%s]json.Unmarshal(respBuf, &respJson) for % error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
		return err
	}
	if respJson.Code != 0 && respJson.Code != 998 {
		logger.ErrorFuncf("[%s]Process transaction for %s failed: %s", hub.TerminalHubID.String, terminal.TerminalID.String, respJson.CustomMessage)
		return errors.New("Process transaction fail")
	}
	if respJson.Code == 998 {
		urlCheck := req.HubUrl + "/ProcessTransactionCheck"
		respCheck, err := client.Post(urlCheck, "application/json", bytes.NewBuffer(reqHubBuf))
		if err != nil {
			logger.ErrorFuncf("[%s]Process transaction check for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
			return err
		}
		defer resp.Body.Close()
		respCheckBuf, err := io.ReadAll(respCheck.Body)
		if err != nil {
			logger.ErrorFuncf("[%s]io.ReadAll(respCheck.Body) for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
			return err
		}
		// Parse response
		var respCheckJson ResponsePayment
		err = json.Unmarshal(respCheckBuf, &respCheckJson)
		if err != nil {
			logger.ErrorFuncf("[%s]json.Unmarshal(respCheckBuf, &respCheckJson) for %s error: %s", hub.TerminalHubID.String, terminal.TerminalID.String, err.Error())
			return err
		}
		if respCheckJson.Code != 0 {
			logger.ErrorFuncf("[%s]Process transaction check for %s failed: %s", hub.TerminalHubID.String, terminal.TerminalID.String, respJson.CustomMessage)
			return errors.New("Process transaction check fail")
		}
		logger.InfoFuncf("[%s]Process transaction check for %s success", hub.TerminalHubID.String, terminal.TerminalID.String)
		return nil
	}
	logger.InfoFuncf("[%s]Process transaction for %s success", hub.TerminalHubID.String, terminal.TerminalID.String)
	return nil
}

func randomAmountWithTipAndCashback() (int, int, int, int) {
	// Sinh ngẫu nhiên một số tiền (amount)
	amount := rand.Intn(200) + 1 // Để giả sử số tiền từ 1 đến 200
	// Sinh ngẫu nhiên một số tiền tip
	tip := rand.Intn(20) // Để giả sử số tiền tip từ 0 đến 19
	// Tính số tiền cashback để tổng cộng chia hết cho 5
	cashback := 5 - (amount+tip)%5
	total := amount + tip + cashback
	return amount, tip, cashback, total
}
