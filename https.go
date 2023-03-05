package httpclient

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpsGet struct {
	url     string
	query   url.Values
	request *http.Request

	status     string
	statusCode int

	duration time.Duration
}

type HttpsPost struct {
	url     string
	body    url.Values
	request *http.Request

	status     string
	statusCode int

	duration time.Duration
}

func (hget *HttpsGet) Init(url string, duration time.Duration, query url.Values) {
	hget.url = url
	hget.query = query
	hget.duration = duration
}

func (hpst *HttpsPost) Init(url string, duration time.Duration, body url.Values) {
	hpst.url = url
	hpst.body = body
	hpst.duration = duration
}

func (hget *HttpsGet) Url() string {
	return hget.url
}

func (hpst *HttpsPost) Url() string {
	return hpst.url
}

func (hget *HttpsGet) Status() string {
	return hget.status
}

func (hpst *HttpsPost) Status() string {
	return hpst.status
}

func (hget *HttpsGet) StatusCode() int {
	return hget.statusCode
}

func (hpst *HttpsPost) StatusCode() int {
	return hpst.statusCode
}

func (hget *HttpsGet) Duration() time.Duration {
	return hget.duration
}

func (hpst *HttpsPost) Duration() time.Duration {
	return hpst.duration
}

func (hget *HttpsGet) SetUrl(url string) {
	hget.url = url
}

func (hpst *HttpsPost) SetUrl(url string) {
	hpst.url = url
}

func (hget *HttpsGet) SetQuery(query url.Values) {
	hget.query = query
}

func (hpst *HttpsPost) SetBody(body url.Values) {
	hpst.body = body
}

func (hget *HttpsGet) SetDuration(duration time.Duration) {
	hget.duration = duration
}

func (hpst *HttpsPost) SetDuration(duration time.Duration) {
	hpst.duration = duration
}

func (hget *HttpsGet) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
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
	if u.Scheme != "https" {
		return nil, errors.New("Scheme Format error")
	}
	geturl = u.Scheme + "://" + u.Host + u.Path

	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return nil, err
	}
	hget.request = req
	hget.createReqHeader(gh, rqh)

	client := new(http.Client)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: u.Hostname()},
	}
	client.Timeout = hget.duration
	client.Transport = tr

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

func (hget *HttpsGet) createReqHeader(gh GeneralHeader, rqh RequestHeader) {
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

func (hpst *HttpsPost) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
	u, err := url.Parse(hpst.url)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: u.Hostname()},
	}
	client.Timeout = hpst.duration
	client.Transport = tr

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

func (hpst *HttpsPost) createReqHeader(gh GeneralHeader, rqh RequestHeader) {
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
