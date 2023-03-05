package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpGet struct {
	url     string
	query   url.Values
	request *http.Request

	status     string
	statusCode int

	duration time.Duration
}

type HttpPost struct {
	url     string
	body    url.Values
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

func (hpst *HttpPost) Init(url string, duration time.Duration, body url.Values) {
	hpst.url = url
	hpst.body = body
	hpst.duration = duration
}

func (hget *HttpGet) Url() string {
	return hget.url
}

func (hpst *HttpPost) Url() string {
	return hpst.url
}

func (hget *HttpGet) Status() string {
	return hget.status
}

func (hpst *HttpPost) Status() string {
	return hpst.status
}

func (hget *HttpGet) StatusCode() int {
	return hget.statusCode
}

func (hpst *HttpPost) StatusCode() int {
	return hpst.statusCode
}

func (hget *HttpGet) Duration() time.Duration {
	return hget.duration
}

func (hpst *HttpPost) Duration() time.Duration {
	return hpst.duration
}

func (hget *HttpGet) SetUrl(url string) {
	hget.url = url
}

func (hpst *HttpPost) SetUrl(url string) {
	hpst.url = url
}

func (hget *HttpGet) SetQuery(query url.Values) {
	hget.query = query
}

func (hpst *HttpPost) SetBody(body url.Values) {
	hpst.body = body
}

func (hget *HttpGet) SetDuration(duration time.Duration) {
	hget.duration = duration
}

func (hpst *HttpPost) SetDuration(duration time.Duration) {
	hpst.duration = duration
}

func (hget *HttpGet) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
	geturl := ""
	if hget.query.Encode() == "" {
		geturl = hget.url
	} else {
		geturl = hget.url + "?" + hget.query.Encode()
	}
	u, err := url.Parse(geturl)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "http" {
		return nil, errors.New("Scheme Format error")
	}
	geturl = u.Scheme + "://" + u.Host + u.Path + "?" + u.RawQuery

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
	hget.request.Header.Set(CONTENT_TYPE, rqh.ContentType())
}

func (hpst *HttpPost) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
	u, err := url.Parse(hpst.url)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "http" {
		return nil, errors.New("Scheme Format error")
	}
	posturl := u.Scheme + "://" + u.Host + u.Path

	req, err := http.NewRequest("POST", posturl, strings.NewReader(hpst.body.Encode()))
	if err != nil {
		return nil, err
	}
	hpst.request = req
	hpst.createReqHeader(gh, rqh)

	client := new(http.Client)
	client.Timeout = hpst.duration

	resp, err := client.Do(hpst.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hpst.status = resp.Status
	hpst.statusCode = resp.StatusCode

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func (hpst *HttpPost) createReqHeader(gh GeneralHeader, rqh RequestHeader) {
	hpst.request.Header.Set(PRAGMA, gh.Pragma())
	hpst.request.Header.Set(CONNECTION, gh.Connection())
	hpst.request.Header.Set(CACHE_CONTROL, gh.CacheControl())

	hpst.request.Header.Set(ACCEPT, rqh.Accept())
	hpst.request.Header.Set(ACCEPT_CHARSET, rqh.AcceptCharset())
	hpst.request.Header.Set(FROM, rqh.From())
	hpst.request.Header.Set(USER_AGENT, rqh.UserAgent())
	hpst.request.Header.Set(REFERER, rqh.Referer())
	hpst.request.Header.Set(CONTENT_TYPE, rqh.ContentType())
}
