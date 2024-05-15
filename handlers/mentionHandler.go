package handlers

import (
	"fmt"
	"strings"

	"github.com/jezlane/go-slack-standup/standup"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func HandleAppMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {

	fmt.Println("Handling AppMentionEvent")
	// Grab the user name based on the ID of the one who mentioned the bot
	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	// Check if the user said StandUp to the bot
	text := strings.ToLower(event.Text)

	if strings.Contains(text, "stand") {
		err := standup.StandupQuestionaire(event, client, user)
		if err != nil {
			return err
		}
	}
	return nil
}
