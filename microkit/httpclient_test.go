package microkit

import (
	"testing"
	"reflect"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

// func setupSuite(tb testing.TB) func(tb testing.TB) {
// 	jsonv := `{"db":{"username":{"type": "string", "value": "test"}}}`
// 	MockServer(jsonv, 8081)

// 	// Return a function to teardown the test
// 	return func(tb testing.TB) {
// 		log.Println("teardown suite")
// 	}
// }

func TestHttpClientReturnValidJson(t *testing.T)  {
	jsonv := `{"db":{"username":{"type": "string", "value": "test"}}}`
	uri := "1"
	MockServer(jsonv, 8081, uri)
	client := new(HttpClient)
	resp := client.Get(strings.TrimSpace(fmt.Sprintf("http://localhost:8081/%s", uri)),"test_key")
	if reflect.TypeOf(resp).String() != "*http.Response" {
		t.Errorf("got %s want %s", reflect.TypeOf(resp), "*http.Response")
	}
}

func TestHttpClient(t *testing.T)  {
	jsonv := `{"db":{"username":{"type": "string", "value": "test"}}}`
	uri := "2"
	MockServer(jsonv, 8082, uri)
	client := new(HttpClient)
	resp := client.Get(strings.TrimSpace(fmt.Sprintf("http://localhost:8082/%s", uri)),"test_key")
	bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)
	if bodyString != jsonv {
		t.Errorf("got %s want %s", jsonv, bodyString)
	}
}