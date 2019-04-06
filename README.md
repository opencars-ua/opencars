[![CircleCI](https://circleci.com/gh/opencars-ua/opencars.svg?style=svg)](https://circleci.com/gh/opencars-ua/opencars)
[![Go Report Card](https://goreportcard.com/badge/github.com/opencars-ua/opencars)](https://goreportcard.com/report/github.com/opencars-ua/opencars)

# OpenCars

## Overview

RESTful API with information about transport just using Ukrainian License Plate.

## Server

Download PostgreSQL dump before starting the server. **NOTE:** This file is pretty big (2.9 *GB*)

```sh
$ gsutil cp gs://opencars/cars.sql
```

Import database to PostgreSQL

```sh
$ cat cars.sql | psql -U postgres
```

Run the server

```sh
$ go run cmd/opencars/main.go
```

## API

```sh
$ http localhost:8080/transport?number="BA2927BT"
```

```json
[
    {
        "body": "ХЕТЧБЕК-В",
        "capacity": 1598,
        "color": "СІРИЙ",
        "date": "2018-09-26",
        "fuel": "ДИЗЕЛЬНЕ ПАЛИВО",
        "id": 8665886,
        "kind": "ЛЕГКОВИЙ",
        "model": "AUDI A1",
        "number": "ВА2927ВТ",
        "own_weight": 1284,
        "registration": "172 - ПЕРВИННА РЕЄСТРАЦІЯ ЛЕГКОВИХ ТЗ, ЯКІ ВВЕЗЕНО З-ЗА КОРДОНУ",
        "registration_address": "3510600000",
        "registration_code": 172,
        "year": 2011
    }
]
```

For more information see [documentation](./docs).

## Parser

Parse CSV files from `data.gov.ua` into one SQL dump.

```sh
$ go run cmd/parser/main.go -path=<PATH> # Path to CSV file
```

## Data

All information was taken from official Ukrainian resource https://data.gov.ua

## License
Project released under the terms of the MIT [license](./LICENSE).
