package microkit

import (
    "fmt"
    "log"
    "net/http"
)

func MockServer (jsonv string, serverPort int, uri string)  {
	go func (jsonv string) {
		http.HandleFunc(fmt.Sprintf("/%s", uri), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s",jsonv)
			
		})
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", serverPort), nil))
	}(jsonv)
}