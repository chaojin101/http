package http

import (
	"fmt"
	"io"
	"os"
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

func TestPostMultipartWithFile(t *testing.T) {
	filename := "http.go"
	imgData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fields := []MultipartField{
		{
			Fieldname: "img",
			File: File{
				Name:        filename,
				ContentType: "application/octet-stream",
			},
			Data: imgData,
		},
		{
			Fieldname: "host",
			Data:      []byte("www.websiteplanet.com"),
		},
		{
			Fieldname: "token",
			Data:      []byte("03AAYGu2SIbdonwYLMdoDpa2s8tSS-C6IIlxyaOwD98QG9m-NuKr0Lcf_MBDUC7FzJR-VC_KH1UtoX-KPeOM_LpM5zlN8j94_cw7ISczEwC_gEL57a-pLWhGEjjArktZ5qJB657oRed9-Q4M4FJHQyyfO9YU2LaarPzUJQK3Hb1jhXZ4D0Hv2UCrPJ5fQDGNAnVZjgJW0vx29FpHmzzXRElOBOML5txieEAfS1sOUystYsXO8pKpi76MU1yNoIVu8IWyhRpoaiFOUx5ygUtP-2w1oT56PCgO7d2t95C_1z4tEICiGqb7NZC1neFO4tOT9lpZwRh39crGJklI82X9gxOmOL91Isfl4Wq5-6IrENjYKaMAGaLpwDWGAkSTIxiooIxOUFrRjCC51cjcThmS5Z6v3UHelX3_iEXyAOw3aO_gCTV4ceNxdS4JusQOFZYaK_5EA9vwUVt6Iw-f7BKOt_nUXYIcsqMtWidTbGWd-byaXvoaMywoaQNPb6XWFrEexT2duEFqinz-EjOd2weFFT7stn91Pmr3JKPb7XKVk2KkaGR1Rudc78uQRZtxa7zb4m3QL8TDlRi5CJ"),
		},
		{
			Fieldname: "mode",
			Data:      []byte("high"),
		}, {
			Fieldname: "path",
			Data:      []byte("16902643787131583908"),
		}, {
			Fieldname: "user_key",
			Data:      []byte(""),
		},
	}
	url := "https://httpbin.org/post"
	resp, err := PostMultipart(url, fields...)
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
