package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) error {
	// Membaca body dari request
	body, err := io.ReadAll(r.Body) // gunakan io.ReadAll (ioutil sudah deprecated)
	if err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}
	defer r.Body.Close()

	// Unmarshal JSON ke dalam X
	err = json.Unmarshal(body, X)
	if err != nil {
		return fmt.Errorf("error unmarshalling request body: %w", err)
	}

	return nil
}
