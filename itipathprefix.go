package itipathprefix

import (
	"net/http"
	"strings"
)

// PathPrefixMatcher Store the prefix to match with the request Path
type PathPrefixMatcher struct {
	prefix string
}

// New is the constructor to ItiPathPrefix
func New(template string) *PathPrefixMatcher {
	return &PathPrefixMatcher{
		prefix: template,
	}
}

// Match returns if the request can be handled by this Route.
func (t *PathPrefixMatcher) Match(req *http.Request) bool {

	if req.URL.Path == t.prefix {
		return true
	}

	req.URL.Path = strings.Trim(req.URL.Path, "/")
	t.prefix = strings.Trim(t.prefix, "/")

	pText := strings.Split(req.URL.Path, "/")
	pPrefix := strings.Split(t.prefix, "/")

	for i := range pPrefix {
		if pPrefix[i] == "*" {
			continue
		}

		if pPrefix[i] != pText[i] {
			if strings.HasPrefix(pText[i], pPrefix[i]) {
				return true
			}
			return false
		}
	}

	return true
}
