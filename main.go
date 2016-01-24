package main

import (
	"flag"
)

type scattrData struct {
	OutUrls []string //list of fanout Urls
}

const defaultScript = `
rule(function(data){
	return data;
})
`

var (
	adminHost  = flag.String("s", "0.0.0.0", "listen on Host for admin interface")
	adminPort  = flag.Int("q", 9090, "use port for admin interface")
	host       = flag.String("b", "0.0.0.0", "listen on Host")
	port       = flag.Int("p", 8080, "use port")
	ConfigFile = flag.String("c", "urls.toml", "use Configfile")
)

func main() {

	flag.Parse()
	banner()

	go StartAdminInterface(*adminHost, *adminPort)

	StartScattrInterface(*host, *port)

}
