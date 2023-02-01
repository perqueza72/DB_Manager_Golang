package test

import (
	"testing"
	"zinc_handler"
)

func TestCreateIndex(t *testing.T) {
	t.Setenv("ZINC_ADMIN_USER", "admin")
	t.Setenv("ZINC_ADMIN_PASSWORD", "Complexpass#123")

	index_handler := zinc_handler.NewIndexHandler()
	index_handler.IndexModel.Name = "another_name"
	got, err := index_handler.CreateIndex()

	if err != nil {
		t.Errorf("Error trying to get index")
		t.Failed()
	}
	t.Logf("This return %v", string(got))

}
