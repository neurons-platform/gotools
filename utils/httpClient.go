package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func HttpPost(url string, data string) string {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		return string(body)

	}
	return ""

}

func HttpPostJsonWtithHeader(url string, data string, header map[string]string) string {
	c := http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))

	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	if Throw(err) {
		resp, err := c.Do(req)
		if Throw(err) {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if Throw(err) {
				return string(body)
			}
		}
	}
	return ""
}

func HttpPostJson(url string, data string) string {
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		return string(body)

	}
	return ""

	//fmt.Println(string(body))
}
