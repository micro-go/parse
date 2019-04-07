package parse

import (
	"fmt"
	"testing"
)

func TestPathAccept(t *testing.T) {
	cases := []struct {
		Path  string
		Rules []interface{}
		Want  []string
	}{
		{"a/b", []interface{}{KeepRule, KeepRule}, []string{"a", "b"}},
		{"a/b/c", []interface{}{KeepRule, KeepRule}, []string{"a", "b"}},
		{"a/b/c", []interface{}{KeepRule, KeepRule, SkipRule}, []string{"a", "b"}},
		{"a/b/c", []interface{}{KeepRule, SkipRule, KeepRule}, []string{"a", "c"}},
		{"bub/1/a/b/c", []interface{}{KeepRule, "1", "a", "b", KeepRule}, []string{"bub", "c"}},
		{"/bub/1/a/b/c", []interface{}{KeepRule, "1", "a", "b", KeepRule}, []string{"bub", "c"}},
		{"bub/1/a/b/c", []interface{}{SkipRule, "1", "a", "b", KeepRule}, []string{"c"}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			have, err := ExtractSep(tc.Path, "/", tc.Rules...)
			if err != nil {
				fmt.Println(err)
				t.Fail()
			} else if arraysMatch(tc.Want, have) == false {
				t.Fail()
			}
		})
	}
}

func TestPathDecline(t *testing.T) {
	cases := []struct {
		Path  string
		Rules []interface{}
		Want  []string
	}{
		{"bub/1/a/b/c", []interface{}{"plus", "1", "a", "b", KeepRule}, []string{"c"}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			have, err := ExtractSep(tc.Path, "/", tc.Rules...)
			if err == nil && arraysMatch(tc.Want, have) == true {
				t.Fail()
			}
		})
	}
}

// Helpers

func arraysMatch(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
