package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type client interface {
	Request()
	createReqHeader()
}

type HttpGet struct {
	url     string
	query   url.Values
	request *http.Request

	status     string
	statusCode int

	duration time.Duration
}

func (hget *HttpGet) Init(url string, duration time.Duration, query url.Values) {
	hget.url = url
	hget.query = query
	hget.duration = duration
}

func (hget *HttpGet) Url() string {
	return hget.url
}

func (hget *HttpGet) Status() string {
	return hget.status
}

func (hget *HttpGet) StatusCode() int {
	return hget.statusCode
}

func (hget *HttpGet) Duration() time.Duration {
	return hget.duration
}

func (hget *HttpGet) SetUrl(url string) {
	hget.url = url
}

func (hget *HttpGet) SetQuery(query url.Values) {
	hget.query = query
}

func (hget *HttpGet) SetDuration(duration time.Duration) {
	hget.duration = duration
}

func (hget *HttpGet) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
	geturl := ""
	if hget.query.Encode() == "" {
		geturl = hget.url
	} else {
		geturl = hget.url + "?" + hget.query.Encode()
	}
	fmt.Println(geturl)

	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}
	hget.request = req
	hget.createReqHeader(gh, rqh)

	client := new(http.Client)
	client.Timeout = hget.duration

	resp, err := client.Do(hget.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hget.status = resp.Status
	hget.statusCode = resp.StatusCode

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func (hget *HttpGet) createReqHeader(gh GeneralHeader, rqh RequestHeader) {
	hget.request.Header.Set(PRAGMA, gh.Pragma())
	hget.request.Header.Set(CONNECTION, gh.Connection())
	hget.request.Header.Set(CACHE_CONTROL, gh.CacheControl())

	hget.request.Header.Set(ACCEPT, rqh.Accept())
	hget.request.Header.Set(ACCEPT_CHARSET, rqh.AcceptCharset())
	hget.request.Header.Set(FROM, rqh.From())
	hget.request.Header.Set(USER_AGENT, rqh.UserAgent())
	hget.request.Header.Set(REFERER, rqh.Referer())
}
