#Отдельно команды для запуска тестов:

Запуск всех тестов в пакете internal/calendar:
go test ./internal/calendar


Подробный вывод:
go test -v ./internal/calendar


Запуск всех тестов во всём проекте:
go test ./...


Запуск конкретного теста:
go test -v -run TestCreateEvent ./internal/calendar


Принудительный прогон без кэша:
go test -v -count=1 ./internal/calendar

============================

##CREATE — создать событие
curl -X POST http://localhost:8080/create_event ^
  -H "Content-Type: application/json" ^
  -d "{\"user_id\":1,\"date\":\"2025-01-01\",\"event\":\"Новый год\"}"

##UPDATE — обновить событие
curl -X POST http://localhost:8080/update_event ^
  -H "Content-Type: application/json" ^
  -d "{\"id\":1,\"event\":\"Новый год (обновлено)\"}"

##DELETE — удалить событие
curl -X POST http://localhost:8080/delete_event ^
  -H "Content-Type: application/json" ^
  -d "{\"id\":1}"

##GET события на день
curl "http://localhost:8080/events_for_day?user_id=1&date=2025-01-01"

##GET события на неделю
curl "http://localhost:8080/events_for_week?user_id=1&date=2025-01-01"

##GET события на месяц
curl "http://localhost:8080/events_for_month?user_id=1&date=2025-01-01"

##Некорректная дата (пример ошибки)
curl "http://localhost:8080/events_for_day?user_id=1&date=2025-99-99"

##Удаление несуществующего события (пример ошибки)
curl -X POST http://localhost:8080/delete_event ^
  -H "Content-Type: application/json" ^
  -d "{\"id\":999}"
