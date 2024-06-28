package drk_http

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

func DecodeJSONBody(r *http.Request, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return ErrInvalidType
	}
	return json.NewDecoder(r.Body).Decode(v)
}
