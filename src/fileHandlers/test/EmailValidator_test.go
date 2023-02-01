package testing

import (
	"filehandlers"
	models "models_zinc"
	"testing"
)

func TestEmailValidator(t *testing.T) {
	email := models.Email{}

	exist := filehandlers.ExistProperty(email, "Message")
	if !exist {
		t.Errorf("Not work properly")
		t.Fail()
		return
	}

	t.Log("Work well.")
}

func TestEmailValidatorPropertyNotExist(t *testing.T) {
	email := models.Email{}

	exist := filehandlers.ExistProperty(email, "NotExist")
	if exist {
		t.Errorf("Not work properly")
		t.Fail()
		return
	}

	t.Log("Work well :)")
}
