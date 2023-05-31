package http

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Code  int
	Error string
}

//JSONError writes a simple json formatted error
func JSONError(w http.ResponseWriter, err string, code int) {
	template := jsonError{}
	template.Code = code
	template.Error = err

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(template)
}
