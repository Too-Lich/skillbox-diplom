package storages

import (
	"diplom/internal/email"
	"diplom/pkg/pars"
	"sort"
)

type EmailStorage []*email.EmailData

func (es *EmailStorage) Add(obj *email.EmailData) {
	*es = append(*es, obj)
}

func NewEmailStorage(filename string) (*EmailStorage, error) {
	return createEmailStorage(filename)
}

func createEmailStorage(filename string) (*EmailStorage, error) {
	emlStr, err := pars.FileToString(filename)
	if err != nil {
		return nil, err
	}

	es := EmailStorage{}
	for _, s := range emlStr {
		res := email.FromSTR(s)
		if res == nil {
			continue
		}
		es.Add(res)
	}

	return &es, nil
}

func (es EmailStorage) catalogingByCountry() map[string]EmailStorage {
	emails := map[string]EmailStorage{}
	for _, email := range es {
		if _, ok := emails[email.Country]; ok {
			emails[email.Country] = EmailStorage{}
		}
		emails[email.Country] = append(emails[email.Country], email)
	}
	return emails
}

type Provider struct {
	Name            string `json:"name"`
	avgDeliveryTime float32
	countEmail      int
}

type providerStorage []*Provider

func (es EmailStorage) createStatisticProviders() providerStorage {
	providers := map[string]*Provider{}
	for _, el := range es {
		provider, ok := providers[el.Provider]
		if !ok {
			providers[el.Provider] = &Provider{
				Name:            el.Provider,
				avgDeliveryTime: float32(el.DeliveryTime),
				countEmail:      1,
			}
			continue
		}
		provider.avgDeliveryTime += float32(el.DeliveryTime)
		provider.countEmail++
	}
	providersList := providerStorage{}
	for _, provider := range providers {
		provider.avgDeliveryTime = provider.avgDeliveryTime / float32(provider.countEmail)
		providersList = append(providersList, provider)
	}
	return providersList
}

func (ps providerStorage) sort() {
	sortF := func(i, j int) bool {
		return ps[i].avgDeliveryTime < ps[j].avgDeliveryTime
	}
	sort.SliceStable(ps, sortF)
}

func (ps providerStorage) BestAndWorst() []providerStorage {
	ps.sort()
	countTop := 3
	if len(ps) < countTop {
		countTop = len(ps)
	}
	var bestProviders, worstProviders providerStorage
	for i := 0; i < countTop; i++ {
		bestProviders = append(bestProviders, ps[i])
		worstProviders = append(worstProviders, ps[len(ps)-1-i])
	}
	return []providerStorage{bestProviders, worstProviders}
}
