package bower

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	"path"
	"log"
)

type Bower struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Keywords    []string `json:"keywords"`
	Ignore    []string `json:"ignore"`
	Main        string `json:"main"`
	License string `json:"license"`
}

func GetBower(projectPath string) Bower {
	bowerFile := path.Join(projectPath, "bower.json")
	bytes, err := ioutil.ReadFile(bowerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not read file '%s'\n", bowerFile)
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
	}
	log.Printf("Read json:%s\n",string(bytes))
	var bow Bower
	err = json.Unmarshal(bytes, &bow)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not unmarshall '%s'\n", bowerFile)
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
	}
	return bow
}
