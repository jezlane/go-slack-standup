# go-slack-standup

go-slack-standup is written in GoLang that provide a Slack Bot that can be used to run Async Daily StandUps within your Slack channel.

## Installation

```bash
git clone https://github.com/jezlane/go-slack-standup.git
```

## Packages

- github.com/slack-go/slack - Library for interacting with Slack
- github.com/gorilla/websocket - Library for WebSockets

## Usage

```bash
go run . -bottoken=[SLACK BOT TOKEN] -apptoken=[SLACK APP TOKEN]

```

### Flags
The following flags are required.

- -bottoken provides the program with your Slack Bot User Token
- -apptoken provides the program with your Slack App User Token

To set these up, refer to the [Setting up the App in Slack](documentation/SettinguptheAppinSlack.md)


## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)