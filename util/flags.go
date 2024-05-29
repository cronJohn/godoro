package util

import (
	"strings"
)

type TagList []string

func (t *TagList) Set(value string) error {
	*t = strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || r == ' '
	})
	return nil
}

func (t *TagList) String() string {
	return strings.Join(*t, ",")
}

func (t *TagList) Type() string {
	return "TagList"
}
