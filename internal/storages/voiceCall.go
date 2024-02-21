package storages

import (
	"diplom/internal/voiceCall"
	"diplom/pkg/pars"
)

type VCStorage []*voiceCall.VoiceCallData

func (vcs *VCStorage) Add(obj *voiceCall.VoiceCallData) {
	*vcs = append(*vcs, obj)
}

func NewVCStorage(filename string) (*VCStorage, error) {
	return createVCStorage(filename)
}

func createVCStorage(filename string) (*VCStorage, error) {
	smsStr, err := pars.FileToString(filename)
	if err != nil {
		return nil, err
	}
	ss := VCStorage{}
	for _, s := range smsStr {
		res := voiceCall.FromStr(s)
		if res == nil {
			continue
		}
		ss.Add(res)
	}
	return &ss, nil
}
