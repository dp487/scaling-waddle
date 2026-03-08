package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the request body into the provided interface
func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil { // Use io.ReadAll instead of ioutil.ReadAll
		if err := json.Unmarshal(body, x); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
