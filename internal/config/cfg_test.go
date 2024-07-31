package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		mockFunc func()
		err      error
		want     *Config
	}{
		{
			name: "success",
			mockFunc: func() {

				t.Setenv("POSTGRES_HOST", "localhost")
				t.Setenv("POSTGRES_USER", "testuser")
				t.Setenv("POSTGRES_PASSWORD", "testpassword")
				t.Setenv("POSTGRES_PORT", "5432")
				t.Setenv("POSTGRES_DB", "testdb")
				t.Setenv("POSTGRES_SLLMODE", "disable")
				t.Setenv("POSTGRES_TIMEZONE", "UTC")
			},
			want: &Config{
				Database: Database{
					Host:     "localhost",
					User:     "testuser",
					Password: "testpassword",
					Port:     "5432",
					DbName:   "testdb",
					SSLMode:  "disable",
					TimeZone: "UTC",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			cfg, err := Load()

			assert.Equal(t, err, tc.err)
			assert.EqualExportedValues(t, cfg, tc.want)
		})
	}
}
