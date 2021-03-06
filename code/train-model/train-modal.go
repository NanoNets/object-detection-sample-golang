package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	modelId := os.Getenv("NANONETS_MODEL_ID")
	apiKey := os.Getenv("NANONETS_API_KEY")

	url := "https://app.nanonets.com/api/v2/ObjectDetection/Model/" + modelId + "/Train/"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("body: ", string(body))
	fmt.Println("NEXT RUN: go build object-detection-sample-golang/code/model-state && ./model-state")
}
