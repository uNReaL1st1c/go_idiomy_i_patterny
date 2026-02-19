# Медицинский центр — поэтапный учебный проект

## Цель проекта
Комплексная система управления медицинским центром с использованием возможностей Go, стандартной библиотеки и популярных технологий: PostgreSQL, Redis, Kafka, HTTP API, gRPC и др.

---

## Этап 1: Фундамент (Go Core + файловая система)
**Сложность:** начальная  
**Что изучаем:** базовый синтаксис, структуры, интерфейсы, JSON, файлы, flag/viper

### Задача
Приложение для учета пациентов. Данные хранятся в JSON-файле.

### Функциональность
- Структуры `Patient`, `Appointment`
- CRUD пациентов (добавить, получить, обновить, удалить)
- Сохранение/загрузка из `patients.json`
- CLI: флаги для операций (`--add`, `--list`, `--find-by-id`)
- Валидация полей (ФИО, дата рождения, пол)

### Технологии
- `encoding/json`, `os`, `io`, `flag`
- unit-тесты для CRUD-логики

---

## Этап 2: HTTP API + REST
**Сложность:** базовая  
**Что изучаем:** net/http, роутинг, middleware, контекст

### Задача
Обернуть логику этапа 1 в REST API.

### Функциональность
- `GET /api/v1/patients` — список пациентов
- `GET /api/v1/patients/:id` — пациент по ID
- `POST /api/v1/patients` — создание
- `PUT /api/v1/patients/:id` — обновление
- `DELETE /api/v1/patients/:id` — удаление
- Middleware: логирование, recovery, CORS
- Конфигурация через env/файл (порт, путь к данным)

### Технологии
- `net/http`, `context`, chi/gin или stdlib
- `godotenv` / env vars

---

## Этап 3: PostgreSQL + репозиторий
**Сложность:** средняя  
**Что изучаем:** sqlx/pgx, миграции, транзакции, connection pool

### Задача
Перенести хранение из JSON в PostgreSQL.

### Функциональность
- Схема: `patients`, `doctors`, `appointments`
- Репозиторий (интерфейс + PostgreSQL-реализация)
- Миграции (golang-migrate или goose)
- Транзакции при создании записи на приём
- Connection pool, таймауты

### Технологии
- `pgx` или `database/sql` + `sqlx`
- `golang-migrate/migrate`

---

## Этап 4: Redis — кэш и сессии
**Сложность:** средняя  
**Что изучаем:** Redis клиент, TTL, кэширование, сессии

### Задача
Добавить кэш для частых запросов и простую аутентификацию по сессии.

### Функциональность
- Кэш: `GET /patients/:id` — сначала Redis, при промахе — PostgreSQL
- Инвалидация кэша при создании/обновлении/удалении пациента
- Сессии (token в Redis, TTL 24ч)
- `POST /auth/login`, `POST /auth/logout`

### Технологии
- `go-redis/redis/v9`
- Интеграционные тесты с testcontainers (опционально)

---

## Этап 5: Kafka — события
**Сложность:** средняя–высокая  
**Что изучаем:** producer/consumer, топики, обработка событий

### Задача
Асинхронная обработка событий: новая запись, отмена, напоминания.

### Функциональность
- Producer: при создании appointment — событие в топик `appointments.created`
- Consumer: при получении события — логирование, обновление статистики, (опционально) отправка в «сервис напоминаний»)
- События: `appointment.created`, `appointment.cancelled`
- Dead letter при ошибках

### Технологии
- `confluent-kafka-go` или `segmentio/kafka-go`

---

## Этап 6: gRPC — внутренняя коммуникация
**Сложность:** высокая  
**Что изучаем:** protobuf, gRPC сервер/клиент, потоковые вызовы

### Задача
Внутренний сервис расписания врачей: HTTP API вызывает gRPC для получения слотов.

### Функциональность
- gRPC-сервис `ScheduleService`: `GetAvailableSlots`, `BookSlot`
- HTTP API вызывает gRPC-клиент для проверки слотов
- Stream: `StreamDailyAppointments(doctorID, date)` — стриминг записей по дате

### Технологии
- `google.golang.org/grpc`
- `protoc` + `protoc-gen-go`, `protoc-gen-go-grpc`

---

## Этап 7: Observability + Production-ready
**Сложность:** высокая  
**Что изучаем:** логирование, метрики, трейсинг, graceful shutdown

### Задача
Сделать приложение готовым к продакшену.

### Функциональность
- Структурированные логи (zerolog/zap)
- Prometheus-метрики (latency, количество запросов, ошибок)
- OpenTelemetry (опционально) или логирование trace-id
- Graceful shutdown: дожидаемся завершения запросов, закрываем БД/Redis/Kafka
- Health checks: `/health`, `/ready`

### Технологии
- `zerolog` / `zap`
- `prometheus/client_golang`
- `context` для отмены при shutdown

---

## Этап 8: Микросервисная декомпозиция (опционально)
**Сложность:** продвинутая  
**Что изучаем:** разбиение на сервисы, событийная шина

### Задача
Разделить монолит на 2–3 сервиса.

### Функциональность
- **Patients Service:** пациенты, HTTP API
- **Appointments Service:** записи, gRPC + Kafka consumer
- **Notifications Service:** напоминания (Kafka consumer)
- Взаимодействие через Kafka и gRPC

---

## Рекомендуемая структура проекта (после этапа 3+)

```
medcenter/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── domain/
│   │   ├── patient.go
│   │   └── appointment.go
│   ├── repository/
│   │   ├── patient.go          # interface
│   │   └── patient_postgres.go
│   ├── service/
│   │   └── patient.go
│   ├── transport/
│   │   ├── http/
│   │   └── grpc/
│   └── kafka/
├── migrations/
├── pkg/
│   └── ...
├── go.mod
├── go.sum
└── PROJECT_PLAN.md
```

---

## Порядок работы

1. Начинаем с **Этапа 1**.
2. После реализации этапа — **code review** (плюсы и минусы).
3. Переходим к следующему этапу.

Можно приступать к Этапу 1.
