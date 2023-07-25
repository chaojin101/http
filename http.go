package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

type MultipartField struct {
	Fieldname string
	Filename  string // optional
	Data      []byte
}

// PostMultipart posts files or data to url.
// The Content-Type header is set to multipart/form-data with random boundary.
func PostMultipart(url string, multipartFields ...MultipartField) (*http.Response, error) {
	reqBodyBuf := new(bytes.Buffer)
	form := multipart.NewWriter(reqBodyBuf)

	for _, multipartField := range multipartFields {
		var formPart io.Writer
		var err error
		if multipartField.Filename == "" {
			formPart, err = form.CreateFormField(multipartField.Fieldname)
			if err != nil {
				return nil, err
			}
		} else {
			formPart, err = form.CreateFormFile(multipartField.Fieldname, multipartField.Filename)
			if err != nil {
				return nil, err
			}
		}
		_, err = formPart.Write(multipartField.Data)
		if err != nil {
			return nil, err
		}
	}

	err := form.Close()
	if err != nil {
		return nil, err
	}

	return http.Post(url, form.FormDataContentType(), reqBodyBuf)
}
