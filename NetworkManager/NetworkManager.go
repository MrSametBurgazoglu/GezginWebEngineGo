package NetworkManager

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type NetworkManager struct {
	client  *http.Client
	BaseUrl string
}

func (receiver *NetworkManager) Initialize() {
	receiver.client = new(http.Client)
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
			targetUrl = receiver.BaseUrl
		} else {
			targetUrl = receiver.BaseUrl[:index]
		}
	}
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
	println(body)
	return body
}
