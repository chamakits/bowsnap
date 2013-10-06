package git

import (
	"errors"
	"os/exec"
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

func CloneGit(uri string) {

}
