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

func CloneAllGit(uri, rootPackagePath string) {

	//Get directory names.
	newCloneDir := GetGitDirectoryName(uri)
	projectPath := path.Join(rootPackagePath, newCloneDir) 

	//First clone the first one
	CloneGit(uri, rootPackagePath)

	//Read the bower.json
	bow := bower.GetBower(projectPath)
	ChangeToVersion(rootPackagePath, projectPath, newCloneDir, bow)

	
	//Get the rest of dev dependencies
}

func ChangeToVersion(rootPackagePath, projectPath, newCloneDir string, bow bower.Bower) {
	//fmt.Printf("%v\n",bow)

	//Get the version from there & Rename according to version there
	renamedDir := fmt.Sprintf("%s-%s",newCloneDir,bow.Version)
	renamedPath := path.Join(rootPackagePath, renamedDir)
	util.RenameIfExists(renamedPath)
	os.Rename(projectPath,renamedPath)
	ChangeToTag(renamedPath, bow.Version)	
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