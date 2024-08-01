package hello

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
