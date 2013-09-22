package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PackageEntry struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	//This does caching at a very basic level.  It can be done currently, because CURRENTLY, no editing is allowed for a PackageEntry
	//SO HUGE NOTE:  If any editing is eventually allowed for a PackageEntry, then it must nil out this cache after the changes.
	jsonString *string
}

func (packageEntry *PackageEntry) ToJson() (retString string) {
	if (*packageEntry).jsonString == nil{
		bytes, err := json.Marshal(packageEntry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:  Unable to unmarshall the package entry.\n")
			fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
			os.Exit(8)
		}
		retString = string(bytes)
		(*packageEntry).jsonString = &retString
	} else {
		retString = *(*packageEntry).jsonString
	}
	return retString
}

func InitPackages(jsonFilePath string, packagesMap *map[string]PackageEntry, packageList *[]PackageEntry) {
	bytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Not able to read bytes from json file provided.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(6)
	}
	//var packageList []PackageEntry
	err = json.Unmarshal(bytes, packageList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Not able to Unmarshall slice of PackageEntry's.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(7)
	}

	for _, packageEntry := range *packageList {
		(*packagesMap)[packageEntry.Name] = packageEntry
	}
}

//This does caching at a very basic level.  It can be done currently, because CURRENTLY, no editing is allowed for the list of all Packages
//SO HUGE NOTE:  If any editing is eventually allowed for the package list, then it must nil out this cache after the changes.
var packagesListJson *string

func PackageListToJson(packagesList *[]PackageEntry) (retString string) {
	
	if packagesListJson == nil {
		bytes, err := json.Marshal(*packagesList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:  Unable to unmarshall the package list.\n")
			fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
			os.Exit(8)
		}
		retString = string(bytes)
		packagesListJson = &retString
	} else {
		retString = *packagesListJson
	}
	
	return 
}
