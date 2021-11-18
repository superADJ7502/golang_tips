package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type CacheEntry struct {
	data []byte
	once *sync.Once
}

type QueryClient struct {
	cache map[string]*CacheEntry
	mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) []byte {
	c.mutex.Lock()
	entry, found := c.cache[name]
	if !found {
		entry = &CacheEntry{
			once: new(sync.Once),
		}
		c.cache[name] = entry
	}
	c.mutex.Unlock()

	entry.once.Do(func() {
		resp, err := http.Get("https://www.baidu.com")
		entry.data, err = ioutil.ReadAll(resp.Body)
		if err != nil {

		}
	})
	return entry.data
}

func main() {

}
