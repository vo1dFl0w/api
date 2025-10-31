# Тестовое задание golang junior+

> REST API на Go с базовыми операциями с wallet с PostgreSQL и запуском в Docker.

**Коротко:**

* Принимает POST-запросы на изменение баланса кошелька.
* Позволяет получить баланс кошелька по UUID.

---

## Стек

* Язык: Go
* База данных: PostgreSQL
* Контейнеризация: Docker + docker-compose

---

## Структура репозитория (важные файлы/папки)

```
/api                     - корень проекта
├─ cmd/api/              - main.go
├─ db/                   - *.sql для создания таблиц в бд
├─ internal/app/         - исходники сервера (/apiserver, /model, /store)
├─ loadtests/            - vegeta targets, payload и скрипт loadtest.sh
├─ config.env            
├─ docker-compose.yml
├─ dockerfile
└─ README.md
```

---

## API

### 1. POST `/api/v1/wallet`

Запрос для изменения баланса (DEPOSIT или WITHDRAW).

**Пример тела (JSON):**

```json
{
  "uuid": "<UUID>",
  "operation": "DEPOSIT",
  "amount": 1000
}
```

**Успешный ответ:** `200 OK` с обновлённым состоянием кошелька.

**Ошибки:**

* `400` — некорректный входной JSON / валидация
* `5xx` — неожиданные ошибки сервера/БД (следует контролировать и фиксировать)

### 2. GET `/api/v1/wallets/{WALLET_UUID}`

Получить баланс кошелька по UUID.

**Пример ответа:**

```json
{
  "uuid": "<UUID>",
  "account": 12345,67
}
```

---

## Быстрый старт (локально)

### 1. Клонировать репозиторий

```bash
git clone https://github.com/vo1dFl0w/api.git
cd api
```

### 2. Создать файл переменных окружения

Скопируйте `config.env.example` или создайте `config.env` с нужными значениями. Пример основных переменных:

```
HTTPAddr=":8080"
DATABASE_URL=postgres://user:password@db:5432/wallets_db?sslmode=disable
```

> В репозитории используется `github.com/joho/godotenv` для загрузки `config.env`.

### 3. Поднять сервисы через docker-compose

```bash
docker-compose up --build
```

После старта:

* API будет доступно на `http://localhost:8080`
* БД поднимается как отдельный контейнер

---

## Тесты

### Unit / Integration

Запустить все тесты на машине:

```bash
# локально
go test ./... -v
```

В проекте присутствуют `*_test.go` — unit и интеграционные тесты для апи/репозитория.

### Нагрузочное тестирование (vegeta)

В репозитории есть `loadtests/` с готовым `loadtest.sh`, `targets.txt` и `payload.json`.

Пример ручного запуска:

```bash
cd loadtests
# по умолчанию rate=1000 duration=30s
./loadtest.sh
```

`loadtest.sh` генерирует отчёты в `loadtests/results/` (results.bin, report.txt, report.json).

---
