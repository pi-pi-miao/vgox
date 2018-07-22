package proxy

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"web/defs"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

//api透传
func Request(b *defs.ApiBody, w http.ResponseWriter, r *http.Request) {

	var err error
	var resp *http.Response
	switch b.Method {
	case http.MethodGet:
		//get所以body为nil
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err%v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err%v", err)
			return
		}
		normalResponse(w, resp)

	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("err%v", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")
		return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(defs.ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}
	w.WriteHeader(200)
	io.WriteString(w, string(res))
}
