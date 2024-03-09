# Домашнее задание 5

[Вернуться на Главную страницу](../../../README.MD)

Дисциплина: Программирование на Golang

Тема 7.Создание RESTful API на Golang

## Описание задания

Реализуйте веб-сервис обработки транзакций. Транзакция должна включать в себя информацию о сумме, валюте, типе (доход/расход), категории, дате и описании транзакции.

### Шаг 1. Определите структуру транзакции:
- ID
- сумма
- валюта: USD, EUR и т. д.
- тип: доход, расход или перевод
- категория: зарплата, еда, жильё и т. д.
- дата
- описание

### Шаг 2. Запустите веб-сервер с помощью метода http.ListenAndServe() и реализуйте обработчик входящих запросов.


### Шаг 3. Разработайте следующие эндпоинты:
создание транзакции: POST /transactions;
получение списка всех транзакций: GET /transactions;
получение одной транзакции по ID: GET /transactions/{id};
обновление транзакции по ID: PUT /transactions/{id};
удаление транзакции по ID: DELETE /transactions/{id}.

### Шаг 4. Проведите тестирование с использованием Postman или cURL.
Примечание
Разработайте структуру вашего API так, чтобы его можно было легко расширить в будущем. Используйте принципы RESTful для создания чистого и понятного API.

## Пример работы программы
---
### Создание транзакции POST /transactions:

![alt text](./assets/GET.png)

```bash
Rq: curl --header "Content-Type: application/json" --request POST --data '{"amount":3120.30,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание 1"}' http://localhost:8080/transactions
Rs: {"id":"1","status":"Success"}

curl --header "Content-Type: application/json" --request POST --data '{"amount":320.30,"currency":"EUR","type":"income","category":"Зарплата","date":"2024-03-02","description":"Описание 2"}' http://localhost:8080/transactions
Rs: {"id":"2","status":"Success"}

curl --header "Content-Type: application/json" --request POST --data '{"amount":30.30,"currency":"JPY","type":"income","category":"Зарплата","date":"2024-03-03","description":"Описание 3"}' http://localhost:8080/transactions
Rs: {"id":"3","status":"Success"}

```

### Получение списка всех транзакций: GET /transactions:

![alt text](./assets/POST.png)

```bash 
GET http://127.0.0.1:8080/transactions/
GET http://127.0.0.1:8080/transactions
```

```json
[
    {"id":"1","amount":3120.3,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание 1"},
    {"id":"2","amount":320.3,"currency":"EUR","type":"income","category":"Зарплата","date":"2024-03-02","description":"Описание 2"},
    {"id":"3","amount":30.3,"currency":"JPY","type":"income","category":"Зарплата","date":"2024-03-03","description":"Описание 3"}
]

```

### Получение одной транзакции по ID: GET /transactions/{id}

```bash 
GET http://127.0.0.1:8080/transactions/2
```

```json
{"id":"2","amount":320.3,"currency":"EUR","type":"income","category":"Зарплата","date":"2024-03-02","description":"Описание 2"}
```

### Обновление транзакции по ID: PUT /transactions/{id}
```bash 
curl --header "Content-Type: application/json" \
    --request PUT \
    --data '{"amount":220.80,"currency":"GBP","type":"income","category":"Зарплата Update","date":"2024-03-04","description":"Описание 2 Update"}'\
    http://localhost:8080/transactions/2

```

```json 
{"id":"2","status":"Success"}
```

**Проверка:**

```bash 
GET http://127.0.0.1:8080/transactions
```

```json
[
    {"id":"1","amount":3120.3,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание 1"},
    {"id":"3","amount":30.3,"currency":"JPY","type":"income","category":"Зарплата","date":"2024-03-03","description":"Описание 3"},
    {"id":"2","amount":220.8,"currency":"GBP","type":"income","category":"Зарплата Update","date":"2024-03-04","description":"Описание 2 Update"}]
```

### Удаление транзакции по ID: DELETE /transactions/{id}

```bash 
curl --header "Content-Type: application/json" --request DELETE --data '' http://localhost:8080/transactions/3

```

```json
{"id":"3","status":"Success"}
```

**Проверка:**

```bash 
GET http://127.0.0.1:8080/transactions
```

```json
[
    {"id":"1","amount":3120.3,"currency":"USD","type":"income","category":"Зарплата","date":"2024-03-01","description":"Описание 1"},
    {"id":"2","amount":220.8,"currency":"GBP","type":"income","category":"Зарплата Update","date":"2024-03-04","description":"Описание 2 Update"}
]
```

## Чек-лист самопроверки:

![alt text](./assets/RESTful.png)