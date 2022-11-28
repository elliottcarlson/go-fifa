package go_fifa_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	fifa "github.com/ImDevinC/go-fifa"
)

type TestClient struct{}

func (c *TestClient) Do(req *http.Request) (*http.Response, error) {
	path := fmt.Sprintf("testdata/%s.json", strings.ReplaceAll(req.URL.Path, "/", "_"))
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	payload, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       io.NopCloser(bytes.NewReader(payload)),
	}
	return resp, nil
}

var _ fifa.HTTPClient = (*TestClient)(nil)
