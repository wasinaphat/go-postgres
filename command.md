#create_migration
migrate create -ext sql -dir db/migration -seq init_schema

install 
go testify pkg


for testing
mockgen -package mockdb  --destination db/mock/store.go  github.com/wasinaphatlilawatthananan/go-postgres/db/sqlc Store