// Package mms осуществляет сбор данных о системе MMS
package mms

import (
	"encoding/json"
	"fmt"
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"strconv"
)

var allowedProviders = []string{"Topolo", "Rond", "Kildy"}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func New(country, provider, bandwidthSTR, avgRespTime string) *MMSData {
	c, err := iso3166.FromAlpha2(country)
	if err != nil {
		return nil
	}
	bandwidth, err := strconv.Atoi(bandwidthSTR)
	if err != nil || (bandwidth < 0 && bandwidth > 100) {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}

	/*
		responseTime, err := strconv.Atoi(avgRespTime)
		if err != nil {
			return nil
		}
	*/
	return &MMSData{
		Country:      c.Name(),
		Bandwidth:    bandwidthSTR,
		ResponseTime: avgRespTime,
		Provider:     provider,
	}

}

// Проверка работы MMS
func CheckMMS() {
	fmt.Println("<-- Проверка работы MMS -->")
	resp, _ := http.Get("http://localhost:8383/mms")

	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	var mm []MMSData

	err := json.Unmarshal(b, &mm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range mm {
		fmt.Println(*New(mm[i].Country, mm[i].Provider, mm[i].Bandwidth, mm[i].ResponseTime))
	}
}
