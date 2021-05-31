# dbtesting
Utility package for running tests in golang with [dbrx](https://github.com/stefanomozart/dbrx) and [dbr](https://github.com/gocraft/dbr).

## Instalation

Use the `go get`command:

```{go}
go get github.com/stefanomozart/dbtesting
```

## Usage

Use the `dbtesting.Setup()` to run the sql scripts needed to setup the test environment, then, run your tests, as in the example bellow:

```{go}
import (
    "reflect"
	"testing"

    "github.com/stefanomozart/dbtesting"
)

func Test_SearchUser(t *testing.T) {
	tests := []struct {
		name    string
		script  string
		filter  *User
		want    []User
		wantErr bool
	}{
		{
			"no records, must return an error",
			"",
			&User{ID: 1},
			nil,
			true,
		},
		{
			"many records, filters one by ID",
			`INSERT INTO
				auth.user(id, username, name, password, status)
			VALUES
				(1, 'user1', 'First User', '5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8', 1),
				(2, 'user2', 'Second User', '5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8', 1),
				(3, 'user3', 'Third User', '5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8', 2);`,
			&User{ID: 1},
			[]User{
                {ID: 1, Username: "user1", Name: "Fist User", Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",  Status: 1},
            }
			false,
		},
		{
			"many recors, filters one by Username",
			`INSERT INTO
				auth.user(id, username, name, password, status)
			VALUES
				(1, 'user1', 'First User', 'test', 1),
				(2, 'user2', 'Second User', 'test', 1),
				(3, 'user3', 'Third User', 'test', 2);`,
			&User{Username: "user2"},
			[]User{
                {ID: 1, Username: "user1", Name: "Fist User", Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",  Status: 1},
            },
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            sess = dtesting.Setup("dbtesting/schema.sql", tt.script)
			
            // assuming your SerachUser function receives the db session as its first parameter
            // and the filter conditions as the second
			got, err := SearchUser(sess, tt.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUser()\ngot  %v\nwant %v", got, tt.want)
			}
		})
	}
}
```

Note that if no database connection is active, the `Setup` function wiil call the `SetupConn(dsn string)` with and empty dsn string, and the SetupConn(dsn string) will use the following environment variables to connect to a postgresql server:

```
dsn = fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    GetEnv("DBTESTING_HOST", "localhost"),
    GetEnv("DBTESTING_PORT", "5432"),
    GetEnv("DBTESTING_USER", "testuser"),
    GetEnv("DBTESTING_PASSWD", "testpassword"),
    GetEnv("DBTESTING_DBNAME", "testdb"),
)
```

## Supported database divres
This package was written especifically to be used with the `postgres` or the
`pgx` database drivers. It was also used with the `SQLite3` driver for 
testing.
