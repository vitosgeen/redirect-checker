package service

import (
	"fmt"
	"net/http"
	"strings"
)

type resultRedirectsCheck struct {
	Redirects map[int]resultRedirectCheck
	Error     string
	Text      string
}

type resultRedirectCheck struct {
	URL        string            `json:"url"`
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Error      string            `json:"error"`
	Text       string            `json:"text"`
}

const CountIterations = 24

func MakeRequest(checkURL string) resultRedirectsCheck {
	var result = make(map[int]resultRedirectCheck)
	nextURL := checkURL
	var lastStatusCode int
	// client := createClientHTTP()
	i := 0
	for i < CountIterations {
		client := createClientHTTP()
		itemResult := checkRedirectIteration(&nextURL, &client)
		result[i] = itemResult
		lastStatusCode = itemResult.StatusCode
		if itemResult.StatusCode == http.StatusOK {
			itemResult.Text = "Done!"
			break
		}
		i++
	}
	if lastStatusCode != http.StatusOK {
		var itemResult resultRedirectCheck
		itemResult.Error = "Somthing went wrong!"
		itemResult.Headers = make(map[string]string)
		result[i+1] = itemResult
	}
	if i == CountIterations {
		var itemResult resultRedirectCheck
		itemResult.Error = "Too many redirects"
		itemResult.Headers = make(map[string]string)
		result[i+1] = itemResult
	}

	return prepareResult(result)
}

func prepareResult(result map[int]resultRedirectCheck) resultRedirectsCheck {
	var generalResult resultRedirectsCheck
	generalResult.Redirects = result
	lastIndex := len(result) - 1
	if len(result[lastIndex].Error) > 0 {
		generalResult.Error = result[lastIndex].Error
	}
	generalResult.Text = result[lastIndex].Text
	return generalResult
}
func checkRedirectIteration(nextURL *string, client *http.Client) resultRedirectCheck {
	var url string
	var itemResult resultRedirectCheck

	url = *nextURL
	resp, err := client.Get(url)
	if err != nil {
		itemResult.Error = err.Error()
		fmt.Println(err)
	}

	itemResult.Headers = collectorHeaders(resp)
	itemResult.Status = resp.Status
	itemResult.StatusCode = resp.StatusCode
	itemResult.URL = resp.Request.URL.String()

	if resp.StatusCode == http.StatusOK {
		itemResult.Text = "Done!"
	} else {
		url = resp.Header.Get("Location")
		*nextURL = url
	}
	return itemResult
}

func collectorHeaders(resp *http.Response) map[string]string {
	headers := make(map[string]string)
	for k, h := range resp.Header {
		var hs = strings.Join(h, ", ")
		headers[k] = hs
	}
	return headers
}
func createClientHTTP() http.Client {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	return client
}
