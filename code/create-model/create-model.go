package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://app.nanonets.com/api/v2/ObjectDetection/Model/"
	apiKey := os.Getenv("NANONETS_API_KEY")

	payload := strings.NewReader("{\"categories\" : [\"TieFighter\", \"MillenniumFalcon\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseMap map[string]*json.RawMessage
	json.Unmarshal(body, &responseMap)

	var model_id string
	json.Unmarshal(*responseMap["model_id"], &model_id)
	fmt.Println("NEXT RUN: 'export NANONETS_MODEL_ID=" + model_id + "'")
	fmt.Println("THEN RUN: go build object-detection-sample-golang/code/upload-training && ./upload-training")
}
