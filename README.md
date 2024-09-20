# AcctKeeper

AcctKeeper is a backend service for managing financial transactions and generating accounting reports.

## Features
- Record financial transactions for multiple user accounts
- Generate monthly and yearly financial reports

## Installation

#### Prerequisites
- Go 1.x
- MySQL or compatible database

## Setup

#### Download Repo
```bash
git clone https://github.com/zz040206655/acctkeeper.git
cd acctkeeper
go mod tidy
```
#### Modify config file
```yaml
server:
  port: ":5050"
database:
  user: "leonard"
  password: "leonard_password"
  host: "localhost"
  port: "3306"
  name: "leonard_db"
```
#### Start and create MySQL database using docker
```bash
./mysql-docker start
./mysql-docker build
./go run ./cmd/main.go
```

## Test

#### Register
```bash
curl -X POST http://localhost:5050/register -H "Content-Type: application/json" -d '{"username": "leo"}'
```
#### Add Transaction
```bash
curl -X POST http://localhost:5050/transaction -H "Content-Type: application/json" -d '{"username": "leo", "amount": 22.22, "type": "bank", "txtime": "2024-09-15T12:00:00Z"}'
```
#### Import Transactions
```bash
curl -X POST http://localhost:5050/import_transactions -H "Content-Type: application/json" -d '[
    {"username": "leo", "amount": 120.4, "type": "bank", "txtime": "2024-09-15T12:00:00Z"},
    {"username": "leo", "amount": -100.0, "type": "cash", "txtime": "2024-09-15T12:02:00Z"},
    {"username": "leo", "amount": 50.0, "type": "bank", "txtime": "2024-09-15T12:00:00Z"}
]'
```
#### Get Report
```bash
curl -X GET "http://localhost:5050/leo/report?year=2024&month=9"
```
