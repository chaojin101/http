package http

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

type MultipartFile struct {
	Fieldname string
	Filename  string
	Data      []byte
}

func PostMultipart(url string, files ...MultipartFile) (*http.Response, error) {
	reqBodyBuf := new(bytes.Buffer)
	multipart := multipart.NewWriter(reqBodyBuf)

	for _, file := range files {
		formFile, err := multipart.CreateFormFile(file.Fieldname, file.Filename)
		if err != nil {
			return nil, err
		}

		_, err = formFile.Write(file.Data)
		if err != nil {
			return nil, err
		}
	}

	err := multipart.Close()
	if err != nil {
		return nil, err
	}

	return http.Post(url, multipart.FormDataContentType(), reqBodyBuf)
}
