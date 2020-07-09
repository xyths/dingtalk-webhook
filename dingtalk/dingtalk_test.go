package dingtalk

import (
	"os"
	"testing"
)

func TestClient_Text(t *testing.T) {
	url := os.Getenv("URL")
	secret := os.Getenv("SECRET")
	text := os.Getenv("TEXT")
	c := New(url, secret)
	err := c.Text(text)
	t.Log(err)
}
