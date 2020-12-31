package DCEAPI

import (
	"net/http"
	"strings"
	"io/ioutil"
	_ "net/url"
	_ "encoding/json"
	"fmt"
	_ "bytes"
)

type Exchange struct {
	Name string
	Apikey string
	Secret string
	Password string
	LimitRate bool
	Debug bool
}

func BaseRequest(method, path string, params, body, headers map[string]string) (string, error) {
	client := &http.Client{}
	fmt.Println(body)
	req, err := http.NewRequest(method, path, strings.NewReader(""))
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	req.Header.Add("User-Agent", "Mozilla")
	req.Header.Add("content-type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
