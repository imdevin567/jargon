package adapter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/imdevin567/jargon/protocol"
)

// HTTPAdapter ...
type HTTPAdapter struct {
	*Adapter
	Path        string
	ContentType string
	localChan   chan []byte
}

// NewHTTPAdapter ...
func NewHTTPAdapter(direction Direction, host string, port int, path, contentType string) *HTTPAdapter {
	adapter := NewAdapter(direction, host, port, HTTP)
	return &HTTPAdapter{
		Adapter:     adapter,
		Path:        path,
		ContentType: contentType,
	}
}

// Start ...
func (httpAdapter *HTTPAdapter) Start() {
	fmt.Println("No setup required for HTTP adapter")
}

// Input = create server ...
func (httpAdapter *HTTPAdapter) Input(c chan []byte) {
	httpAdapter.localChan = c
	mux := http.NewServeMux()
	mux.HandleFunc(httpAdapter.Path, httpAdapter.handlePost)
	http.ListenAndServe(fmt.Sprintf(":%v", httpAdapter.Port), mux)
}

// Output = HTTP POST ...
func (httpAdapter *HTTPAdapter) Output(c chan []byte) {
	url := fmt.Sprintf("http://%v:%v%v", httpAdapter.Host, httpAdapter.Port, httpAdapter.Path)
	for {
		msg := <-c
		_, err := http.Post(url, httpAdapter.ContentType, bytes.NewBuffer(msg))
		if err != nil {
			// TODO: Handle error
		}
	}
}

// handlePost ...
func (httpAdapter *HTTPAdapter) handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		} else {
			httpAdapter.localChan <- body
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
