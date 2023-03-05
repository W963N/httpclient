package httpclient

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

type GeneralHeader struct {
	pragma       string
	cacheControl string
	connection   string
}

func (gh *GeneralHeader) Pragma() string {
	return gh.pragma
}

func (gh *GeneralHeader) CacheControl() string {
	return gh.cacheControl
}

func (gh *GeneralHeader) Connection() string {
	return gh.connection
}

func (gh *GeneralHeader) SetPragma(pragma string) {
	gh.pragma = pragma
}

func (gh *GeneralHeader) SetCacheControl(cachecontrol string) {
	gh.cacheControl = cachecontrol
}

func (gh *GeneralHeader) SetConnection(connection string) {
	gh.connection = connection
}

func (gh *GeneralHeader) Init() {
	gh.pragma = NO_STORE
	gh.cacheControl = NO_STORE
	gh.connection = CLOSE
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

type RequestHeader struct {
	accept        string
	acceptCharset string
	from          string
	referer       string
	userAgent     string
	authorization string
}

func (rqh *RequestHeader) Accept() string {
	return rqh.accept
}

func (rqh *RequestHeader) AcceptCharset() string {
	return rqh.acceptCharset
}

func (rqh *RequestHeader) From() string {
	return rqh.from
}

func (rqh *RequestHeader) Referer() string {
	return rqh.referer
}

func (rqh *RequestHeader) UserAgent() string {
	return rqh.userAgent
}

func (rqh *RequestHeader) Authorization() string {
	return rqh.authorization
}

func (rqh *RequestHeader) SetAccept(accept string) {
	rqh.accept = accept
}

func (rqh *RequestHeader) SetAcceptCharset(acceptcharset string) {
	rqh.acceptCharset = acceptcharset
}

func (rqh *RequestHeader) SetFrom(from string) {
	rqh.from = from
}

func (rqh *RequestHeader) SetReferer(referer string) {
	rqh.referer = referer
}

func (rqh *RequestHeader) SetUserAgent(useragent string) {
	rqh.userAgent = useragent
}

func (rqh *RequestHeader) SetAuthorization(authorization string) {
	rqh.authorization = authorization
}

func (rqh *RequestHeader) Init() {
	mt := []mimeType{
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_PLAIN},
		{Type: MIME_TYPE_APPLICATION, Subtype: MIME_SUBTYPE_JSON},
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_HTML},
		{Type: MIME_TYPE_TEXT, Subtype: MIME_SUBTYPE_CSV},
	}
	rqh.accept = rqh.CreateAccept(mt)
	rqh.acceptCharset = UTF8
	rqh.from = ""
	rqh.referer = ""
	rqh.userAgent = user_agent
}

type mimeType struct {
	Type    string
	Subtype string
}

func (rqh *RequestHeader) CreateAccept(mimetype []mimeType) string {
	mt := []string{}

	for _, m := range mimetype {
		lo := []string{m.Type, m.Subtype}
		mt = append(mt, strings.Join(lo, "/"))
	}
	return strings.Join(mt, ",")
}
