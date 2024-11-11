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
    
    M := GoGrok.Message{Role: "system", Content: "You are a test Assistant"}
    messages := GoGrok.Messages{Messages: []GoGrok.Message{M}, Model: "grok-beta"}
    
    GoGrok.GetChatCompletion(messages)
    
}
