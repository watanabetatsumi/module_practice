package greeting

import (
	"log"

	"training/greeting/hello"
)

func Hello(msg string) string {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := hello.Greet(msg)
    if err != nil {
        log.Fatal(err)
    }

    return message
}
