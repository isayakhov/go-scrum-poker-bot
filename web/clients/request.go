package clients

import "encoding/json"

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Json    interface{}
}

func (r *Request) ToBytes() ([]byte, error) {
	if r.Json != nil {
		result, err := json.Marshal(r.Json)
		if err != nil {
			return []byte{}, err
		}
		return result, nil
	}

	return []byte{}, nil
}
