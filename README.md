# scattr

Scattr can be used to scatter a HTTP request to several endpoints and gather the responses and pass it back to the client.
Scatter can be configured to add javascript filter which can further modify the gathered responses to make it useful to the clients. Scattr has builtin admin interface to easily  add / remove endpoint and filter scripts .

The following schematic gives a very high-level overview of the workings of **Scattr**

![Scattr-image](https://github.com/gophergala2016/scattr/blob/master/screenshots/Scattr.jpg "Scattr")

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
    	use Configfile (default "urls.toml")
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

## Installation
  ```
  $ go get https://github.com/gophergala2016/scattr
  $ go build
  ```
- Start the scattr and admin servers(default set to port 8080 and 9090 respectively)

  ```
    $ scattr [-c] [-p] [-b]
  ```

### Terminal Switches

  ``` -c urls.toml ```
    use the fanout(scattr)page(default "urls.toml")

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
