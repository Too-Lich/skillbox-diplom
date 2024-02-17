package storages

import (
	"diplom/internal/mms"
	"encoding/json"
	"io"
	"net/http"
)

type MMSStorage []*mms.MMSData

func (ms *MMSStorage) Add(obj *mms.MMSData) {
	*ms = append(*ms, obj)
}

func NewMMSStorage(url string) (*MMSStorage, error) {
	return createMMSStorage(url)
}

func createMMSStorage(url string) (*MMSStorage, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var buf []map[string]interface{}
	if err = json.Unmarshal(content, &buf); err != nil {
		return nil, err
	}
	ms := MMSStorage{}
	for _, elem := range buf {
		m := mms.New(elem["country"].(string), elem["provider"].(string), elem["bandwidth"].(string), elem["response_time"].(string))
		if m != nil {
			ms = append(ms, m)
		}
	}

	return &ms, nil
}
