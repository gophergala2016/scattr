# scattr

This application can be used to fan out a http request to several other endpoints. The other endpoints can be configured within this application.
The http request comes to this server and will be fanned-out to the pre-defined endpoints. And finally the server collects the responses from those endpoints and returns a consolidated response.

## Usage
This application can be used to fan out a http request to several other endpoints. The other endpoints can be configured within this application.

![Tuning-image](https://github.com/gophergala2016/scattr/blob/master/screenshots/Scattr.jpg "Scattr")

## Guide

- Install golang packages as follows :

    ```
      $ go get github.com/BurntSushi/toml
      $ go get github.com/robertkrimen/otto
      $ go get github.com/mreiferson/go-httpclient
    ```
- Install the project and build it

  ```
  $ go get https://github.com/gophergala2016/scattr
  $ go build
  ```
- Start the scattr and admin servers(default set to port 8080 and 9090 respectively)

  ```
    $ skeddy [-c] [-l] [-p] [-b]
  ```

### Terminal Switches

  ``` -c urls.toml ```
    use the fanout(scattr)page(default "urls.toml")

  ``` -l script.txt ```
    get the script (default script.txt)

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
