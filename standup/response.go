package standup

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

func StandupResponse(callback slack.InteractionCallback, api *slack.Client) error {
	user, err := api.GetUserInfo(callback.User.ID)
	if err != nil {
		return err
	}
	user_slug := user.Profile.DisplayName + "(" + user.Name + ")"

	sourceThreadTimestamp := callback.Message.Timestamp
	fmt.Println("sourceThreadTimestamp", string(sourceThreadTimestamp))

	stateMap := callback.BlockActionState.Values
	q1 := stateMap["question001"]["plain_text_input-action"].Value
	q2 := stateMap["question002"]["plain_text_input-action"].Value
	q3 := stateMap["question003"]["radio_buttons-action"].SelectedOption.Text.Text
	q4 := stateMap["question004"]["plain_text_input-action"].Value

	fmt.Println("StandUp Button Clicked by", user.Profile.DisplayName, "(", user.Name, ")")
	fmt.Println(q1, q2, q3, q4)

	ResponseJSON := fmt.Sprintf(ResponseBlockJSON, user_slug, q1, q2, q3, q4)
	var blocks slack.Blocks

	err = blocks.UnmarshalJSON([]byte(ResponseJSON))
	if err != nil {
		log.Fatal(err)
	}

	responseBlock := []slack.Block{}
	responseBlock = append(responseBlock, blocks.BlockSet...)

	_, _, err = api.PostMessage(callback.Channel.ID,
		slack.MsgOptionBlocks(responseBlock...),
		slack.MsgOptionTS(string(sourceThreadTimestamp)),
	)
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil

}
