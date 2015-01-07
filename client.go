package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	Email(*Email) (*EmailResponse, error)
	EmailBatch([]*Email) ([]*EmailResponse, error)
}

type Error struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.ErrorCode, e.Message)
}

type client struct {
	serverToken string
}

func NewClient(ServerToken string) Client {
	return &client{ServerToken}
}

func (c *client) Email(e *Email) (*EmailResponse, error) {
	if payload, err := json.Marshal(e); err == nil {
		reader := bytes.NewReader(payload)
		if newReq, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", reader); err == nil {
			newReq.Header.Set("Content-Type", "application/json")
			newReq.Header.Set("Accept", "application/json")
			newReq.Header.Set("X-Postmark-Server-Token", c.serverToken)
			if resp, err := http.DefaultClient.Do(newReq); err == nil {
				decoder := json.NewDecoder(resp.Body)
				if resp.StatusCode == http.StatusOK {
					response := new(EmailResponse)
					if err = decoder.Decode(response); err == nil {
						return response, nil
					} else {
						return nil, err
					}
				} else {
					response := new(Error)
					if err = decoder.Decode(&response); err == nil {
						return nil, response
					} else {
						return nil, err
					}
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *client) EmailBatch(e []*Email) ([]*EmailResponse, error) {
	if payload, err := json.Marshal(e); err == nil {
		reader := bytes.NewReader(payload)
		if newReq, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", reader); err == nil {
			newReq.Header.Set("Content-Type", "application/json")
			newReq.Header.Set("Accept", "application/json")
			newReq.Header.Set("X-Postmark-Server-Token", c.serverToken)
			if resp, err := http.DefaultClient.Do(newReq); err == nil {
				decoder := json.NewDecoder(resp.Body)
				if resp.StatusCode == http.StatusOK {
					var response []*EmailResponse
					if err = decoder.Decode(&response); err == nil {
						return response, nil
					} else {
						return nil, err
					}
				} else {
					response := new(Error)
					if err = decoder.Decode(&response); err == nil {
						return nil, response
					} else {
						return nil, err
					}
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
