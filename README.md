# EQUIFAX TEST - DASHBOARD API

This is a Backend service in a BFF style developed in Golang that currently fetches user data and todos from https://dummyjson.com, aggregates them and exposes a single endpoint `GET /dashboard/:id`.

## Install
```bash
git clone https://github.com/MaximilianoTapiaMonsalve/Equifax-Test.git
```


## Setup 
```bash
go mod tidy
```

## Run
```bash
go run main.go
```

If you use AIR you can run it with
```bash
air
```

## Test

```bash
go test ./test/... -v
```

## Usage

```bash
curl http://localhost:3000/dashboard/1
```
