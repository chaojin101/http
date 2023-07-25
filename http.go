package http

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

type File struct {
	Name        string
	ContentType string // optional: default is application/octet-stream
}

type MultipartField struct {
	Fieldname string
	File      File // optional
	Data      []byte
}

// PostMultipart posts files or data to url.
// The request header Content-Type is set to multipart/form-data with random boundary.
func PostMultipart(url string, multipartFields ...MultipartField) (*http.Response, error) {
	reqBodyBuf := new(bytes.Buffer)
	form := multipart.NewWriter(reqBodyBuf)

	for _, multipartField := range multipartFields {
		var formPart io.Writer
		var err error
		if multipartField.File.Name == "" {
			formPart, err = form.CreateFormField(multipartField.Fieldname)
			if err != nil {
				return nil, err
			}
		} else {
			if multipartField.File.ContentType == "" {
				formPart, err = form.CreateFormFile(multipartField.Fieldname, multipartField.File.Name)
				if err != nil {
					return nil, err
				}
			} else {
				formPart, err = createFormFileWithContentType(form, multipartField.Fieldname, multipartField.File.Name, multipartField.File.ContentType)
				if err != nil {
					return nil, err
				}
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

func createFormFileWithContentType(w *multipart.Writer, fieldname, filename, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			fieldname, filename))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}
