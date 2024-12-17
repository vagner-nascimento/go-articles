package monkeypatching_test

import (
	"articles/tests/mocks/monkeypatching"
	"database/sql"
	"testing"
)

func TestOpenDd(t *testing.T) {
	tests := []struct {
		name                 string
		description          string
		user, pass, addr, db string
		expectedError        error
		sqlOpener            func(string, string) (*sql.DB, error)
	}{
		{
			name:          "connection success",
			description:   "should return sql.DB and error nil",
			user:          "user",
			pass:          "pass",
			addr:          "addr",
			db:            "db",
			expectedError: nil,
			sqlOpener: func(s1, s2 string) (*sql.DB, error) {
				return &sql.DB{}, nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.description)

			monkeypatching.SqlOpen = test.sqlOpener

			db, err := monkeypatching.OpenDd(test.user, test.pass, test.addr, test.db)
			if err != test.expectedError {
				t.Errorf("expected %v \ngot %v", test.expectedError, err)
			} else if db == nil {
				t.Errorf("expected a DB, got nil")
			}
		})
	}
}
