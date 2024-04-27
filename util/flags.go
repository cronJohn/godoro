package util

import "strings"

type FlagList []string

func (t *FlagList) Set(value string) error {
	*t = strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || r == ' '
	})
	return nil
}

func (t *FlagList) String() string {
	return strings.Join(*t, ",")
}

func (t *FlagList) Type() string {
	return "FlagList"
}
