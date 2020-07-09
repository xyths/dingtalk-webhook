package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	BaseUrl string
	Secret  string
}

func New(url, secret string) *Client {
	return &Client{
		BaseUrl: url,
		Secret:  secret,
	}
}

func (c *Client) Sign(timestamp int64) string {
	key := fmt.Sprintf("%d\n%s", timestamp, c.Secret)
	h := hmac.New(sha256.New, []byte(c.Secret))
	h.Write([]byte(key))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *Client) send(timestamp int64, sign string, data string) error {
	addr := fmt.Sprintf("%s&timestamp=%d&sign=%s", c.BaseUrl, timestamp, url.QueryEscape(sign))
	resp, err := http.Post(addr, "application/json", strings.NewReader(data))
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("http response is NOT 200")
	}

	dec := json.NewDecoder(resp.Body)

	var r Response
	if err := dec.Decode(&r); err != nil {
		return err
	}
	if r.Code == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("errcode: %d, errmsg: %s", r.Code, r.Message))
	}
}

func (c *Client) Text(message string) error {
	timestamp := time.Now().Unix() * 1000
	s := c.Sign(timestamp)

	text := MessageText{
		Type: "text",
		Text: Text{
			Content: message,
		},
		At: &At{
			AtAll: false,
		},
	}
	b, _ := json.Marshal(text)
	data := string(b)
	return c.send(timestamp, s, data)
}
