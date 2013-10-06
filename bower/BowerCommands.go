package bower

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/chamakits/bowsnap/model"
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

func GetPackageRepo(registryUrl, packageName string) string {
	resp, err := http.Get(fmt.Sprintf("%s/%s",registryUrl, packageName))

	var pack model.PackageEntry
	readBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(readBytes, pack)
	if err != nil {
		fmt.Fprintf(os.Stderr,"ERROR: Could not unmarshall json.\n")
		fmt.Fprintf(os.Stderr,"ERROR:%v\n",err)
	}

	return pack.Url
}
 
func GetPackageVersion(packageName string) {
	//Read bower.json to get version.
}