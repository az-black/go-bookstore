package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		return err
	}
}