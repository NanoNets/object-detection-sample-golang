package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	annotationsDirectory := "annotations/json"
	imageDirectory := "images/"
	modelID := os.Getenv("NANONETS_MODEL_ID")
	apiKey := os.Getenv("NANONETS_API_KEY")
	fmt.Print("modelId: ", modelID, "\n")
	fmt.Print("apiKey: ", apiKey, "\n")
	fileList := []string{}
	err := filepath.Walk(annotationsDirectory, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || filepath.Ext(path) == ".DS_Store" {
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, filePath := range fileList {
		folder := strings.Split(filePath, "/")
		length := len(folder)

		fileName := strings.Split(folder[length-1], ".")[0]
		buffer, _ := ioutil.ReadFile(filePath)
		jsonText := string(buffer)
		url := "https://app.nanonets.com/api/v2/ObjectDetection/Model/" + modelID + "/UploadFile/"

		file, err := os.Open(imageDirectory + fileName + ".jpg")
		if err != nil {
			fmt.Println("Image not found", err)
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", filepath.Base(imageDirectory+fileName+".jpg"))
		if err != nil {
			fmt.Println("Image not found", err)
		}
		_, err = io.Copy(part, file)

		writer.WriteField("data", `[{"filename":"`+fileName+".jpg"+`", "object": `+jsonText+`}]`)
		writer.WriteField("id", modelID)

		contentType := writer.FormDataContentType()

		err = writer.Close()
		if err != nil {
			fmt.Println("Problem writing the formdata", err)
		}

		req, _ := http.NewRequest("POST", url, body)

		req.Header.Add("Content-Type", contentType)
		req.SetBasicAuth(apiKey, "")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body2, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body2))
	}
}
