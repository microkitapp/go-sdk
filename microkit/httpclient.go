package microkit

import (
	"net/http"
	"fmt"
	"time"
)

type HttpClient struct {}

func (c HttpClient) Get (url string, key string) *http.Response {
	client := &http.Client{
        Timeout: time.Second * 10,
    }
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Printf("Got error %s", err.Error())
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
    
    res, err := client.Do(req)
    if err != nil {
		fmt.Printf("Got error %s", err.Error())
    }
	

	return res
} 