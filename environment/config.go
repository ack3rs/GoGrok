package environment

import (
	"bytes"
	"os"
	"runtime"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var HOSTNAME, _ = os.Hostname()
var VERSION = "Development"
var runningConfig = make(map[string]string)

func init() {
	SetUpLogging()
}

func GetEnvString(key string, EnvDefault string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return EnvDefault
}

func GetEnvInt(key string, EnvDefault int) int {
	if os.Getenv(key) != "" {
		intValue, err := strconv.Atoi(os.Getenv(key))
		if err == nil {
			return EnvDefault
		}
		return intValue
	}
	return EnvDefault
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GetRunningConfig(key string, defaultValue string) string {

	if runningConfig[key] != "" {
		return runningConfig[key]
	}

	if GetEnvString(key, defaultValue) != "" {
		runningConfig[key] = GetEnvString(key, defaultValue)
		return GetEnvString(key, defaultValue)
	}

	return defaultValue
}

func SetRunningConfig(key string, defaultValue string) {

	runningConfig[key] = GetEnvString(key, defaultValue)

}
