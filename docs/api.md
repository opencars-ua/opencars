# OpenCars API
API для отпримання публічної інформації про транспортні засоби України.

## Version: 0.1.5

### /vehicle/operations

#### GET
##### Summary:

Пошук оперіцій за реєстраційним номером

##### Description:

Отримати перелік оперіцій за реєстраційним номером

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| number | path | Реєстраційний номер | Yes | string |
| limit | path | Максимальна кількість операцій | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Список оперіції з номерним знаком | [ [Operation](#operation) ] |
| 400 | Помилковий запит | [Error](#error) |
| 405 | Помилковий метод | [Error](#error) |

### /vehicle/registrations

#### GET
##### Summary:

Пошук інформації за номером свідоцтва

##### Description:

Отримати інформацію про реєстрацію авто за номером свідоцтва

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| code | path | Номер свідоцтва про реєстрацію авто | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Список реєстрацій авто за номером свідоцтва | [ [Registration](#registration) ] |
| 400 | Помилковий запит | [Error](#error) |
| 404 | Не знайдено | [Error](#error) |
| 405 | Помилковий метод | [Error](#error) |
| 503 | Ресурс з даними про свідоцтва недоступний | [Error](#error) |

### Models


#### Error

Повідомлення про помилку

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string | Опис помилки | No |

#### Operation

Детальна інформація про операцію з транспортним засобом

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string | Адресса у форматі КОАТУУ | No |
| body | string | Тип кузова | No |
| brand | string | Марка | No |
| capacity | number (int32) | Об'єм двигуна | No |
| color | string | Колір транспортного засобу | No |
| date | string | Дата проведення операції | No |
| description | string | Тип операції | No |
| fuel | string | Тип палива | No |
| kind | string | Тип авто | No |
| model | string | Модель | No |
| number | string | Реєстраційний номер | No |
| office_id | number (int64) | Код місця проведення операції | No |
| office_name | string | Місце проведення операції | No |
| operation | string | Універсальний код операції | No |
| purpose | string | Тип авто | No |
| weight | number (int32) | Маса без навантаження | No |
| year | number (int32) | Рік випуску | No |
| vin | string | Номер шасі | No |

#### Registration

Детальна інформація про операцію з транспортним засобом

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| body | string | Тип кузова | No |
| brand | string | Марка | No |
| capacity | number (int32) | Об'єм двигуна | No |
| category | string | Категорія | No |
| code | string | Номер свідоцтва | No |
| color | string | Колір | No |
| date | string | Дата реєстрації | No |
| first_reg | string | Дата першої реєстрації | No |
| fuel | string | Тип палива | No |
| kind | string | Тип авто | No |
| model | string | Модель | No |
| number | string | Реєстраційний номер | No |
| own_weight | number (int32) | Маса без навантаження | No |
| total_weight | number (int32) | Повна маса | No |
| vin | string | Номер шасі | No |
| year | number (int32) | Рік випуску | No |