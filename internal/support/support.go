//Функция, которая получает данные о текущей загрузке команды службы поддержки
//по API для дальнейшего прогноза потенциального времени ожидания ответа.

package support

type Support struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func (s *Support) Check() bool {
	return s.Topic != "" && s.ActiveTickets != 0
}
