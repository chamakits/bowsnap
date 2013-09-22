package server

import (
	"github.com/gorilla/mux"
	"os"
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
func StartServer(snapshotVersion string) {
	initRegistry(snapshotVersion)
	r := mux.NewRouter()
	//r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/packages", packagesHandler)
	r.HandleFunc("/packages/{name}", specificPackageHandler)
	r.HandleFunc("/packages/search/{name}", searchHandler)
	http.Handle("/", r)
}

func initRegistry(snapshotVersionPath string){
	doesFileExist, err := checkFile(snapshotVersionPath)
	if err != nil {
		fmt.Fprintf(os.Stderr,"ERROR:  File provided (\"%s\") for bower registry snapshot does not exist.\n",snapshotVersionPath)
		fmt.Fprintf(os.Stderr,"ERROR-MESSAGE:  %v\n",err)
	}
}

func handler(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}

func packagesHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	category := vars["name"]
	
}
