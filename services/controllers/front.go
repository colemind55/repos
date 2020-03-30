package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers ...
func RegisterControllers() {
	ex := newEmailExistsController()
	http.Handle("/emailExists", *ex)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
