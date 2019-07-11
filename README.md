# Goupon

## Requirements

- Install go
- Install dependencies
```bash
go get ./...
```

## Usage 

Change the values on the configuration file (`config/config.json`) for your desired ones:

```json
{
    "port": "3000",
    "dbConfig": {
        "dbDriver": "postgres",
        "dbHost": "localhost",
        "dbPort": "5432",
        "dbUser": "postgres",
        "dbPassword": "demodemo",
        "dbName": "postgres"
    }
}
```

And execute:

```bash
go run cmd/goupon/main.go
```

You can change the configuration file path in the `config` argument:

```bash
Usage of cmd/goupon/main.go
    -config string
        Configuration file (default "./config/config.json")
```

## API Specification