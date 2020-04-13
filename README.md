# Project Codename Transfer Banking Simulate

Simluation to transfer banking

## Usage

Basic instructions to get started with the project

### Backend

Docker start:
```
$ docker run --name postgres -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres
```

Migration Database
```
$ soda migrate up
```

Seed database:
```
$ go run config/seed.go
```

Run backend:
```
$ go run main.go
```
