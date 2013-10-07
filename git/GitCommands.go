package git

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"github.com/chamakits/bowsnap/bower"
	"github.com/chamakits/bowsnap/util"
	"regexp"
	"path"
	"fmt"
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

func CloneAllGit(uri, newPath string) {
	//First clone the first one
	CloneAndToTag(uri, newPath)
	newCloneDir := GetGitDirectoryName(uri)
	
	//Read the bower.json
	projectPath := path.Join(newPath, newCloneDir)
	bow := bower.GetBower(projectPath)
	fmt.Printf("%v\n",bow)

	//Get the version from there & Rename according to version there
	renamedDir := fmt.Sprintf("%s-%s",newCloneDir,bow.Version)
	renamedPath := path.Join(newPath, renamedDir)
	os.Rename(projectPath,renamedPath)
	
	//Get the rest of dev dependencies
}

func CloneAndToTag(uri, gitPath, tag string) {
	CloneGit(uri, gitPath)
	ChangeToTag(gitPath, tag)
}

func CloneGit(uri, path string) {
	cwd := util.GetCWD()
	os.MkdirAll(path, 0700)
	os.Chdir(path)

	gitCmd := exec.Command("git", "clone", uri)
	output, gitErr := gitCmd.Output()
	if gitErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Git executed with issues when cloning.\n")
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", gitErr)
	}
	log.Printf("GIT-OUTPUT:%v\n", string(output))
	os.Chdir(cwd)
}

func ChangeToTag(gitPath, tag string) {
	cwd := util.GetCWD()
	os.Chdir(gitPath)

	gitToTag := exec.Command("git","checkout", tag)
	output, gitErr := gitToTag.Output()
	if gitErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Git executed with issues when checkouting out tag \"%s\".\n",tag)
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", gitErr)
	}
	log.Printf("GIT-OUTPUT:%v\n", string(output))

	err := os.RemoveAll(".git")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not delete .git directory.\n", err)
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
	}
	os.Chdir(cwd)
}

func GetGitDirectoryName(uri string) string {
	reg := regexp.MustCompile(`.*/(.*)\.git`)
	return reg.FindStringSubmatch(uri)[1]
}