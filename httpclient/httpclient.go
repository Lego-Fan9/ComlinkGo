package httpclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

var ErrMaxRetriesExceeded = errors.New("max retries exceeded")
var ErrUnknownHTTP = errors.New("unknown error making http.Do")

var Client *HTTPClient

const (
	maxRetries     = 5
	retryDelay     = 2 * time.Second
	maxIdleConns   = 50
	fiveSeconds    = 5 * time.Second  //nolint:revive,stylecheck
	fifteenSeconds = 15 * time.Second //nolint:revive,stylecheck
	thirtySeconds  = 30 * time.Second //nolint:revive,stylecheck
	ninetySeconds  = 90 * time.Second //nolint:revive,stylecheck
)

type HTTPClient struct {
	Client *http.Client
	Ctx    context.Context
	Wg     *sync.WaitGroup
}

func Init(ctx context.Context, wg *sync.WaitGroup) *HTTPClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   fiveSeconds,
			KeepAlive: thirtySeconds,
		}).DialContext,
		MaxIdleConns:        maxIdleConns,
		IdleConnTimeout:     ninetySeconds,
		TLSHandshakeTimeout: fifteenSeconds,
	}

	Client = &HTTPClient{
		Client: &http.Client{
			Timeout:   thirtySeconds,
			Transport: transport,
		},
		Ctx: ctx,
		Wg:  wg,
	}

	return Client
}

func (c *HTTPClient) DoWithRetry(req *http.Request) (*http.Response, error) {
	var err error

	reqRoot, _, err := cloneRequest(req)
	if err != nil {
		return nil, err
	}

	for attempt := range maxRetries {
		var reqTemp *http.Request

		reqTemp, reqRoot, err = cloneRequest(reqRoot)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrUnknownHTTP, err)
		}

		resp, errr := c.DoWithoutRetry(reqTemp)
		if errr == nil {
			return resp, nil
		}

		err = errr

		time.Sleep(retryDelay << attempt)
	}

	if err == nil {
		err = ErrMaxRetriesExceeded
	}

	return nil, err
}

func cloneRequest(req *http.Request) (*http.Request, *http.Request, error) {
	var bodyBytes []byte

	var err error

	if req.Body != nil {
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, nil, err //nolint:wrapcheck
		}

		req.Body.Close()
	}

	makeClone := func() *http.Request {
		c := req.Clone(req.Context())
		if bodyBytes != nil {
			c.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		return c
	}

	return makeClone(), makeClone(), nil
}

func (c *HTTPClient) DoWithoutRetry(req *http.Request) (*http.Response, error) {
	c.Wg.Add(1)
	defer c.Wg.Done()

	req = req.WithContext(c.Ctx)

	return c.Client.Do(req)
}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	return c.DoWithRetry(req)
}

func (c *HTTPClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Content-Type", contentType)

	return c.DoWithRetry(req)
}
