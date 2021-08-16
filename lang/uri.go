package lang

import (
	"net/url"
	"strconv"
)

// URI 提供一个只读的URI接口
type URI interface {
	URL() *url.URL
	String() string

	Scheme() string
	User() string // username and password information
	Host() string // host or host:port
	Port() int
	Path() string     // path (relative paths may omit leading slash)
	Query() string    // encoded query values, without '?'
	Fragment() string // fragment for references, without '#'
}

// CreateURI 创建URI
func CreateURI(url *url.URL) URI {

	uri := &innerURI{}
	port, err := strconv.Atoi(url.Port())
	host := url.Hostname()

	uri.inner = *url
	uri.host = host
	if err == nil {
		uri.port = port
	}

	return uri
}

// ParseURI 解析URI
func ParseURI(str string) (URI, error) {
	u1, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	return CreateURI(u1), nil
}

////////////////////////////////////////////////////////////////////////////////

type innerURI struct {
	inner url.URL
	host  string
	port  int
}

func (inst *innerURI) _Impl() URI {
	return inst
}

func (inst *innerURI) Scheme() string {
	return inst.inner.Scheme
}

func (inst *innerURI) User() string {
	return inst.inner.User.String()
}

func (inst *innerURI) Host() string {
	return inst.host
}

func (inst *innerURI) Port() int {
	return inst.port
}

func (inst *innerURI) Path() string {
	return inst.inner.Path
}

func (inst *innerURI) Query() string {
	return inst.inner.RawQuery
}

func (inst *innerURI) Fragment() string {
	return inst.inner.Fragment
}

func (inst *innerURI) URL() *url.URL {
	result := &url.URL{}
	*result = inst.inner
	return result
}

func (inst *innerURI) String() string {
	return inst.inner.String()
}
