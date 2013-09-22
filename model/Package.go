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
}

func (packageEntry PackageEntry) ToJson() string {
	bytes, err := json.Marshal(packageEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Unable to unmarshall the package entry.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(8)
	}
	return string(bytes)
}

func InitPackages(jsonFilePath string, packagesMap *map[string]*PackageEntry, packageList *[]PackageEntry) {
	bytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Not able to read bytes from json file provided.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(6)
	}

	err = json.Unmarshal(bytes, packageList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Not able to Unmarshall slice of PackageEntry's.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(7)
	}

	for _, packageEntry := range *packageList {
		(*packagesMap)[packageEntry.Name] = &packageEntry
		fmt.Println(packageEntry)
		fmt.Println(&packageEntry)
		fmt.Println((*packagesMap)[packageEntry.Name])
	}
	fmt.Println((*packagesMap)["jquery"])
	fmt.Println((*packagesMap)["bootstrap"])
	for key, val := range *packagesMap {
		fmt.Println(key)
		fmt.Println(*val)
	}
}

func PackageListToJson(packagesList *[]PackageEntry) string {
	bytes, err := json.Marshal(*packagesList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:  Unable to unmarshall the package list.\n")
		fmt.Fprintf(os.Stderr, "ERROR-MESSAGE:  %v\n", err)
		os.Exit(8)
	}
	return string(bytes)
}
