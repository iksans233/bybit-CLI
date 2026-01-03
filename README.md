# **Simple BYBIT CLI (demo)**

### Simple bybit CLI tool for trade activity, show equity, market ticker using native http package(spot only)

### requirements:

Bybit ApiKey & SecretKey

https://github.com/joho/godotenv

https://github.com/spf13/cobra

### Domain

Rest API: https://api-demo.bybit.com

### Usage:
#### Get wallet balance:

go run main.go account [symbol]

Example: go run main.go account USDT

#### Get Ticker Price:
go run main.go market [symbol]

**Note: Always use uppercase for symbol**

#### Place order:
Example: go run main.go buy --symbol ENAUSDT --qty 100 --price 0.1 --type limit --side buy --category spot

#### Cancelall:
go run main.go cancelall --category spot

#### List pending order:
go run main.go showorder


###### Reference: https://bybit-exchange.github.io/docs/v5