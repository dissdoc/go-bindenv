package internal

import (
	"fmt"
	"os"
)

func MustParam(param string) string {
	value := os.Getenv(param)
	if value == "" {
		panic(fmt.Sprintf("%s is not set", param))
	}

	return value
}
