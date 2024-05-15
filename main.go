package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

const slackDebug = false

func main() {
	var slackBotToken string
	var slackAppToken string
	flag.StringVar(&slackBotToken, "bottoken", "x", "Slack Bot Token Key")
	flag.StringVar(&slackAppToken, "apptoken", "x", "Slack App Token Key")
	flag.Parse()
	fmt.Println("Slack Bot Token Key:", slackBotToken)
	fmt.Println("Slack App Token Key:", slackAppToken)

	api := slack.New(
		slackBotToken,
		slack.OptionDebug(slackDebug),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(slackAppToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(slackDebug),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	go clientEventManage(api, client)

	client.Run()
}
