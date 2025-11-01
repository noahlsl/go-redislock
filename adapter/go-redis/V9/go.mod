module github.com/noahlsl/go-redislock/adapter/go-redis/V9

go 1.24.5

replace github.com/noahlsl/go-redislock => ../../..

require (
	github.com/noahlsl/go-redislock v1.3.0
	github.com/redis/go-redis/v9 v9.16.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
)
