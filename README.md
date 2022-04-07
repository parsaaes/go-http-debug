# go-http-debug
Useful utilities to be used in http clients as a debugging tool.
## Examples
### Transport With Dump
This is an http.RoundTripper implementation that wraps another http.RoundTripper to add dump ability using `net/http/httputil` package.
To use this, simply use it as a transport in the http client. You can specify the internal transport to be used. Otherwise, it will use the http.DefaultTransport.
```go
package main

import (
	"bytes"
	"fmt"
	ghd "github.com/parsaaes/go-http-debug"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: &ghd.TransportWithDump{
			Handler: func(req string, resp string) {
				fmt.Println(req)
				fmt.Println(resp)
			},
		},
	}

	body := []byte(`{"abc": 123}`)

	req, _ := http.NewRequest(http.MethodPost, "https://example.com", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")

	client.Do(req)
}
```

Output:

```
POST / HTTP/1.1
Host: example.com
User-Agent: Go-http-client/1.1
Content-Length: 12
Content-Type: application/json
Accept-Encoding: gzip

{"abc": 123}
HTTP/2.0 200 OK
Accept-Ranges: bytes
Cache-Control: max-age=604800
Content-Type: text/html; charset=UTF-8
Date: Thu, 07 Apr 2022 14:03:27 GMT
Etag: "3147526947+gzip"
Expires: Thu, 14 Apr 2022 14:03:27 GMT
Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT
Server: EOS (vny/044E)
Vary: Accept-Encoding

<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
```

### Transport With Curl
Sometimes it's easier to communicate with other people using curl commands to talk about HTTP requests. This is an http.RoundTripper implementation that wraps another http.RoundTripper and returns a curl translation of the request using [http2curl](https://github.com/moul/http2curl).
To use this, simply use it as a transport in the http client. You can specify the internal transport to be used. Otherwise, it will use the http.DefaultTransport.
```go
package main

import (
	"bytes"
	"fmt"
	ghd "github.com/parsaaes/go-http-debug"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: &ghd.TransportWithCurl{
			Handler: func(curl string) {
				fmt.Println(curl)
			},
		},
	}

	body := []byte(`{"abc": 123}`)

	req, _ := http.NewRequest(http.MethodPost, "https://example.com", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")

	client.Do(req)
}
```
Output:
```
curl -X 'POST' -d '{"abc": 123}' -H 'Content-Type: application/json' 'https://example.com'
```