package utils

import (
	"net/http"
	"io/ioutil"
)

func HttpGet(url string) (r string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	r = string(body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return r, nil
}

func GetStatusCode(url string) (code int, err error) {
	res, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	code = res.StatusCode
	return code, nil
}

