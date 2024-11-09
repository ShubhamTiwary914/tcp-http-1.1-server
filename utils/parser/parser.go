package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	HTTPSchema "server/types"
)

func ParseHttpRequest(rawRequest string) (HTTPSchema.Headers, HTTPSchema.Body, HTTPSchema.RequestInfo) {
	// Split on double CRLF (Unix mostly)
	parts := strings.Split(rawRequest, "\r\n\r\n")
	headers, info := ParseHeadersInfo(parts[0])
	body := ParseBody(parts[1])
	return headers, body, info
}

func ParseHeadersInfo(headers string) (HTTPSchema.Headers, HTTPSchema.RequestInfo) {
	//split string into lines
	var reqHeaders HTTPSchema.Headers = make(HTTPSchema.Headers)
	var info HTTPSchema.RequestInfo
	lines := strings.Split(headers, "\r\n")

	// First line contains request info
	if len(lines) > 0 {
		parts := strings.Split(lines[0], " ")
		if len(parts) >= 3 {
			info.Method = parts[0]
			info.Path = parts[1]
			info.Version = parts[2]
		}
	}

	// Remaining lines are headers
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			reqHeaders[key] = value
		}
	}
	return reqHeaders, info
}

func ParseBody(body string) HTTPSchema.Body {
	// Clean the input
	body = strings.TrimSpace(body)
	body = strings.Trim(body, "\x00") // Remove null characters (!issue with json parser)
	body = strings.ReplaceAll(body, "'", "\"")

	// Now decode the proper JSON
	result := make(map[string]string)
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println("Error ", err)
		return nil
	}
	return result
}
