package greeting

import (
	"log"
    "fmt"

	"greeting/hello"

	hellohello "github.com/watanabetatsumi/module_practice/greeting/sample"
)

func Hello(msg string) string {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := hello.Greet(msg)
    if err != nil {
        log.Fatal(err)
    }
    message2, err := hellohello.Greet(msg)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s",message2)

    return message
}
