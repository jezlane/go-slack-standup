package handlers

import (
	"fmt"

	"github.com/jezlane/go-slack-standup/standup"

	"github.com/slack-go/slack"
)

func HandleButtonClickEvent(callback slack.InteractionCallback, api *slack.Client) error {
	fmt.Println("Handling Button Click Event")

	switch callback.ActionCallback.BlockActions[0].Value {
	case "click_me_123":
		err := standup.StandupResponse(callback, api)
		if err != nil {
			return err
		}
	}

	return nil
}
