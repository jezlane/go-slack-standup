package standup

const QuestionaireBlockJSON = `
[
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*Daily Stand-Up*\n\n:information_source: All Three Questions are Required"
		}
	},
	{
		"type": "divider"
	},
	{
		"type": "input",
		"block_id": "question001",
		"element": {
			"type": "plain_text_input",
			"multiline": true,
			"action_id": "plain_text_input-action"
		},
		"label": {
			"type": "plain_text",
			"text": "Q1 - What did you accomplish since the last check-in?"
		}
	},
	{
		"type": "divider"
	},
	{
		"type": "input",
		"block_id": "question002",
		"element": {
			"type": "plain_text_input",
			"multiline": true,
			"action_id": "plain_text_input-action"
		},
		"label": {
			"type": "plain_text",
			"text": "Q2 - What are you working on next?"
		}
	},
	{
		"type": "divider"
	},
	{
		"type": "input",
		"block_id": "question003",
		"element": {
			"type": "radio_buttons",
			"options": [
				{
					"text": {
						"type": "plain_text",
						"text": ":white_check_mark: Not Blocked",
						"emoji": true
					},
					"value": "question003-value0"
				},
				{
					"text": {
						"type": "plain_text",
						"text": ":wave: Need Assistance",
						"emoji": true
					},
					"value": "question003-value1"
				},
				{
					"text": {
						"type": "plain_text",
						"text": ":octagonal_sign: Blocked",
						"emoji": true
					},
					"value": "question003-value2"
				}
			],
			"action_id": "radio_buttons-action"
		},
		"label": {
			"type": "plain_text",
			"text": "Q3 - What's your status moving forward?",
			"emoji": true
		}
	},
	{
		"type": "input",
		"block_id": "question004",
		"element": {
			"type": "plain_text_input",
			"multiline": true,
			"action_id": "plain_text_input-action"
		},
		"label": {
			"type": "plain_text",
			"text": ":speech_balloon:  Please elaborate on your choice.(optional)",
			"emoji": true
		}
	},
	{
		"type": "divider"
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": " "
		},
		"accessory": {
			"type": "button",
			"text": {
				"type": "plain_text",
				"text": "Submit",
				"emoji": true
			},
			"value": "click_me_123",
			"action_id": "actionId-0"
		}
	}
]
`
