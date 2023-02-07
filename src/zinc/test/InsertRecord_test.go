package test

import (
	"filehandlers"
	"testing"
	"zinc_handler"
)

func TestInserRecord(t *testing.T) {

	path := "./../../../static/enron_mail_20110402/maildir/allen-p/_sent_mail/1."

	t.Setenv("ZINC_ADMIN_USER", "admin")
	t.Setenv("ZINC_ADMIN_PASSWORD", "Complexpass#123")

	lines, errGetLines := filehandlers.GetLines(&path)
	if errGetLines != nil {
		t.Error(errGetLines)
		t.Fail()
	}

	email, errParsing := filehandlers.ParseStrings2Email(lines)
	if errParsing != nil {
		t.Error(errParsing)
	}

	request := zinc_handler.Model2IRequestData(email)
	response, err := zinc_handler.InsertSingleRecord(request)
	if err != nil {
		t.Errorf("Eror trying to insert. %v", err)
	}

	t.Log("Works!!!")
	t.Log(string(response))
}
