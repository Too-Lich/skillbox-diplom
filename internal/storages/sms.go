package storages

import (
	"diplom/internal/sms"
	"diplom/pkg/pars"

	"sort"
)

type SMSStorage []*sms.SMSData

func (ss *SMSStorage) Add(obj *sms.SMSData) {
	*ss = append(*ss, obj)
}

func NewSMSStorage(filename string) (*SMSStorage, error) {
	return createSMSStorage(filename)
}

func createSMSStorage(filename string) (*SMSStorage, error) {
	smsStr, err := pars.FileToString(filename)
	if err != nil {
		return nil, err
	}

	ss := SMSStorage{}
	for _, s := range smsStr {
		res := sms.FromSTR(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss, nil
}

func (ss SMSStorage) SortCountry() {
	sortF := func(i, j int) bool {
		return ss[i].Country < ss[j].Country
	}
	sort.SliceStable(ss, sortF)
}

func (ss SMSStorage) SortProvider() {
	sortF := func(i, j int) bool {
		return ss[i].Provider < ss[j].Provider
	}
	sort.SliceStable(ss, sortF)
}
