# Status of Services
## Финальная работа по курсу Golang-разработчик
### Сервис собирает данные о работе систем SMS, MMS, VoiceCall, Email, Billing, Support, Incidents.
Для работы требуется запускать симулятор исходных данных:   
go run simulator/main.go

Потом стартуем сам сервис, командой:    
go run cmd/main.go

По умолчанию, исходные данные поступают:    
"SMS" > из симулятора в data/sms.data   
"VoiceCall" > из симулятора в data/voice.data  
"Email" > из симулятора в data/email.data  
"Billing" > из симулятора в data/billing.data   
"MMS" > http://127.0.0.1:8383/mms   
"Support"  > "http://127.0.0.1:8383/support     
"Incidents"  > http://127.0.0.1:8383/accendent

Вывод обработанных данных по адресу: http://localhost:8080/     
Html-страница для обзора статистики: status_page.html