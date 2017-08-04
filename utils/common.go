package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteResponse marshalls json into an object and writes server response
func WriteResponse(w http.ResponseWriter, code int, object interface{}) {
	data, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, string(data))
}
