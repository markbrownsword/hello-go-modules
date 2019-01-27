// Command hello-cli
package main

import (
	"fmt"
	"github.com/markbrownsword/hello-go-modules/speak"
	"time"
)

func main() {
	english := speak.NewEnglish("Good Day", "Good Evening")
	result := speak.Greeting(english, time.Now())

	fmt.Println(result)
}
