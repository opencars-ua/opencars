[![CircleCI](https://circleci.com/gh/opencars/opencars.svg?style=svg)](https://circleci.com/gh/opencars/opencars)
[![Go Report Card](https://goreportcard.com/badge/github.com/opencars/opencars)](https://goreportcard.com/report/github.com/opencars/opencars)

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
        "address": "8036100000",
        "body": "СЕДАН-B",
        "brand": "MERCEDES-BENZ",
        "capacity": 2987,
        "color": "ЧОРНИЙ",
        "date": "2018-10-18",
        "description": "ПЕРВИННА РЕЄСТРАЦІЯ ЛЕГКОВИХ ТЗ, ЯКІ ВВЕЗЕНО З-ЗА КОРДОНУ",
        "fuel": "ДИЗЕЛЬНЕ ПАЛИВО",
        "kind": "ЛЕГКОВИЙ",
        "model": "S 350 D",
        "number": "АА1234ВР",
        "office_id": 8046,
        "office_name": "ТСЦ 8046",
        "operation": 172,
        "person": "P",
        "purpose": "ЗАГАЛЬНИЙ",
        "weight": 1825,
        "year": 2017
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
