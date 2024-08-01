package hellohello

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf("Hi,Hi, %v. Welcome!Welcome!", name)
    return message, nil
}
