package reporter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DingtalkGroupMessage struct {
	accessToken string
	secret string
}

type DingtalkGroupConfig struct {
	Id string
	Name string
	AccessToken string
	Secret string
}


func (dingtalkGroupMessage *DingtalkGroupMessage) SetAccessToken(accessToken string )  {
	dingtalkGroupMessage.accessToken=accessToken
}

func (dingtalkGroupMessage *DingtalkGroupMessage) SetSecret(secret string)  {
	dingtalkGroupMessage.secret=secret
}

func (dingtalkGroupMessage *DingtalkGroupMessage) SendMessage(messageContent string) string  {
	text := map[string]string{
		"content": messageContent,
	}
	postData := map[string]interface{}{
		"msgtype": "text",
		"text":    text,
	}
	postContent, _ := json.Marshal(postData)
	timestamp := time.Now().UnixMilli()
	requestUrl:=fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token="+dingtalkGroupMessage.accessToken+"&timestamp=%v&sign=%s",timestamp,dingtalkGroupMessage.makSign(timestamp))
	resp, err := http.Post(requestUrl,
		"application/json",
		strings.NewReader(string(postContent)))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func (dingtalkGroupMessage *DingtalkGroupMessage) makSign(timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, dingtalkGroupMessage.secret)
	hash := hmac.New(sha256.New, []byte(dingtalkGroupMessage.secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	return sign
}