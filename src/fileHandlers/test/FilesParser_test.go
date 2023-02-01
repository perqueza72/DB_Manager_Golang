package testing

import (
	"filehandlers"
	"testing"
)

func TestFolderParser(t *testing.T) {

	emails, err := filehandlers.FolderFiles2Email("/home/perqueza72/Workspace/Truora/prueba_tecnica/statics/enron_mail_20110402/maildir/allen-p/_sent_mail")
	if err != nil {
		t.Error(err)
	}

	for _, email := range emails {
		t.Logf(email.Message)
	}

	t.Log("Works!!!")
}

func TestFileParser(t *testing.T) {

	path := "/home/perqueza72/Workspace/Truora/prueba_tecnica/statics/enron_mail_20110402/maildir/allen-p/_sent_mail/1."
	lines, errGetLines := filehandlers.GetLines(path)
	if errGetLines != nil {
		t.Error(errGetLines)
	}

	email, errParsing := filehandlers.ParseStrings2Email(lines)
	if errParsing != nil {
		t.Error(errParsing)
	}

	expect := " Mon, 14 May 2001 16:39:00 -0700 (PDT)"

	if email.Date != expect {
		t.Errorf("Date must be %v. It returns %v", expect, email.Date)
		t.Fail()
	}

	t.Log("Works!!!")
}

func TestParseStrings2Email(t *testing.T) {
	got, _ := filehandlers.GetLines("/home/perqueza72/Workspace/Truora/prueba_tecnica/statics/enron_mail_20110402/maildir/allen-p/all_documents/1.")

	email, err := filehandlers.ParseStrings2Email(got)
	if err != nil {
		t.Fail()
	}

	subjectExpected := " December 14, 2000 - Bear Stearns' predictions for telecom in Latin America"
	if email.Subject != subjectExpected {
		t.Errorf(`\nSubject expect:\n\t"%v"\nGotted :\n\t"%v"\n`, subjectExpected, email.Subject)
		t.Fail()
	}

}
