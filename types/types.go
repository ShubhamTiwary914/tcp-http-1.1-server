package types

type RequestInfo struct {
	Method  string
	Path    string
	Version string
	Query   map[string]string
}

type Headers map[string]string

type Body map[string]string
