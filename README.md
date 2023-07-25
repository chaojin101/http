# Extend official go http package

```
go get https://github.com/chaojin101/http
```

## Example

### http.PostFiles

short example:

```go
fields := []http.MultipartField{
  {
    Fieldname: "img",
    File: http.File{
      Name: "1.jpg",
    },
    Data: []byte("hello world"),
  },
}
url := "https://httpbin.org/post"
resp, err := http.PostMultipart(url, fields...)
```

short output:

```json
{
  "files": {
    "img": "hello world"
  },
  "headers": {
    "Content-Type": "multipart/form-data; boundary=54667f9c32b0610e8fd32ffa4db11783c8a34820c6a2a81cd31e48e643db"
  }
}
```

full runnable example:

```go
package main

import (
	"fmt"
	"io"

	"github.com/chaojin101/http"
)

func main() {
	fields := []http.MultipartField{
		{
			Fieldname: "img",
			File: http.File{
				Name: "1.jpg",
			},
			Data: []byte("hello world"),
		},
	}
	url := "https://httpbin.org/post"
	resp, err := http.PostMultipart(url, fields...)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBodyData))
}
```

full output:

```json
{
  "args": {},
  "data": "",
  "files": {
    "img": "hello world"
  },
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "247",
    "Content-Type": "multipart/form-data; boundary=1805df65e71312c6bb872b2a04a4695877a3b5c5b6963d4348d4fe517b0e",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-64bf7a18-6b7158c06fe32e8e6b13da72"
  },
  "json": null,
  "origin": "xxx.xxx.xxx.xxx",
  "url": "https://httpbin.org/post"
}
```
