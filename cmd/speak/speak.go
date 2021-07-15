// Command hello-cli
package main

import (
	"fmt"
	"github.com/markbrownsword/hello-go-modules/language"
	"github.com/markbrownsword/hello-go-modules/speak"
	"time"
)

func main() {
	english := language.NewEnglish("Good Day", "Good Evening")
	result := speak.Greeting(english, time.Now())

	fmt.Println(result)
}
