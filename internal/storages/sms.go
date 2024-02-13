package storages

import "diplom/internal/sms"

type SMSStorage []*sms.SMSData

func (ss *SMSStorage) Add(obj *sms.SMSData) {
	*ss = append(*ss, obj)
}
