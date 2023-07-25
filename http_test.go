package http

import (
	"io"
	"testing"
)

func TestPostMultipart(t *testing.T) {
	fields := []MultipartField{
		{
			Fieldname: "img",
			Data:      []byte("123"),
		},
	}
	url := "https://httpbin.org/post"
	resp, err := PostMultipart(url, fields...)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
			return
		}
	}()

	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(respBodyData))
}
