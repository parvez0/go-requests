package requests

import (
	"fmt"
	"os"
	"time"
)

var level = os.Getenv("GO_REQUEST_LOGS")

func log(msg string)  {
	if level != ""{
		currentTime := time.Now().Format("2006-01-02T15:04:05Z")
		fmt.Println(currentTime, "[ GO-REQUEST ]", msg)
	}
}
