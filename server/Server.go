package server

import (
	"fmt"
	"github.com/chamakits/bowsnap/model"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"log"
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

var packagesMap map[string]*model.PackageEntry
var packagesList []model.PackageEntry

const defaultPackageInitAmount = 10000

func StartServer(snapshotVersion string) {
	packagesMap = make(map[string]*model.PackageEntry, defaultPackageInitAmount)
	packagesList = make([]model.PackageEntry, defaultPackageInitAmount)
	initRegistry(snapshotVersion)
	r := mux.NewRouter()
	r.HandleFunc("/", packagesHandler)
	r.HandleFunc("/packages", packagesHandler)
	r.HandleFunc("/packages/{name}", specificPackageHandler)
	r.HandleFunc("/packages/search/{name}", searchHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8123",nil)
	fmt.Println("End of StartServer")
}

func initRegistry(snapshotVersionPath string) {
	doesFileExist := checkFileExist(snapshotVersionPath)
	if !doesFileExist {
		fmt.Fprintf(os.Stderr, "ERROR:  File provided (\"%s\") for bower registry snapshot does not exist.\n", snapshotVersionPath)
	}
	model.InitPackages(snapshotVersionPath, &packagesMap, &packagesList)
}

func checkFileExist(filePath string) bool {
	_, err := os.Lstat(filePath)
	return err == nil
}

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}

func packagesHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%v", model.PackageListToJson(&packagesList))
}

func specificPackageHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	packageName := vars["name"]
	log.Printf("URL Parameter packageName=%s\n",packageName)
	
	packageFound := packagesMap[packageName]
	log.Printf("Found package: %+v\n",packageFound)
	
	var jsonString string
	if packageFound != nil {
		jsonString = (*packageFound).ToJson()
	}
	log.Printf("Json string:\"%s\"\n",jsonString)
	fmt.Fprintf(response, jsonString)

}

func searchHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	lookup := vars["name"]
	foundPackages := make([]model.PackageEntry, 512)
	for _, currPackage := range packagesList {
		if strings.Contains(currPackage.Name, lookup) {
			foundPackages = append(foundPackages, currPackage)
		}
	}
	fmt.Fprintf(response, model.PackageListToJson(&foundPackages))
}
