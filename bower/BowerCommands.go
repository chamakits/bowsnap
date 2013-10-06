package bower

import (
	"fmt"
	"net/http"
	"os"
)

func GetPackageList(registryUrl string) (resp *http.Response) {
	resp, err := http.Get(registryUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Could not reach bower repository URL at \"%s\".\n", registryUrl)
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:%v\n", err)
		os.Exit(2)
	}
	return
}

func GetPackageRepo(packageName string) {

}
