package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	auth := getAuth()
	if auth.Code != 0 {
		return
	}
	data := auth.Data
	data.Host = "http://10.33.196.197:82"
	uObj, _ := url.Parse(data.Host)

	baseURL := &cos.BaseURL{BucketURL: uObj}
	client := cos.NewClient(baseURL, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:     data.TmpSecretID,
			SecretKey:    data.TmpSecretKey,
			SessionToken: data.SessionToken,
		},
	})

	path := fmt.Sprintf("%s/%s/go.mod", data.Bucket, data.Prefix)
	localPath := "./go.mod"
	_, err := client.Object.PutFromFile(context.Background(), path, localPath, nil)
	if err != nil {
		panic(err.Error())
	}

}

type VrAuth struct {
	RequestID    string `json:"request_id"`
	TraceID      string `json:"trace_id"`
	BusinessCode string `json:"business_code"`
	OsiRequestID string `json:"osi_request_id"`
	Code         int    `json:"code"`
	Status       string `json:"status"`
	Data         struct {
		AppID        string `json:"app_id"`
		Bucket       string `json:"bucket"`
		DownloadHost string `json:"download_host"`
		DownloadType string `json:"download_type"`
		Expire       int    `json:"expire"`
		Host         string `json:"host"`
		IsAccelerate string `json:"is_accelerate"`
		Prefix       string `json:"prefix"`
		Region       string `json:"region"`
		SessionToken string `json:"sessionToken"`
		TmpSecretID  string `json:"tmpSecretId"`
		TmpSecretKey string `json:"tmpSecretKey"`
		TTL          int    `json:"ttl"`
	} `json:"data"`
	Cost int `json:"cost"`
}

func getAuth() *VrAuth {
	retData := &VrAuth{}
	url := "http://10.33.196.197/__proxy__/vrauth/auth"
	method := "POST"

	data := map[string]interface{}{
		"ak":              "AAA",
		"auth_folder_uri": "/",
		"idenname":        "AAA",
		"expires":         "46374653",
	}

	bytes, _ := json.Marshal([]string{"auth_folder_id"})
	data["primaryids"] = string(bytes)

	bytes, _ = json.Marshal(data)

	payload := strings.NewReader(string(bytes))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err := json.Unmarshal(body, retData); err != nil {
		return nil
	}

	return retData
}
