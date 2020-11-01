package requests

import (
	"fmt"
	"os"
	"time"
)

var level = os.Getenv("GO_REQUEST_LOGS")

func log(msg string)  {
	if level != ""{
		fmt.Println(time.Now().String(), "[ GO-REQUEST ]", msg)
	}
}
