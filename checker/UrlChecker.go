package checker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UrlChecker struct {
	url string
}

type UrlCheckRules struct {
	Name string
	Url string
	Recipient []string
}

func (urlChecker *UrlChecker) SetUrl(url string)  {
	urlChecker.url=url
}

func (urlChecker *UrlChecker) Check() (error,string) {

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(urlChecker.url)

	if err != nil {
		return errors.New(fmt.Sprintf("request errror %s",err.Error())),""
	}
	if resp.StatusCode!=200 {
		return  errors.New(fmt.Sprintf("http status code:%d",resp.StatusCode)),"";
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	return err,string(result)
}