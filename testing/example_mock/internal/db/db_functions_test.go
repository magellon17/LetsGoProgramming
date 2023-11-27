package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names: []string{"Boris, Boris228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := Service{DB: mockDB}

	for i, row := range testTable {
		mock.
			ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDbRows(row.names)).
			WillReturnError(row.errExpected)

		names, err := dbService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	type args struct {
		columnName string
		tableName  string
	}

	testTable := []struct {
		args          args
		names         []string
		namesExpected []string
		errExpected   error
	}{
		{
			args:          args{"names", "users"},
			namesExpected: []string{"Boris", "Roma", "Dima"},
		},
		{
			args:        args{"names", "users"},
			names:       nil,
			errExpected: errors.New("ExpectedError"),
		},
	}

	dbService := Service{DB: mockDB}

	for i, row := range testTable {
		mock.
			ExpectQuery("SELECT DISTINCT " + row.args.columnName + " FROM " + row.args.tableName).
			WillReturnRows(mockDbRows(row.namesExpected)).
			WillReturnError(row.errExpected)

		names, err := dbService.SelectUniqueValues(row.args.columnName, row.args.tableName)
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.namesExpected, names, "row: %d, expected names: %s, actual names: %s", i, row.namesExpected, names)
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}
