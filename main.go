package main

import (
	"flag"
)



var (
	adminHost  = flag.String("s", "0.0.0.0", "listen on Host for admin interface")
	adminPort  = flag.Int("q", 9090, "use port for admin interface")
	host       = flag.String("b", "0.0.0.0", "listen on Host")
	port       = flag.Int("p", 8080, "use port")
	ConfigFile = flag.String("c", "config.toml", "use Configfile")
)

// const defaultScript = `
// rule(function(data){
// 	return data[0].Response
// })
// `


func main() {
	flag.Parse()
	banner()

	go StartAdminInterface(*adminHost, *adminPort)

	StartScattrInterface(*host, *port)

}
