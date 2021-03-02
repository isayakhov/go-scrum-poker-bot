package clients

import "encoding/json"

type Response struct {
	Status  int
	Headers map[string][]string
	Body    []byte
	Error   error
}

func (r *Response) Json(to interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	return json.Unmarshal(r.Body, to)
}
