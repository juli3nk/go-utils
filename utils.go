package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RecoverFunc() {
	if r := recover(); r != nil {
		log.Println("Recovered:", r)
	}
}

func CreateSlice(input, sep string) []string {
	result := []string{}

	items := strings.Split(input, sep)

	for _, item := range items {
		result = append(result, strings.TrimSpace(item))
	}

	return result
}

func ConvertSliceToMap(sep string, slice []string) map[string]string {
	result := make(map[string]string)

	if len(slice) > 0 {
		for _, s := range slice {
			if !strings.Contains(s, sep) {
				continue
			}

			split := strings.Split(s, sep)

			result[split[0]] = split[1]
		}
	}

	return result
}

func Exit(err error) {
	fmt.Println(err)

	os.Exit(1)
}

func RemoveLastChar(s string) string {
	strLen := len(s) - 1
	newStr := s
	if strLen > 0 {
		newStr = s[0:strLen]
	}

	return newStr
}

func GetReflectValue(k reflect.Kind, i interface{}) reflect.Value {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != k {
		fmt.Printf("%v type can't have attributes inspected\n", val.Kind())
		return reflect.Value{}
	}

	return val
}

// https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func StripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}

	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)

	return str
}

func GetEnv(envKey string) (string, error) {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		return "", fmt.Errorf("environment variable %q not set", envKey)
	}

	return val, nil
}

func GetEnvDefault(envKey, defValue string) string {
	val, err := GetEnv(envKey)
	if err != nil {
		return defValue
	}

	return val
}
