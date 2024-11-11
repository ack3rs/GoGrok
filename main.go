package main

import (
    "os"
    
    "GoGrok/environment"
)

var HOSTNAME, _ = os.Hostname()
var VERSION = "Development"

func main() {
    
    environment.SetUpEnv()
    
}
