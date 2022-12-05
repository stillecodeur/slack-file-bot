package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4482316430561-4466789675397-EtIHTz1rUCzI5hnb4qm3G7je")
	os.Setenv("CHANNEL_ID", "C04DQMJ4545")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	chanelArr := []string{os.Getenv("CHANNEL_ID")}
	filesArr := []string{"awsgsg-intro.pdf"}

	for i := 0; i < len(filesArr); i++ {
		params := slack.FileUploadParameters{
			Channels: chanelArr,
			File:     filesArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error %v", err)
		}
		fmt.Printf("Name:%s URL:%s", file.Name, file.URL)
	}
}
