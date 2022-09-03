package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createClientHTTP(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/testredirect", nil)
	ra := []*http.Request{r}
	tests := []struct {
		name string
		want bool
	}{
		{"Just client", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createClientHTTP()
			err := got.CheckRedirect(r, ra)
			assert.Equal(t, tt.want, (err != nil))
		})
	}
}

func Test_prepareResult(t *testing.T) {
	ts1 := map[int]resultRedirectCheck{
		0: {URL: "http://testurl.com", StatusCode: 301, Status: "Moved Permanently"},
		1: {URL: "https://testurl.com", StatusCode: 200, Status: "OK"},
	}
	ts1w := resultRedirectsCheck{
		Redirects: ts1,
	}
	ts2 := map[int]resultRedirectCheck{
		0: {URL: "http://testurl.com", StatusCode: 500, Status: "Error", Error: "Server Error"},
	}
	ts2w := resultRedirectsCheck{
		Redirects: ts2,
		Error:     ts2[0].Error,
	}

	tests := []struct {
		name   string
		result map[int]resultRedirectCheck
		want   resultRedirectsCheck
	}{
		{"Ok map", ts1, ts1w},
		{"Error 500 map", ts2, ts2w},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prepareResult(tt.result)
			assert.Equal(t, tt.want, got)
		})
	}
}
