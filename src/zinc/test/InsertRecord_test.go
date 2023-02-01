package test

import (
	"bytes"
	"encoding/json"
	"filehandlers"
	"testing"
	"zinc_handler"
)

func TestInserRecord(t *testing.T) {

	path := "./../../../../statics/enron_mail_20110402/maildir/allen-p/_sent_mail/1."

	t.Setenv("ZINC_ADMIN_USER", "admin")
	t.Setenv("ZINC_ADMIN_PASSWORD", "Complexpass#123")

	lines, errGetLines := filehandlers.GetLines(path)
	if errGetLines != nil {
		t.Error(errGetLines)
		t.Fail()
	}

	email, errParsing := filehandlers.ParseStrings2Email(lines)
	if errParsing != nil {
		t.Error(errParsing)
	}

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(email)
	response, err := zinc_handler.InsertSingleRecord("email", &buf)
	if err != nil {
		t.Errorf("Eror trying to insert. %v", err)
	}

	t.Log("Works!!!")
	t.Log(string(response))
}
