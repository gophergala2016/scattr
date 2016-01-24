# scattr

This is a light weight server that users an incoming http request as a trigger to initiate http requests to one or more configured end points, waits until it receives responses for each such request and then combines the individual responses to prepare the final consolidated response for the original request. Custom logic can be scripted via JavaScript to control how the individual responses are processed.
Scattr has builtin admin interface to easily  add / remove endpoint and filter scripts .

The following schematic gives a very high-level overview of the workings of **Scattr**

![Scattr-image](https://github.com/gophergala2016/scattr/blob/master/screenshots/scattr.png "Scattr")
![Scattr-image](https://github.com/gophergala2016/scattr/blob/master/screenshots/screenshot.png "Scattr")

## Installation
To install **Scattr** simply run

```
$ go get https://github.com/gophergala2016/scattr
```

This will download Scattr to `$GOPATH/src/github.com/gophergala2016/scattr`. Now cd to this directory and run
`go build` to create `scattr` binary. Place this binary somewhere in PATH to install it.


## Usage
To run scattr simply execute the binary in any folder using the following command-line

```
$ scattr
```

Scattr can be run from command-line using the following syntax.

```
$ scattr -h
Usage of scattr:
  -b string
    	listen on Host (default "0.0.0.0")
  -c string
    	use Configfile (default "config.toml")
  -p int
    	use port (default 8080)
  -q int
    	use port for admin interface (default 9090)
  -s string
    	listen on Host for admin interface (default "0.0.0.0")

```
for example to run scattr

Sample script to filter the response :

```
eval(function(data){
  return data.Url
});

```

### Command Line Flags

  ``` -c urls.toml ```
    use the fanout(scattr)page(default "config.toml")

  ``` -p port ```
    scattr port to listen (default 8080)

  ``` -b host-name ```
    bind to ip address (default "0.0.0.0")

  ``` -q port ```
    admin port to listen (default 9090)

  ``` -s host-name ```
    bind to ip address (default "0.0.0.0")

# License

MIT
