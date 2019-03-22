package Utils

import (
	"errors"
	"fmt"
	"strings"
)

func ParseHeaderKeyValue(s string) (string, string, error) {
	var (
		key   string
		value string
		er    error
	)
	header := strings.SplitAfterN(s, ":", 2)
	if len(header) != 2 {
		er = errors.New(fmt.Sprintf("Wrong header value. Header: %s\n", s))
	} else {
		key = header[0]
		key = key[:len(key)-1]
		value = header[1]
	}

	return key, value, er
}
