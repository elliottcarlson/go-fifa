package go_fifa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

///////////////////////////////////

///////////////////////////////////

const (
	defaultAPIBaseURL = "https://api.fifa.com/api/v1"
	defaultUserAgent  = "go-fifa"
	defaultLanguage   = "en-US,en"
)

type Client struct {
	opts *Options
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Options struct {
	HTTPClient HTTPClient
	APIBaseURL string
	UserAgent  string
	Language   string
}

func NewClient(options *Options) (*Client, error) {
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}

	if options.APIBaseURL == "" {
		options.APIBaseURL = defaultAPIBaseURL
	}

	if options.UserAgent == "" {
		options.UserAgent = defaultUserAgent
	}

	if options.Language == "" {
		options.Language = defaultLanguage
	}

	client := &Client{
		opts: options,
	}

	return client, nil
}

func (c *Client) get(path string, respData interface{}, reqData interface{}) (interface{}, error) {
	return c.sendRequest(http.MethodGet, path, respData, reqData, false)
}

// func (c *Client) post(path string, respData interface{}, reqData interface{}, hasJSONBody bool) (interface{}, error) {
// 	return c.sendRequest(http.MethodPost, path, respData, reqData, hasJSONBody)
// }

func (c *Client) sendRequest(method string, path string, respData interface{}, reqData interface{}, hasJSONBody bool) (interface{}, error) {
	req, err := c.newRequest(method, path, reqData, hasJSONBody)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(req, respData)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *Client) newRequest(method string, path string, data interface{}, hasJSONBody bool) (*http.Request, error) {
	url := c.opts.APIBaseURL + path
	if hasJSONBody {
		return c.newJSONRequest(url, method, data)
	}
	return c.newStandardRequest(url, method, data)
}

func (c *Client) newJSONRequest(url string, method string, data interface{}) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) newStandardRequest(url string, method string, data interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if data != nil {
		v, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = v.Encode()
	}
	return req, nil
}

func (c *Client) doRequest(req *http.Request, resp interface{}) error {
	c.setRequestHeaders(req)
	response, err := c.opts.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute API request: %s", err.Error())
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("invalid response code: %s", response.Status)
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %s", err.Error())
	}
	return nil
}

func (c *Client) setRequestHeaders(req *http.Request) {
	if c.opts.UserAgent != "" {
		req.Header.Add("User-Agent", c.opts.UserAgent)
	}
	if c.opts.Language != "" {
		req.Header.Add("Accept-Language", c.opts.Language)
	}
}
