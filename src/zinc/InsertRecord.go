package zinc_handler

import (
	constants "constants_project"
	"io"
	"log"
	models "models_zinc"
	"net/http"
)

func InsertSingleRecord(index_name string, data models.IRequestData) ([]byte, error) {
	req, err := http.NewRequest("POST", constants.ZINC_HOST+"/api/"+index_name+"/_doc", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return body, nil
}
