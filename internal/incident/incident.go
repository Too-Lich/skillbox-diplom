//Сбор данных о системе истории инцидентов
//В общем потоке данных система истории инцидентов просто содержит массив инцидентов за последнюю неделю.
//Мы будем возвращать нашим пользователям историю инцидентов, чтобы разгрузить саппорт
//от однообразных вопросов. Например, у нас происходят неполадки в системе SMS,
//саппорт создает инцидент. Пользователи на нашей странице Statuspage могут посмотреть что у нас неполадки,
//и на какое-то время отключить свои системы, до завершения работ по восстановлению не создавать новых
//тикетов саппорту. Таким образом мы сообщаем нашим пользователям о возможных и бывших проблемах с сервисами.

package incident

import "slices"

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}

func (i *IncidentData) Check() bool {
	return i.Topic != "" && i.checkStatus()
}

func (i *IncidentData) checkStatus() bool {
	possibleStatus := []string{"closed", "active"}
	return slices.Contains(possibleStatus, i.Status)
}
