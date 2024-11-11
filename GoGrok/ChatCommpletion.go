package GoGrok

import (
	"encoding/json"
	"fmt"
	l "log/slog"

	"GoGrok/environment"
)

func GetChatCompletion(messages Messages) MessageResponse {

	out := MessageResponse{}
	server := environment.GetEnvString("XAI_SERVER", "https://api.x.ai")

	payload, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Error marshaling messages:", err)
		l.With("Error", err).Error("Error marshaling payload")
		return out
	}

	request, _, _ := webRequest(server+"/v1/chat/completions", payload)

	err = json.Unmarshal(request, &out)
	if err != nil {
		l.With("Error", err).Error("Error unmarshaling response")
		return MessageResponse{}
	}

	return out
}
