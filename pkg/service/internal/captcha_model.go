package internal

import (
	"encoding/json"
)

type Captcha struct {
	Key     string `json:"-"`
	Content string
	Purpose int32
}

func (c *Captcha) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
func (c *Captcha) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}
