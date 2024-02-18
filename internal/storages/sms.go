package storages

import "diplom/internal/sms"

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
}
