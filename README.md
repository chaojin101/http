# Extend official go http package

## Usage

```
go get https://github.com/chaojin101/http
```

### http.PostFiles

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
