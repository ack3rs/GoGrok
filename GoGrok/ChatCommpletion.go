package GoGrok

import "GoGrok/environment"

func GetChatCompletion(messages Messages) {
    
    server := environment.GetEnvString("XAI_SERVER", "https://api.x.ai")
    
    request, t, x := webRequest(server + "/v1/chat/completions")
    environment.Dump(request)
    environment.Dump(t)
    environment.Dump(x)
    
}
