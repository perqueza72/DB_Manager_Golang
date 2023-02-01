package zinc_handler

import (
	"bytes"
	constants "constants_project"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	models "models_zinc"
	"net/http"
	"os"
)

type IndexHandler struct {
	IndexModel *models.ZincIndex
}

func NewIndexHandler() *IndexHandler {
	jsonFile, err := os.Open("./../../standard_index_structure.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened index structure")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var standardIndex models.ZincIndex
	json.Unmarshal(byteValue, &standardIndex)

	return &IndexHandler{
		IndexModel: &standardIndex,
	}
}

func (IndexHandler *IndexHandler) createRequestData() *models.IRequestData {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(IndexHandler.IndexModel)
	requestData := models.IRequestData(&buf)
	return &requestData
}

func (IndexHandler *IndexHandler) CreateIndex() ([]byte, error) {

	requestData := IndexHandler.createRequestData()
	req, err := http.NewRequest("POST", constants.ZINC_HOST+"/api/index", *requestData)
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
