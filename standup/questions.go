package standup

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func StandupQuestionaire(event *slackevents.AppMentionEvent, client *slack.Client, user *slack.User) error {

	var blocks slack.Blocks
	err := blocks.UnmarshalJSON([]byte(QuestionaireBlockJSON))
	if err != nil {
		log.Fatal(err)
	}

	msgBlock := []slack.Block{}
	msgBlock = append(msgBlock, blocks.BlockSet...)

	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionBlocks(msgBlock...))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
