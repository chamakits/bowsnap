package util

import (
	"os"
	"fmt"
)

func GetCWD() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Could not get working directory.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:%v\n", err)
		os.Exit(4)
	}
	return currentWorkingDirectory
}
