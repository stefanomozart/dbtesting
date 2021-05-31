package dbrx

import (
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stefanomozart/dbrx"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name   string
		schema string
		script string
		want   int
	}{
		{
			"criar tabela e inserir uma linha",
			"test_schema.sql",
			`INSERT INTO
				teste(id, texto)
			VALUES
				(1, 'isso é um teste');`,
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dml := Setup(tt.schema, tt.script)

			var got int
			dml.Select("count(*)").From("teste").LoadOne(&got)

			if got != tt.want {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetupConn(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		wantErr bool
	}{
		{
			"without dns",
			"",
			false,
		},
		{
			"with dns",
			fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbrx.GetEnv("DBTESTING_HOST", "localhost"),
				dbrx.GetEnv("DBTESTING_PORT", "5432"),
				dbrx.GetEnv("DBTESTING_USER", "postgres"),
				dbrx.GetEnv("DBTESTING_PASSWD", ""),
				dbrx.GetEnv("DBTESTING_DBNAME", "postgres"),
			),
			false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := dbrx.SetupConn(tt.dsn)
			if got == nil {
				t.Errorf("SetupConn() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
