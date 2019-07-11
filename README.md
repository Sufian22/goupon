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

To execute the tests just run the following command:
```bash
go test ./... 
```

__NOTE: The tests will be executed against the postgres database specified on the configuration file__


## API Specification

### Create Coupon

Endpoint:

```bash
POST /api/coupons
```

Request body:

```json
{
    "name": "testname",
    "brand": "testbrand",
    "value": 12.2
}
```

### Get Coupon

Endpoint:

```bash
GET /api/coupons/{id}
```

```bash
- id => Coupon identifier, only integers allowed
```

### List Coupons

Endpoint:

```bash
GET /api/coupons
```

Query parameters:

```bash
- q => Filtering value, it will be compared against coupon "name" and coupon "brand", should satisfy "^[a-z0-9]*$" pattern
- orderBy => Order field, only "name" and "brand" are allowed. Default "name"
- order => Sort order, only "ASC" and "DESC" are allowed. Default "ASC"
- quantity => Number of coupons that want to be listed, only "integers" allowed. Default "10"
```

### Update Coupon

```bash
PUT /api/coupons/{id}
```

Request body:

```json
{
    "name": "newname",
    "brand": "newbrand",
    "value": 15
}
```
