package lib

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type ArgsFormatter struct {
	KeyPrefix              string
	KeyPostfix             string
	DisableSplitMultiwords bool
}

func (kf ArgsFormatter) formatValue(v reflect.Value) string {
	return fmt.Sprintf("%v", v.Interface())
}

func (kf ArgsFormatter) formatKey(key string) string {
	var buf bytes.Buffer

	if kf.DisableSplitMultiwords {
		buf.WriteString(key)
	} else {
		for index, r := range key {
			if unicode.IsUpper(r) && index != 0 {
				buf.WriteByte('-')
			}
			buf.WriteRune(r)
		}
	}

	key = strings.ToLower(buf.String())
	key = kf.KeyPrefix + key + kf.KeyPostfix
	return key
}

func (kf ArgsFormatter) FormatArgs(md interface{}) []string {
	var args []string

	vType := reflect.TypeOf(md)
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	val := reflect.ValueOf(md)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	ielements := vType.NumField()

	for i := 0; i < ielements; i++ {
		t := vType.Field(i)
		v := val.Field(i)

		// skip zero values
		if v.Interface() == reflect.Zero(t.Type).Interface() {
			continue
		}
		tagOpts := strings.Split(t.Tag.Get("arg"), ",")
		name := tagOpts[0]
		if name == "" {
			name = t.Name
		}
		if name != "-" {
			args = append(args, kf.formatKey(name))
		}

		if len(tagOpts) > 0 {
			tagOpts = tagOpts[1:]
		}
		if !stringInSlice("omitvalue", tagOpts) {
			args = append(args, kf.formatValue(v))
		}
	}
	return args
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
