package testing

import (
	"filehandlers"
	models "models_zinc"
	"testing"
)

func TestGetLines(t *testing.T) {
	path := "./../../../static/enron_mail_20110402/maildir/allen-p/all_documents/1."
	_, err := filehandlers.GetLines(&path)

	if err != nil {
		t.Errorf("Error getting lines. %v", err)
		t.Fail()
	}

	t.Logf("getted")

}

func TestSetField(t *testing.T) {
	email := models.Email{}
	fieldName := "Message-ID"
	value := "5"

	filehandlers.SetField(&email, fieldName, value)

	if email.Message != value {
		t.Errorf("%v is different from %v", email.Message, value)
		t.Fail()
	}

	t.Log("Ok")
}

func TestSetFieldDate(t *testing.T) {
	email := models.Email{}
	fieldName := "Date"
	value := "5"

	filehandlers.SetField(&email, fieldName, value)

	if email.Date != value {
		t.Errorf("%v is different from %v", email.Date, value)
		t.Fail()
	}

	t.Log("Ok")
}
