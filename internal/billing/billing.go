//у компании есть автоматизированная система биллинга для реализации оплаты услуг
//в ручном и автоматизированном режиме. Команда биллинга контролирует свои сервисы
//как простое состояние работает / не работает и для экономии места использует битовую маску

package billing

import (
	"errors"
	"log"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func New(path string) (*BillingData, error) {
	content, err := ReadFile(path)
	if len(content) != 6 || err != nil {
		return nil, errors.New("Неправильные входные данные")
	}
	return &BillingData{
		CreateCustomer: check(content[5]),
		Purchase:       check(content[4]),
		Payout:         check(content[3]),
		Recurring:      check(content[2]),
		FraudControl:   check(content[1]),
		CheckoutPage:   check(content[0]),
	}, nil
}

func ReadFile(filePath string) ([]byte, error) {
	log.Printf("Чтение файла `%v`", filePath)
	content, err := os.ReadFile(filePath)
	return content, err
}

func check(status byte) bool {
	return status == byte('1')
}
