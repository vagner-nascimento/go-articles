package httpcalls

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetHttp(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	pRes := &Response{}
	err = json.Unmarshal(body, pRes)

	return pRes, err
}
