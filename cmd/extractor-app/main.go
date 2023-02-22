package main

import (
    "github.com/aws/aws-lambda-go/lambda"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func HandleRequest(ctx context.Context) {

	url := "http://ec2-44-211-82-189.compute-1.amazonaws.com/v1/application"
	fmt.Println("HTTP json post:", url)

	dataMap := map[string]interface{}{
		"application_owner":       os.Getenv("APPLICATION_OWNER"),
		"application_name":        os.Getenv("APPLICATION_NAME"),
		"application_description": os.Getenv("APPLICATION_DESCRIPTION"),
		"application_language":    os.Getenv("APPLICATION_LANGUAGE"),
	}

	jsonBytes, _ := json.Marshal(dataMap)
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
}

func main() {
	lambda.Start(HandleRequest)
}
