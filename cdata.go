package mxj

import (
	"fmt"
	"strconv"
)

var xmlCdata bool

func XMLCdata(b bool) {
	xmlCdata = b
}

func makeCdata(value interface{}) string {
	result := ""
	switch value.(type) {
	case string:
		result = fmt.Sprintf("%s %s %s", "<![CDATA[", value.(string), "]]>")
	case float64:
		result = fmt.Sprintf("%s %s %s", "<![CDATA[", strconv.FormatFloat(value.(float64), 'f', -1, 64), "]]>")
	default:
		// noop
	}
	return result
}
