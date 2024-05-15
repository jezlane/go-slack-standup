package standup

const ResponseBlockJSON = `
[
	{
		"type": "header",
		"text": {
			"type": "plain_text",
			"text": "SU Report - %s",
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
			"text": "*What did you accomplish since the last check-in?* \n %s \n"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*What are you working on next?* \n %s \n"
		}
	},
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "*What's your status moving forward?*\n %s \n %s \n"
		}
	}
]
`
