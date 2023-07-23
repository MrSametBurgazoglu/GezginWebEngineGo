package NetworkManager

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type NetworkManager struct {
	client     *http.Client
	BaseUrl    string
	CurrentUrl string
}

func (receiver *NetworkManager) Initialize() {
	receiver.client = new(http.Client)
}

func (receiver *NetworkManager) GetPage(url string) []byte {
	if strings.HasPrefix(url, "http") {
		firstIndex := strings.Index(url, "//")
		if firstIndex == -1 {
			firstIndex = 0
		}
		index := strings.Index(url[firstIndex+2:], "/")
		if index == -1 {
			receiver.BaseUrl = url
		} else {
			receiver.BaseUrl = url[:index+firstIndex]
		}
		receiver.CurrentUrl = url
	} else if strings.HasPrefix(url, "www") {
		index := strings.Index(url, "/")
		if index == -1 {
			receiver.BaseUrl = url
		} else {
			receiver.BaseUrl = url[:index]
		}
		receiver.CurrentUrl = url
	}
	req, err := http.NewRequest("GET", receiver.CurrentUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "GezginWebBrowserEngine")

	resp, err := receiver.client.Do(req)
	if err != nil {
		println(err.Error())
		panic("wow there is problem")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}

func (receiver *NetworkManager) Get(url string) []byte {
	targetUrl := ""
	if strings.HasPrefix(url, "/") {
		targetUrl = receiver.BaseUrl + url
	} else if strings.HasPrefix(url, "http") {
		targetUrl = url
	} else {
		index := strings.LastIndexAny(receiver.BaseUrl, "/")
		if index == -1 {
			targetUrl = receiver.CurrentUrl + "/" + url
		} else {
			targetUrl = receiver.CurrentUrl[:index] + url
		}
	}
	println(targetUrl, "target url")
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "GezginWebBrowserEngine")

	resp, err := receiver.client.Do(req)
	if err != nil {
		println(err.Error())
		panic("wow there is problem")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}
