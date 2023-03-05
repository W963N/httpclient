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
	CONTENT_TYPE   = "Content-Type"
)

const (
	UTF8                    = "utf-8"
	MIME_TYPE_TEXT          = "text"
	MIME_TYPE_APPLICATION   = "application"
	MIME_SUBTYPE_HTML       = "html"
	MIME_SUBTYPE_JSON       = "json"
	MIME_SUBTYPE_CSV        = "csv"
	MIME_SUBTYPE_PLAIN      = "plain"
	MIME_SUBTYPE_URLENCODED = "x-www-form-urlencoded"
	MIME_SUBTYPE_JS         = "javascript"
)

var user_agent = "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

type RequestHeader struct {
	accept        string
	acceptCharset string
	from          string
	referer       string
	userAgent     string
	authorization string
	contentType   string
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

func (rqh *RequestHeader) ContentType() string {
	return rqh.contentType
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

func (rqh *RequestHeader) SetContentType(contenttype string) {
	rqh.contentType = contenttype
}

func (rqh *RequestHeader) SetAuthorization(authorization string) {
	rqh.authorization = authorization
}

func (rqh *RequestHeader) Init() {
	mt := []MimeType{
		{mType: MIME_TYPE_TEXT, mSubtype: MIME_SUBTYPE_PLAIN},
		{mType: MIME_TYPE_APPLICATION, mSubtype: MIME_SUBTYPE_JSON},
		{mType: MIME_TYPE_TEXT, mSubtype: MIME_SUBTYPE_HTML},
		{mType: MIME_TYPE_TEXT, mSubtype: MIME_SUBTYPE_CSV},
		{mType: MIME_TYPE_TEXT, mSubtype: MIME_SUBTYPE_JS},
	}
	rqh.accept = rqh.createMimetype(mt)
	rqh.acceptCharset = UTF8
	rqh.from = ""
	rqh.referer = ""
	rqh.contentType = rqh.createMimetype([]MimeType{
		{mType: MIME_TYPE_APPLICATION, mSubtype: MIME_SUBTYPE_URLENCODED},
	})
	rqh.userAgent = user_agent
}

type MimeType struct {
	mType    string
	mSubtype string
}

func (m *MimeType) Type() string {
	return m.mType
}

func (m *MimeType) Subtype() string {
	return m.mSubtype
}

func (m *MimeType) SetType(mType string) {
	m.mType = mType
}

func (m *MimeType) SetSubtype(mSubtype string) {
	m.mSubtype = mSubtype
}

func (rqh *RequestHeader) createMimetype(mimetype []MimeType) string {
	mt := []string{}

	for _, m := range mimetype {
		lo := []string{m.mType, m.mSubtype}
		mt = append(mt, strings.Join(lo, "/"))
	}
	return strings.Join(mt, ",")
}
