package storages

import (
	"diplom/configs"
	"diplom/internal/billing"
)

type ResultSetT struct {
	SMS       []SMSStorage               `json:"sms"`
	MMS       []MMSStorage               `json:"mms"`
	Email     map[string]providerStorage `json:"email"`
	VoiceCall VCStorage                  `json:"voice_call"`
	Billing   billing.BillingData        `json:"billing"`
	Support   []int                      `json:"support"`
	Incidents IncidentStorage            `json:"incident"`
}

type ResultT struct {
	Status bool        `json:"status"` // True - если все этапы сбора данных прошли. Иначе = false
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора данных прошли. Иначе = nil
	Error  []string    `json:"errors"` // Пустая строка, если все этапы сбора данных прошли. Иначе - тест ошибки
}

var config = configs.GetConfig()

func GetResultData() ResultT {
	data := ResultSetT{}
	status := ResultT{
		Status: true,
		Data:   &data,
		Error:  nil,
	}
	return status
}
