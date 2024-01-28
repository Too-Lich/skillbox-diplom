/*
Package voiceCall осуществляет сбор данных о системе тестовых звонков.
Общий принцип работы Voice Call похож на SMS, но в данном случае он имеет более широкий набор значений.
Компания старается улучшать качество системы, поэтому делает тестовые звонки,
позволяет записывать разговоры и оценивает параметры влияющие на разговор
*/
package voiceCall

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var allowedProviders = []string{"TransparentCalls", "E-Voice", "JustPhone"}

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           int     `json:"bandwidth"`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianCallsTime     int     `json:"median_calls_time"`
}

func New(country string, bandwidth, responseTime int, provider string, conStability float32, ttfb, voicePurity, medCallTime int) *VoiceCallData {
	if _, err := iso3166.FromAlpha2(country); err != nil {
		return nil
	}
	if bandwidth < 0 && bandwidth > 100 {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}
	return &VoiceCallData{
		Country:             country,
		Bandwidth:           bandwidth,
		ResponseTime:        responseTime,
		Provider:            provider,
		ConnectionStability: conStability,
		TTFB:                ttfb,
		VoicePurity:         voicePurity,
		MedianCallsTime:     medCallTime,
	}
}

func FromStr(str string) *VoiceCallData {
	listStr := strings.Split(str, ";")
	if len(listStr) < 8 {
		return nil
	}
	bandwidth, err := strconv.Atoi(listStr[1])
	if err != nil {
		return nil
	}
	responseTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil
	}
	conStability, err := strconv.ParseFloat(listStr[4], 32)
	if err != nil {
		return nil
	}
	ttfb, err := strconv.Atoi(listStr[5])
	if err != nil {
		return nil
	}
	voicePurity, err := strconv.Atoi(listStr[6])
	if err != nil {
		return nil
	}
	medCallTime, err := strconv.Atoi(listStr[7])
	if err != nil {
		return nil
	}

	return New(listStr[0], bandwidth, responseTime, listStr[3], float32(conStability), ttfb, voicePurity, medCallTime)
}
