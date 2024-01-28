//Для оценки качества доставки писем компания самостоятельно раз в минуту отправляет
//письма каждому провайдеру на почтовые ящики распределенные по всему миру и
//с помощью API проверяет через какое время приходит письмо.
//Письма отправляются от разных провайдеров, чтобы получить медианное время доставки.
//Значение 0 в значении времени доставки означает что письмо не было получено в течение часа

package email

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var allowedProviders = []string{
	"Gmail",
	"Yahoo",
	"Hotmail",
	"Orange",
	"Comcast",
	"AOL",
	"Live",
	"RediffMail",
	"GMX",
	"Proton Mail",
	"Yandex",
	"Mail.ru",
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func New(country string, provider string, deliveryTime int) *EmailData {
	if _, err := iso3166.FromAlpha2(country); err != nil {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}
	return &EmailData{
		Country:      country,
		Provider:     provider,
		DeliveryTime: deliveryTime,
	}
}

func fromSTR(str string) *EmailData {
	listStr := strings.Split(str, ";")
	if len(listStr) < 3 {
		return nil
	}
	deliveryTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil
	}
	return New(listStr[0], listStr[1], deliveryTime)
}
