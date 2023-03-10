package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// recived body in json then we have to parse
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil { //unmarshelling
			return
		}
	}
}
