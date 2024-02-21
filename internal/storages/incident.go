package storages

import (
	"diplom/internal/incident"
	"diplom/pkg/apiRequest"
	"diplom/pkg/pars"
	"fmt"
	"sort"
)

type IncidentStorage []*incident.IncidentData

func NewIncidentStorage(url string) (*IncidentStorage, error) {
	return createIncidentStorage(url)
}

func createIncidentStorage(url string) (*IncidentStorage, error) {
	resp := apiRequest.Get(url)
	is := IncidentStorage{}
	if resp == nil {
		return &is, fmt.Errorf("ошибка получения данных Incident")
	}
	if err := pars.JSON(&is, resp.Body); err != nil {
		return &is, err
	}
	deleteErrData(is)
	return &is, nil
}

func (is IncidentStorage) sort() {
	sortF := func(i, j int) bool {
		return is[i].Status < is[j].Status
	}
	sort.SliceStable(is, sortF)
}
