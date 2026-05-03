module appfactory/upgrade-service

go 1.23.0

require (
	appfactory/shared-go v0.0.0
	github.com/jackc/pgx/v5 v5.7.6
)

replace appfactory/shared-go => ../shared/go
