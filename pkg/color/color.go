package color

import (
	"math/rand"
)

// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[90m"
var White  = "\033[97m"
var LightGray = "\033[37m"
var LightRed = "\033[91m"
var LightGreen = "\033[92m"
var LightYellow = "\033[93m"
var LightBlue = "\033[94m"
var LightPurple = "\033[95m"
var LightCyan = "\033[96m"


func Random() string {
	colors := []string{Red, Green, Yellow, Blue, Purple, Cyan, Gray, White, LightGray, LightRed, LightGreen, LightYellow, LightBlue, LightPurple, LightCyan}
	return colors[rand.Intn(len(colors))]
}
