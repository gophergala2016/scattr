package main

import (
	"flag"
	// "io/ioutil"
	// "fmt"
	// "os"
)

type scattrData struct {
	OutUrls []string
}


var (
	adminHost  = flag.String("s", "0.0.0.0", "listen on Host for admin interface")
	adminPort  = flag.Int("q", 9090, "use port for admin interface")
	host       = flag.String("b", "0.0.0.0", "listen on Host")
	port       = flag.Int("p", 8080, "use port")
	ConfigFile = flag.String("c", "urls.toml", "use Configfile")
	ScriptFile = flag.String("l", "script.txt", "specify script file.")
)

const defaultScript = `
rule(function(data){
	function readSingleFile(evt) {
    var f = evt.target.files[0];
		console.log("************");
	console.log(f);
}
})
`
// func returnString() string {
// 	bytes, err := ioutil.ReadFile(*ScriptFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("FUNCTION", string(bytes))
// 	return string(bytes)
// }

func main() {
	flag.Parse()
	banner()

	go StartAdminInterface(*adminHost, *adminPort)

	StartScattrInterface(*host, *port)

}
