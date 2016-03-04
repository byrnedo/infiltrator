# Infiltrator

Does simple web/net check of an upstream endpoint through an api.

Use case is when doing remote checks using consul along with a docker 'overlay' network.
If consul is not within docker itself it cant access the overlay. Solution here was to run 
`Infiltrator` inside the overlay network but to listen on the `bridge` interface for api 
calls.


```
Usage of ./infiltrator:
  -host string
        interface to bind to (default "127.0.0.1")
  -port int
        port to bind to (default 8080)
```

## Endpoints

URL: /v1/connect

Does a plain tcp connect attempt. 200 on success.
    
Query params: 

    - host
    - port

URL: /v1/http

Does a http request to the given url. 200 on success, upstream status code in body.
    
Query params: 

    - url
