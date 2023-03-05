package main

import "strings"

const (
	PRAGMA        = "Pragma"
	CACHE_CONTROL = "Cache-Control"
	CONNECTION    = "Connection"
)

const (
	NO_STORE = "no-store"
	NO_CACHE = "no-cache"
	CLOSE    = "close"
)

type generalHeader struct {
	Pragma       string
	CacheControl string
	Connection   string
}

func (gh *generalHeader) New() {
	gh.Pragma = NO_STORE
	gh.CacheControl = NO_STORE
	gh.Connection = CLOSE
}

const (
	ACCEPT         = "Accept"
	ACCEPT_CHARSET = "Accept-Charset"
	FROM           = "From"
	REFERER        = "Referer"
	USER_AGENT     = "User-Agent"
	AUTHORIZATION  = "Authorization"
)

const (
	UTF8                  = "utf-8"
	MIME_TYPE_TEXT        = "text"
	MIME_TYPE_APPLICATION = "application"
	MIME_SUBTYPE_HTML     = "html"
	MIME_SUBTYPE_JSON     = "json"
	MIME_SUBTYPE_CSV      = "csv"
	MIME_SUBTYPE_PLAIN    = "plain"
)

var user_agent = "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

type requestHeader struct {
	Accept        string
	AcceptCharset string
	From          string
	Referer       string
	UserAgent     string
	Authorization string
}

func (rqh *requestHeader) New() {
	mt := []mimeType{
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_PLAIN},
		{Type: MIME_TYPE_APPLICATION, Subtype: MIME_SUBTYPE_JSON},
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_HTML},
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_CSV},
	}
	rqh.Accept = rqh.CreateAccept(mt)
	rqh.AcceptCharset = UTF8
	rqh.From = ""
	rqh.Referer = ""
	rqh.UserAgent = user_agent
}

type mimeType struct {
	Type    string
	Subtype string
}

func (rqh *requestHeader) CreateAccept(mimetype []mimeType) string {
	mt := []string{}

	for _, m := range mimetype {
		lo := []string{m.Type, m.Subtype}
		mt = append(mt, strings.Join(lo, "/"))
	}
	return strings.Join(mt, ",")
}
