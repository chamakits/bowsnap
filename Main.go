package main

import (
	"flag"
	"fmt"
	"github.com/chamakits/bowsnap/git"
	"github.com/chamakits/bowsnap/server"
	"github.com/chamakits/bowsnap/snap"
	"github.com/chamakits/bowsnap/util"
	"log"
	"os"
)

func main() {
	var (
		serverFlag          bool
		snapshotVersionFlag string
		newSnapshotFlag     bool
		newSnapshotNameFlag string
		repoUrlFlag         string
		portFlag            int
	)
	git.CanFindGit()
	//git.CloneGit("git://github.com/components/jquery.git","packages/jquery")
	git.CloneAllGit("git://github.com/components/jquery.git","packages/jquery")
	initFlags(&serverFlag, &snapshotVersionFlag, &portFlag, &newSnapshotFlag, &newSnapshotNameFlag, &repoUrlFlag)

	if serverFlag {
		log.Printf("Starting server on port:%d\n", portFlag)
		server.StartServer(snapshotVersionFlag, portFlag)
	} else if newSnapshotFlag {
		snap.TakeNewSnapshot(repoUrlFlag, newSnapshotNameFlag)
	}
}

const bowerDefaultRepo = "http://bower.herokuapp.com/packages"

func initFlags(serverFlag *bool, snapshotVersionFlag *string, portFlag *int, newSnapshotFlag *bool, newSnapshotNameFlag *string, repoUrlFlag *string) {
	//Init flag setup
	flag.BoolVar(serverFlag, "s", false, "Run the mirrored bower server.")
	flag.StringVar(snapshotVersionFlag, "sname", "LATEST.json", "OPTIONAL:Choose version of snapshot to serve.  When used, must be with the \"s\" flag as well.")
	flag.IntVar(portFlag, "port", 8199, "OPTIONAL: Port to start the server on. When used, must be with the \"s\" flag. ")

	flag.BoolVar(newSnapshotFlag, "n", false, "Create a new snapshot of bower.")
	flag.StringVar(newSnapshotNameFlag, "nname", "", "OPTIONAL: Name for new snapshot to copy. When used, must be with the \"n\" flag. ")
	flag.StringVar(repoUrlFlag, "repo", "", "OPTIONAL: Url of the bower repo to take snapshot from. When used, must be with the \"n\" flag. ")

	flag.Parse()

	//Do flag validation
	if len(os.Args) <= 1 {
		*serverFlag = true
	} else if *serverFlag && *newSnapshotFlag {

	} else if *serverFlag && len(*snapshotVersionFlag) < 1 {
		//*snapshotVersionFlag="LATEST"
		fmt.Fprintf(os.Stderr, "ERROR:  Tried to run program in \"server\" mode while providing a blank \"version\".\n")
		flag.Usage()
		fmt.Fprintf(os.Stderr, "ERROR:  Exiting because of incorrect flag.\n")
		os.Exit(1)
	} else if *newSnapshotFlag && len(*newSnapshotNameFlag) < 1 {
		*newSnapshotNameFlag = getDefSnapshotName()
	}
	if *newSnapshotFlag && len(*repoUrlFlag) < 1 {
		*repoUrlFlag = bowerDefaultRepo
	}

}

func getDefSnapshotName() string {
	return "BowerSnapshot_" + util.GetTimeStamp() + ".json"
}


