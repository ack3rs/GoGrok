package main

import (
	l "log/slog"
	"os"

	"GoGrok/GoGrok"
	"GoGrok/environment"
)

var HOSTNAME, _ = os.Hostname()
var VERSION = "Development"

func main() {

	environment.SetUpEnv()
	l.With("HOSTNAME", HOSTNAME).With("VERSION", VERSION).Info("Starting")

	messages := GoGrok.Messages{
		Messages: []GoGrok.Message{
			{Role: "system", Content: "You are an AI that generates random text descriptions of rooms from a list of keywords, that will be used in a text adventure game. "},
			{Role: "user", Content: "Mountain, Snow, Old Path, Cave in the Distant"},
		},
		Model: "grok-beta"}

	Response := GoGrok.GetChatCompletion(messages)

	environment.Dump(Response)
}
