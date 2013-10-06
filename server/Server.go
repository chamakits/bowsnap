package server

import (
	"fmt"
	"github.com/chamakits/bowsnap/model"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

/**
Basing it on: https://github.com/bower/registry/
URLs:
get '/packages' do
get '/packages/:name' do
get '/packages/search/:name' do
post '/packages' do
Package.create(
:name => params[:name],
:url  => params[:url]
)
**/

var packagesMap map[string]model.PackageEntry
var packageList []model.PackageEntry

const defaultPackageInitAmount = 10000

func StartServer(snapshotVersion string, port int) {
	packagesMap = make(map[string]model.PackageEntry, defaultPackageInitAmount)
	packageList = make([]model.PackageEntry, defaultPackageInitAmount)
	initRegistry(snapshotVersion)
	r := mux.NewRouter()
	r.Schemes("https")
	//r.HandleFunc("/", packagesHandler)
	r.HandleFunc("/packages", packagesHandler)
	r.HandleFunc("/packages/{name}", specificPackageHandler)
	r.HandleFunc("/packages/search/{name}", searchHandler)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Println("End of StartServer")
}

func initRegistry(snapshotVersionPath string) {
	doesFileExist := checkFileExist(snapshotVersionPath)
	if !doesFileExist {
		fmt.Fprintf(os.Stderr, "ERROR:  File provided (\"%s\") for bower registry snapshot does not exist.\n", snapshotVersionPath)
	}
	model.InitPackages(snapshotVersionPath, &packagesMap, &packageList)
}

func checkFileExist(filePath string) bool {
	_, err := os.Lstat(filePath)
	return err == nil
}

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}

var allPackageKey *string

func packagesHandler(response http.ResponseWriter, request *http.Request) {
	if allPackageKey == nil {
		namesKey := getPackagePath(request)
		allPackageKey = &namesKey
	}

	fmt.Fprintf(response, "%v", model.PackageListToJson(&packageList, *allPackageKey))
}

func getPackagePath(request *http.Request) string {
	return request.URL.Path
}

func specificPackageHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	packageName := vars["name"]

	packageFound, found := packagesMap[packageName]

	var jsonString string
	if found {
		jsonString = packageFound.ToJson()
	}
	fmt.Fprintf(response, jsonString)

}

func searchHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	lookup := vars["name"]
	//log.Printf("SearchHandler, name:%s\n", lookup)
	namesString := ""
	foundPackages := make([]model.PackageEntry, 0)
	for _, currPackage := range packageList {
		//		log.Printf("Lookup:%s; currPackage:%v\n",lookup, currPackage)
		//		log.Printf("Currpackage.Name:\"%s\", Lookup:\"%s\", contains:\"%v\"\n", currPackage.Name, lookup, strings.Contains(currPackage.Name, lookup))
		if strings.Contains(currPackage.Name, lookup) {
			foundPackages = append(foundPackages, currPackage)
			namesString = namesString + currPackage.Name
		}
	}
	//	log.Printf("Found size:%s\n", len(foundPackages))
	fmt.Fprintf(response, model.PackageListToJson(&foundPackages, getPackagePath(request)))
}
