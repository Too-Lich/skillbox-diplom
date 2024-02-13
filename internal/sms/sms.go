// Package sms осуществляет сбор данных о системе SMS
package sms

import (
	"fmt"
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
	"strings"
)

// Список допустимы провайдеров в массиве данных
var allowedProviders = []string{"Topolo", "Rond", "Kildy"}

// SMSData - структура для данных SMS
type SMSData struct {
	Country     string `json:"country"`
	Bandwidth   int    `json:"bandwidth"`
	AvgRespTime int    `json:"avg_resp_time"`
	Provider    string `json:"provider"`
}

// New - каждая строка должна содержать 4 поля (alpha-2 код страны, пропускная способность канала от 0% до 100%, среднее время ответа в ms, название компании провайдера
func New(country string, bandwidth int, avgRespTime int, provider string) *SMSData {
	c, err := iso3166.FromAlpha2(country)
	if err != nil {
		return nil
	}

	if bandwidth < 0 && bandwidth > 100 {
		return nil
	}

	if !slices.Contains(allowedProviders, provider) {
		return nil
	}

	return &SMSData{
		Country:     c.Name(),
		Bandwidth:   bandwidth,
		AvgRespTime: avgRespTime,
		Provider:    provider,
	}
}

func FromSTR(str string) *SMSData {
	listStr := strings.Split(str, ";")

	if len(listStr) < 4 {
		return nil
	}

	bandwidth, err := strconv.Atoi(listStr[1])
	if err != nil {
		return nil
	}

	avgRespTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil

	}

	return New(listStr[0], bandwidth, avgRespTime, listStr[3])
}

// Проверка работы функции FromSTR
func CheckSms() {
	rawData, _ := os.ReadFile("../sms.data")
	strData := strings.Split(string(rawData), "\n")
	for i := range strData {
		data := FromSTR(strData[i])
		fmt.Println(data)
	}
}
