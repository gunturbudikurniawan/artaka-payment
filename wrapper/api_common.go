package wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moul.io/http2curl"
	"net/http"
	"reflect"
)

func Call(method string, url string, header *http.Header, body interface{}, result interface{}) *error {
	reqBody := []byte("")
	var err error
	var client = &http.Client{}

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())
	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	if header != nil {
		request.Header = *header
	}

	request.Header.Set("Content-Type", "application/json")

	command, _ := http2curl.GetCurlCommand(request)
	fmt.Println(command)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(respBody))

	if err != nil {
		panic(err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		panic(err)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		panic(err)
	}

	return nil

}
