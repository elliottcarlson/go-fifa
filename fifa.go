package go_fifa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

///////////////////////////////////

///////////////////////////////////

const (
	defaultAPIBaseURL = "https://api.fifa.com/api/v3"
	defaultUserAgent  = "go-fifa"
	defaultLanguage   = "en-US,en"
)

type Client struct {
	Client     HTTPClient
	ApiBaseURL string
	UserAgent  string
	Language   string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *Client) get(path string, respData interface{}, reqData interface{}) (interface{}, error) {
	return c.sendRequest(http.MethodGet, path, respData, reqData)
}

func (c *Client) sendRequest(method string, path string, respData interface{}, reqData interface{}) (interface{}, error) {
	req, err := c.newRequest(method, path, reqData)
	if err != nil {
		return nil, err
	}
	err = c.doRequest(req, respData)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *Client) newRequest(method string, path string, data interface{}) (*http.Request, error) {
	if c.ApiBaseURL == "" {
		c.ApiBaseURL = defaultAPIBaseURL
	}
	url := c.ApiBaseURL + path
	return c.newStandardRequest(url, method, data)
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
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	response, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute API request: %s", err.Error())
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("invalId response code: %s", response.Status)
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return fmt.Errorf("failed to decode API response: %s", err.Error())
	}
	return nil
}

func (c *Client) setRequestHeaders(req *http.Request) {
	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if c.Language == "" {
		c.Language = defaultLanguage
	}
	if c.Language != "" {
		req.Header.Add("Accept-Language", c.Language)
	}
}
