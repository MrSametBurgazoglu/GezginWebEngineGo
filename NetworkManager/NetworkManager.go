package NetworkManager

import (
	"io"
	"log"
	"net/http"
)

type NetworkManager struct {
	client *http.Client
	Url    string
}

func (receiver *NetworkManager) Initialize() {
	receiver.client = new(http.Client)
}

func (receiver *NetworkManager) Get(url string) []byte {
	req, err := http.NewRequest("GET", receiver.Url+url, nil)
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
