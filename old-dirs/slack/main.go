package main

import (
	"github.com/bluele/slack"
)

const (
	hookURL = "https://hooks.slack.com/services/xxxxxx/xxxxxx/xxxxxxxxxxxxxx"
)

func main() {
	hook := slack.NewWebHook(hookURL)
	err := hook.PostMessage(&slack.WebHookPostPayload{
		Text: "hello!",
		// Channel: "#test-channel",
		Attachments: []*slack.Attachment{
			{
				Text:  "danger",
				Color: "danger",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
