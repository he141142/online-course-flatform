package response

import "encoding/json"

func (r *httpResponse) ResponseSuccess(status int, data interface{}) {
	r.WriteHeader(status)
	r.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(r).Encode(data)
}
