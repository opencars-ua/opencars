# OpenCars

[![CircleCI](https://circleci.com/gh/opencars/opencars.svg?style=svg)](https://circleci.com/gh/opencars/opencars)
[![Go Report Card](https://goreportcard.com/badge/github.com/opencars/opencars)](https://goreportcard.com/report/github.com/opencars/opencars)

## Overview

RESTful API with information about transport just using Ukrainian License Plate.

## Server

Download PostgreSQL dump before starting the server. **NOTE:** This file is pretty big (3.1 _GB_)

```sh
$ gsutil cp gs://opencars/cars.sql cars.sql
```

Import database to PostgreSQL

```sh
$ cat cars.sql | psql -U postgres

```

Run the server

```sh
$ go run cmd/server/main.go
```

## API

```sh
$ http localhost:8080/transport?number="АА9359РС"
```

```json
[
  {
    "address": "8036600000",
    "body": "УНІВЕРСАЛ-B",
    "brand": "TESLA",
    "capacity": 0,
    "color": "ЧОРНИЙ",
    "date": "05.06.2019",
    "description": "ПЕРЕРЕЄСТРАЦIЯ У ЗВ`ЯЗКУ ЗI ЗМIНОЮ АНКЕТНИХ ДАНИХ ВЛАСНИКА",
    "fuel": "ЕЛЕКТРО",
    "kind": "ЛЕГКОВИЙ",
    "model": "MODEL X",
    "number": "АА9359РС",
    "office_id": 12290,
    "office_name": "ТСЦ 8041",
    "operation": 340,
    "person": "P",
    "purpose": "ЗАГАЛЬНИЙ",
    "weight": 2485,
    "year": 2016,
    "vin": "5YJXCCE40GF010543"
  }
]
```

For more information see [documentation](./docs/api.md).

## Parser

Parse CSV files from `data.gov.ua` into one SQL dump.

```sh
$ go run cmd/parser/main.go -path=<PATH> # Path to CSV file
```

## Data

All information was taken from [official Ukrainian resource](https://data.gov.ua)

## License

Project released under the terms of the MIT [license](./LICENSE).
