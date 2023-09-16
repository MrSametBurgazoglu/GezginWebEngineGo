package ResourceManager

import (
	"gezgin_web_engine/NetworkManager"
	"os"
)

type Resource struct {
	data []byte
}

func (receiver *Resource) GetData() []byte {
	return receiver.data
}

type ResourceManager struct {
	NetworkManager *NetworkManager.NetworkManager
	cache          map[string]*Resource
	Online         bool
}

func (receiver *ResourceManager) Initialize() {
	receiver.cache = make(map[string]*Resource)
}

func (receiver *ResourceManager) CreateResourceFromWeb(url string) {
	data := receiver.NetworkManager.Get(url)
	resource := Resource{data: data}
	receiver.cache[url] = &resource
}

func (receiver *ResourceManager) CheckResource(url string) bool {
	if !receiver.Online {
		return true
	}
	if receiver.cache[url] != nil {
		return true
	}
	return false
}

func (receiver *ResourceManager) GetResource(url string) (*Resource, error) {
	if !receiver.Online {
		file, err := os.ReadFile("exampleHtmlFiles/" + url)
		if err != nil {
			return nil, err
		}
		return &Resource{data: file}, nil
	}
	return receiver.cache[url], nil
	/* try to get resource from cache
	if its not in cache then look if internal get from file manager if its not get from network manager
	*/
}
