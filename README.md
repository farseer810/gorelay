# gorelay
## installation
```shell
$ go get github.com/farseer810/gorelay
```    

## how to run
```shell
$ gorelay [--port <PORT>]
```   

the following code allow us to run this program on port 8000(default to be 8080 if the option is not provided)
```shell
$ gorelay --port 8000
```    

## usage
this program is basically a web proxy program: receives GET and POST requests, 
and resend it to the URI which the Gorelay-Targeturi header indicates.
here's an example:
```python3
target_uri = "http://example.com/targeturl"
params = {} # some GET data
data = {} # some POST data
headers = {'Gorelay-Targeturi': target_uri}

# send a GET request
requests.get('http://localhost:8000', params=params, headers=headers)

# send a POST request
requests.post('http://localhost:8000', data=data, headers=headers)
```
