package service

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCheckMailSentInEventBuffer(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"o_id", "o_unique_id", "o_sender", "o_receiver", "o_template", "o_created"}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT (.+) FROM event_buffer").WithArgs(1).WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1,1"))
	mock.ExpectCommit()
}

func TestUpdateSendDate(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE event_buffer").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if _, err = updateSendDate(2); err != nil {
		t.Errorf("error was not expected while updating send date: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestValidateEmail(t *testing.T) {
	got := validateEmail("test@test.com")
	want := true

	if got != want {
		t.Error("Validation if argument is mail address")
	}
}

func TestValidateNotEmpty(t *testing.T) {
	got := validateNotEmpty("teststring")
	want := true

	if got != want {
		t.Error("Valdiation if argument is not empty")
	}
}
