package dingtalk_webhook

import (
	"github.com/royeo/dingrobot"
	"os"
	"testing"
)

func TestSendText(t *testing.T) {
	url := os.Getenv("URL")
	secret := os.Getenv("SECRET")
	text := os.Getenv("TEXT")

	robot := dingrobot.NewRobot(url)
	robot.SetSecret(secret)

	atMobiles := []string{"13810325645"}
	isAtAll := false

	err := robot.SendText(text, atMobiles, isAtAll)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}
