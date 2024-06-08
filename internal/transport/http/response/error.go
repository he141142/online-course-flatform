package response

import "encoding/json"

func (r *httpResponse) ResponseError(status int, err error) {
	r.WriteHeader(status)
	r.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(r).Encode(map[string]interface{}{"error": err.Error()})
}
