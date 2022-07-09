package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Gitee struct {
	Commit struct {
		Author struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		Committer struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"committer"`
		Message string `json:"message"`
		Parents []struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"parents"`
		Sha  string `json:"sha"`
		Tree struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
	} `json:"commit"`
	Content struct {
		Links struct {
			HTML string `json:"html"`
			Self string `json:"self"`
		} `json:"_links"`
		DownloadURL string `json:"download_url"`
		HTMLURL     string `json:"html_url"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		Sha         string `json:"sha"`
		Size        int64  `json:"size"`
		Type        string `json:"type"`
		URL         string `json:"url"`
	} `json:"content"`
}

var randomFactor *rand.Rand

func init() {
	randomFactor = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := randomFactor.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no image")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	base64 := base64.StdEncoding.EncodeToString(data)

	nowTime := time.Now()
	fileName := RandString(10) + ".png"

	urlStr := fmt.Sprintf("https://gitee.com/api/v5/repos/seekgo/image/contents/%d/%d/%s", nowTime.Year(), nowTime.Month(), fileName)

	form := &url.Values{}
	form.Add("access_token", "testcdfd0f00106b5c33f798cabd165dc76b")
	form.Add("branch", "master")
	form.Add("content", string(base64))
	form.Add("message", fmt.Sprintf("upload %s", fileName))

	reqest, err := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{Timeout: 200 * time.Second}
	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	real := &Gitee{}
	err = json.Unmarshal(resp, real)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(real.Content.DownloadURL)

}
