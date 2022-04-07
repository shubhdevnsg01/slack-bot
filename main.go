package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2450072143366-3359324819986-7haonCpqCjbY00dkvDimP4SU")
	os.Setenv("CHANNEL_TO", "C02E4D0T2AC")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"FINAL450.xlsx"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{

			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name:%s,Url:%s\n", file.Name, file.URL)
	}

}
