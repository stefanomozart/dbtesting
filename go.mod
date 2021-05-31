module github.com/stefanomozart/dbtesting

go 1.15

require (
	github.com/gocraft/dbr v0.0.0-20190714181702-8114670a83bd
	github.com/gocraft/dbr/v2 v2.7.1
	github.com/jackc/pgx/v4 v4.11.0
	github.com/stefanomozart/dbrx v0.0.0-20210402130815-2c5b4eabb7da
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
)

replace github.com/stefanomozart/dbrx => /home/stefano/dev/gopath/src/github.com/stefanomozart/dbrx
