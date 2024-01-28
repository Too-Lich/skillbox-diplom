// Package mms осуществляет сбор данных о системе MMS
package mms

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
)

var allowedProviders = []string{"Topolo", "Rond", "Kildy"}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    int    `json:"bandwidth"`
	ResponseTime int    `json:"response_time"`
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
	ResponseTime, err := strconv.Atoi(avgRespTime)
	if err != nil {
		return nil
	}
	return &MMSData{
		Country:      c.Name(),
		Bandwidth:    bandwidth,
		ResponseTime: ResponseTime,
		Provider:     provider,
	}
}
