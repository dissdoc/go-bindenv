package gobindenv

import "fmt"

var Version string = "0.1.0"

func BindEnv(message string) {
	fmt.Println("[LOG]" + message)
}
