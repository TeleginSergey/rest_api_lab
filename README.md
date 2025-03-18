# REST API для работы с альбомами

Этот проект представляет собой простое REST API для управления данными об альбомах. API поддерживает базовые CRUD-операции (Create, Read, Update, Delete) для сущности "альбомы". Проект использует PostgreSQL в качестве базы данных и написан на Go.

## Содержание
1. [Требования](#требования)
2. [Запуск проекта](#запуск-проекта)
3. [Использование API](#использование-api)
   1. [Контакты](#контакты)
      1. [Получить все контакты](#получить-все-контакты)
      2. [Создать новый контакт](#создать-новый-контакт)
      3. [Обновить контакт](#обновить-контакт)
      4. [Удалить контакт](#удалить-контакт)
   2. [Группы](#группы)
      1. [Получить все группы](#получить-все-группы)
      2. [Создать новую группу](#создать-новую-группу)
      3. [Обновить группу](#обновить-группу)
      4. [Удалить группу](#удалить-группу)
4. [Проблемы и решения](#проблемы-и-решения)
5. [Лицензия. MIT LICENSE](#лицензия-mit-license)

---

## Требования

- Go (версия 1.20 или выше)
- Docker
---

## Запуск проекта

Запустите docker-compose.yml:
```
docker compose up --build
```

## Использование API

API доступен на порту `8080`. Ниже приведены примеры запросов с использованием `curl`.

---

### Контакты

#### Получить все контакты

```bash
curl -X GET http://localhost:8080/contacts
```

#### Создать новый контакт
```bash
curl -X POST http://localhost:8080/contacts \
-H "Content-Type: application/json" \
-d '{
    "id": 3,
    "username": "alice",
    "givenname": "Alice",
    "familyname": "Smith",
    "phone": [
        {
            "typeid": 3,
            "countrycode": 1,
            "operator": 789,
            "number": 5559876
        }
    ],
    "email": ["alice.smith@example.com"],
    "birthday": "2000-10-10T00:00:00Z"
}'
```

#### Обновить контакт
```bash
curl -X PUT http://localhost:8080/contacts/1/username/johndoe_new
```

#### Удалить контакт
```bash
curl -X DELETE http://localhost:8080/contacts/1
```

### Группы

#### Получить все группы
```bash
curl -X GET http://localhost:8080/groups
```

#### Создать новую группу
```bash
curl -X POST http://localhost:8080/groups \
-H "Content-Type: application/json" \
-d '{
    "id": 2,
    "title": "Work",
    "description": "Colleagues from work",
    "contacts": [3]
}'
```

#### Обновить группу
```bash
curl -X PUT http://localhost:8080/groups/1/title/BestFriends
```

#### Удалить группу
```bash
curl -X DELETE http://localhost:8080/groups/1
```

## Проблемы и решения

Если порт 9080 уже занят, выполните следующие шаги:
1. Найдите процесс, занимающий порт:
```bash
lsof -i :8080
```
2. Завершите процесс:
```bash
kill -9 <PID>
```

## Лицензия. MIT LICENSE
Этот проект распространяется под лицензией MIT. Подробности см. в файле [LICENSE](LICENSE).
