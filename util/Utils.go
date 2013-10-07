package util

import (
	"fmt"
	"os"
	"time"
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

func RenameIfExists(path string){
	os.Rename(path, fmt.Sprintf("%s-Pre-%s",path, GetTimeStamp()))
}

const TIME_LAYOUT = "Jan-02-2006_15-04-05-MST"
func GetTimeStamp() string {
	now := time.Now()
	return now.Format(TIME_LAYOUT)
}