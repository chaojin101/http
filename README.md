# Extend official go http package

```
go get https://github.com/chaojin101/http
```

## Example

### http.PostFiles

short example:

```go
files := []http.File{
    {
        Fieldname: "file",
        Filename:  "hello.txt",
        Data:      []byte("Hello, World!"),
    },
}
url := "https://httpbin.org/post"
resp, err := http.PostFiles(url, files)
```

short output:

```json
{
  "files": {
    "file": "Hello, World!"
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
	files := []http.File{
		{
			Fieldname: "file",
			Filename:  "hello.txt",
			Data:      []byte("Hello, World!"),
		},
	}
	url := "https://httpbin.org/post"
	resp, err := http.PostFiles(url, files...)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

full output:

```json
{
  "args": {},
  "data": "",
  "files": {
    "file": "Hello, World!"
  },
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "254",
    "Content-Type": "multipart/form-data; boundary=54667f9c32b0610e8fd32ffa4db11783c8a34820c6a2a81cd31e48e643db",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-64bf407e-2a36af6b2026e4ec21284eb6"
  },
  "json": null,
  "origin": "xx.xx.xx.xx",
  "url": "https://httpbin.org/post"
}
```
