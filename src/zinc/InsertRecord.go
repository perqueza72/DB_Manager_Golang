package zinc_handler

import (
	"bytes"
	constants "constants_project"
	"encoding/json"
	"io"
	"log"
	models "models_zinc"
	"net/http"
)

func Model2IRequestData(model interface{}) *models.IRequestData {
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(model)
	request := models.IRequestData(&buf)
	return &request
}

func InsertSingleRecord(data *models.IRequestData) ([]byte, error) {
	req, err := http.NewRequest("POST", constants.ZINC_HOST+"/api/"+constants.ZINC_EMAIL_INDEX+"/_doc", *data)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}

	return body, nil
}
