package requests

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetUser(n string) (string, error) {
	if len(n) == 0 || n == "" {
		return "", errors.New("empty name")
	}

	r, err := http.Get("https://api.github.com/users/" + n)

	if err != nil {
		fmt.Printf("Received error %v\n", err.Error())
		return "", err
	}

	bodyRaw, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Received error %v\n", err.Error())
		return "", err
	}

	body := string(bodyRaw)

	fmt.Println(body)
	return body, nil
}
