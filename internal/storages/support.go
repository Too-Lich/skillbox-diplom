package storages

import (
	"diplom/internal/support"
	"diplom/pkg/apiRequest"
	"diplom/pkg/pars"
	"fmt"
	"log"
)

type SupportStorage []*support.Support

func NewSupportStorage(url string) (*SupportStorage, error) {
	return createSupportStorage(url)
}

func createSupportStorage(url string) (*SupportStorage, error) {
	resp := apiRequest.Get(url)
	ss := SupportStorage{}
	if resp == nil {
		return &ss, fmt.Errorf("Ошибка данных Support")
	}
	log.Printf("Чтение url   %v", url)
	if err := pars.JSON(&ss, resp.Body); err != nil {
		return &ss, err
	}
	deleteErrData(ss)
	return &ss, nil
}

func (s SupportStorage) CurrentLoad() (int, int) {
	countTickets := 0
	for _, el := range s {
		countTickets += el.ActiveTickets
	}
	loadStats := 1
	switch {
	case countTickets < 9:
		loadStats = 1
	case countTickets < 16:
		loadStats = 2
	case countTickets > 16:
		loadStats = 3
	}
	avgTimeTicket := 60 / 18
	fullTimeTicket := avgTimeTicket * countTickets
	waitingTime := fullTimeTicket + avgTimeTicket
	return loadStats, waitingTime
}
