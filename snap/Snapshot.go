package snap

import (
	"fmt"
	"github.com/chamakits/bowsnap/bower"
	"io/ioutil"
	"os"
	"path"
)

func TakeNewSnapshot(repoUrl string, newSnapshotNameFlag string) {
	resp := bower.GetPackageList(repoUrl)

	readBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Could not read response from repository URL at \"%s\".\n", repoUrl)
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(3)
	}
	//fmt.Println(string(readBytes))
	currentWorkingDirectory := getCWD()
	writeFile(&readBytes, currentWorkingDirectory, newSnapshotNameFlag)
	createLatestSymlink(currentWorkingDirectory, newSnapshotNameFlag)

}

func getCWD() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Could not get working directory.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:%v\n", err)
		os.Exit(4)
	}
	return currentWorkingDirectory
}

func writeFile(readBytes *[]byte, currentWorkingDirectory string, newFileName string) {
	snapshotFilePath := path.Join(currentWorkingDirectory, newFileName)
	ioutil.WriteFile(snapshotFilePath, *readBytes, 0644)
}

func createLatestSymlink(currentWorkingDirectory string, newFileName string) {
	symlinkPath := path.Join(currentWorkingDirectory, "LATEST.json")
	snapshotFilePath := path.Join(currentWorkingDirectory, newFileName)
	_, err := os.Lstat(symlinkPath)
	var linkPathExists, isLink bool
	//Checking that the link's path exists.  Doesn't guarantee its a symlink though, just that the file exists.
	if err == nil {
		linkPathExists = true
	}

	//If the file exists, now check if it is actually a symlink.
	if linkPathExists {
		_, err = os.Readlink(symlinkPath)
		if err == nil {
			isLink = true
		} else {
			fmt.Fprintf(os.Stderr, "ERROR:  Found existing file which is not a symlink at path \"%s\".\n", symlinkPath)
			fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:%v\n", err)
			os.Exit(5)
		}
	}

	if linkPathExists && isLink {
		os.Remove(symlinkPath)
	}

	os.Symlink(snapshotFilePath, symlinkPath)
}
