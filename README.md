# transfer-system-golang

A simple money transfer system written in Go, using PostgreSQL for storage.

## Features
- Account management
    - Create new account
    - Look up account by ID
- Transaction management
    - Create new transaction

## Tech Stack
- Programming language: Go
- Database: PostgreSQL

## Prerequisites
- Go 1.24+
- PostgreSQL

## Environment Variables
Set the following in your `.env` file:
- `APP_PORT` (default: 8000)
- `POSTGRES_HOST`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DBNAME`

## How to Run
1. Prepare a PostgreSQL database.
2. Update `.env` with your database credentials.
3. Set up environment variables: `source .env`.
4. Run the app: `go run cmd/api-server/main.go`.
5. The app will automatically run the migrations.

## API Examples

**Create Account**
```sh
curl -X POST http://localhost:8000/accounts -d '{"account_id":1,"initial_balance":"100.00"}' -H "Content-Type: application/json"
```

**Get Account**
```sh
curl http://localhost:8000/accounts/1
```

**Create Transaction**
```sh
curl -X POST http://localhost:8000/transactions -d '{"source_account_id":1,"destination_account_id":2,"amount":"10.00"}' -H "Content-Type: application/json"
```

## Running Tests

To run unit tests:
```sh
go test ./...
```

## Project Structure

- `cmd/`  
  Application entrypoint(s). Contains the main executable.

- `internal/domains/`  
  Business logic and core domain models. Contains service logic and repository interfaces for accounts, transactions, and money.

- `internal/infrastructure/`  
  Infrastructure code such as HTTP handlers, database implementations, configuration, and migrations.

- `internal/infrastructure/httpserver/`  
  HTTP server setup and request handlers.

- `internal/infrastructure/db/`  
  Database connection, repository implementations, and migration files.

- `internal/infrastructure/config/`  
  Application configuration loading.

This structure separates domain logic from infrastructure code, following Clean Architecture principles.

## Bonus
### Setting Up TigerBeetle (Optional)
TigerBeetle is an ultra-fast, distributed financial ledger database. To experiment with TigerBeetle, follow these steps:

1. Download TigerBeetle Get the latest release from [TigerBeetle Releases](https://github.com/tigerbeetle/tigerbeetle/releases).
2. Format the Database `tigerbeetle format --cluster=0 --replica=0 --replica-count=1 --development ./0_0.tigerbeetle`
3. Start the TigerBeetle Server `tigerbeetle start --addresses=3000 --development ./0_0.tigerbeetle`
4. Open the TigerBeetle REPL `tigerbeetle repl --cluster=0 --addresses=3000`
5. Create an Tresury Account In the REPL, run: `create_accounts id=1 ledger=1 code=1`
6. Update `.env` 
   ```
   export FEATURE_FLAG_TIGERBEETLE="ON"
   export TIGERBEETLE_ADDRESS="3000"
   ```

For more details, see the [TigerBeetle documentation](https://docs.tigerbeetle.com/).


## License
This project is released into the public domain under [The Unlicense](LICENSE).