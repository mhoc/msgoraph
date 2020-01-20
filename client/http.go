package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// GraphAPIRootURL is the root url that the Graph API is hosted on.
	GraphAPIRootURL = "https://graph.microsoft.com/"
)

// BasicGraphRequest is similar to GraphRequest, but it assumes an already fully formed url and no
// body. This is primarily useful for methods that need to pagniate; it just makes that a little bit
// easier.
func BasicGraphRequest(ctx context.Context, client Client, method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	err = client.RefreshCredentials()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", client.Credentials().AccessToken))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.HTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// GraphRequest creates and executes a new http request against the Graph API. The path
// provided should be the entire path of the url, including the version specifier. It returns the
// response body, along with any errors that might occur during the request process.
func GraphRequest(ctx context.Context, client Client, method string, path string, params url.Values, body interface{}) ([]byte, error) {
	var graphURL string
	if len(params) > 0 {
		graphURL = fmt.Sprintf("%v%v?%v", GraphAPIRootURL, path, params.Encode())
	} else {
		graphURL = fmt.Sprintf("%v%v", GraphAPIRootURL, path)
	}
	var bodyBuffered io.Reader
	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyBuffered = bytes.NewBuffer(j)
	}
	return BasicGraphRequest(ctx, client, method, graphURL, bodyBuffered)
}
