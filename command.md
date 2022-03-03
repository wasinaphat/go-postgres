#create_migration
migrate create -ext sql -dir db/migration -seq init_schema

install 
go testify pkg