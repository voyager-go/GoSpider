package tools

import (
	"fmt"
	"go.spider/global"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Url string
	Method string
	Headers *http.Header
	Body io.Reader
	Handle Handle
	Client http.Client
}

func (r *Request) Execute() error{
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return err
	}
	req.Header = *r.Headers
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error status code: %d", resp.StatusCode)
	}
	r.Handle.Worker(resp.Body, r.Url)
	defer resp.Body.Close()
	return nil
}


func (r *Request) NewRequest(method, requestUrl, userAgent string, handle Handle, body io.Reader) (*Request, error) {
	if _, err := url.Parse(requestUrl); err != nil{
		return nil, err
	}
	hdr := http.Header{}
	if userAgent != "" {
		hdr.Add("User-Agent", userAgent)
	}else {
		hdr.Add("User-Agent", global.UserAgent)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
	return &Request{
		Url:     requestUrl,
		Method:  method,
		Headers: &hdr,
		Body:    body,
		Handle:  handle,
		Client:  client,
	}, nil
}
