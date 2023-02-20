package main

import (
       "context"
       "os"
       "github.com/aws/aws-lambda-go/lambda"
       "encoding/json"
       "bytes"
       "fmt"
       "net/http"
)

func HandleRequest(ctx context.Context) {

      url := "https://webhook.site/d2876b61-681d-43cf-8435-00f43cd36203"
      fmt.Println("HTTP json post:", url)

      dataMap := map[string]interface{}{
      		"application_name": os.Getenv("APPLICATION_NAME"),
      		"application_description": os.Getenv("APPLICATION_DESCRIPTION"),
      		"application_language": os.Getenv("APPLICATION_LANGUAGE"),
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