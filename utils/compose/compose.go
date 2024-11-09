package compose

import (
	"encoding/json"
	"fmt"
)

func ComposeHttpResponse(version, status, content_type, body string) string {
	return fmt.Sprintf("%s %s\r\nContent-Type: %s\r\n\r\n%s", version, status, content_type, body)
}

func ComposeString_toJson(body map[string]string) string {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
