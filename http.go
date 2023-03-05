package httpclient

import (
	"io/ioutil"
	"net/http"
)

type client interface {
	Request()
	createReqHeader()
}

type HttpGet struct {
	url     string
	request *http.Request

	Status     string
	StatusCode int
}

func (hget *HttpGet) Url(url string) {
	hget.url = url
}

func (hget *HttpGet) Request(gh GeneralHeader, rqh RequestHeader) ([]byte, error) {
	req, err := http.NewRequest("GET", hget.url, nil)
	if err != nil {
		return nil, err
	}
	hget.request = req

	client := new(http.Client)
	resp, err := client.Do(hget.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hget.Status = resp.Status
	hget.StatusCode = resp.StatusCode

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func (hget *HttpGet) createReqHeader(gh GeneralHeader, rqh RequestHeader) {
	hget.request.Header.Set(PRAGMA, gh.Pragma)
	hget.request.Header.Set(CONNECTION, gh.Connection)
	hget.request.Header.Set(CACHE_CONTROL, gh.CacheControl)

	hget.request.Header.Set(ACCEPT, rqh.Accept)
	hget.request.Header.Set(ACCEPT_CHARSET, rqh.AcceptCharset)
	hget.request.Header.Set(FROM, rqh.From)
	hget.request.Header.Set(USER_AGENT, rqh.UserAgent)
	hget.request.Header.Set(REFERER, rqh.Referer)
}
