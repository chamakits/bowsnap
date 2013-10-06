package git

import (
	"errors"
	"os/exec"
	"os"
	"log"
)

func CanFindGit() (bool, error) {
	gitCmd := exec.Command("git", "--version")
	_, gitErr := gitCmd.Output()
	if gitErr != nil {
		//log.Printf("ERROR:%v\n",gitErr)
		return false, errors.New("Could not find git in PATH")
	}

	//log.Printf("Found git version:%v\n",string(stdout))
	return true, nil
}

func CloneGit(uri, path string) {
	os.MkdirAll(path, 0700)
	os.Chdir(path)
	
	gitCmd := exec.Command("git", "clone", uri)
	output, gitErr := gitCmd.Output()
	if gitErr != nil {
		log.Printf("ERROR: Git executed with issues when cloning.\n")
		log.Printf("ERROR:%v\n",gitErr)
	}
	log.Printf("GIT-OUTPUT:%v\n",string(output))
	
}
