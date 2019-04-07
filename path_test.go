package parse

import (
	"fmt"
	"testing"
)

func TestPathExtractSepAccept(t *testing.T) {
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

func TestPathExtractSepDecline(t *testing.T) {
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

func TestPathExtractOne(t *testing.T) {
	cases := []struct {
		Path  string
		Rules []interface{}
		Want  string
	}{
		{"a", []interface{}{KeepRule}, "a"},
		{"a", []interface{}{SkipRule}, ""},
		{"a/b", []interface{}{KeepRule, KeepRule}, "a"},
		{"a/b", []interface{}{SkipRule, KeepRule}, "b"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			opts := ExtractOpts{Separator: "/"}
			have := ExtractOne(tc.Path, opts, tc.Rules...)
			if have != tc.Want {
				fmt.Println("Mismtach have", have, "want", tc.Want)
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
