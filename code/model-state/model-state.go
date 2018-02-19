package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	modelId := os.Getenv("NANONETS_MODEL_ID")
	apiKey := os.Getenv("NANONETS_API_KEY")

	fmt.Println(modelId, apiKey)

	url := "https://app.nanonets.com/api/v2/ObjectDetection/Model/" + modelId

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, "")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseMap map[string]*json.RawMessage
	json.Unmarshal(body, &responseMap)

	var state int
	var status string
	json.Unmarshal(*responseMap["state"], &state)
	json.Unmarshal(*responseMap["status"], &status)
	fmt.Println("state: ", state, "status: ", status)
	if state != 5 {
		fmt.Println("The model isn't ready yet, it's status is:", status)
		fmt.Println("We will send you an email when the model is ready. If your imapatient, run this script again in 10 minutes to check.")
	} else {
		fmt.Println("NEXT RUN: go build object-detection-sample-golang/code/prediction && ./prediction ./images/videoplayback0051.jpg")
	}
}
