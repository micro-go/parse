package parse

import (
	"strings"
)

// func ExtractSep() is a wrapper on Extract(), first splitting the string based on the separator.
func ExtractSep(path, separator string, rules ...interface{}) ([]string, error) {
	if separator == "" {
		return nil, badRequestErr
	}
	all := strings.Split(path, separator)
	return Extract(all, rules...)
}

// func Extract() takes a list of strings and returns portions of it, according to the rules.
// Rules can contain either strings or an ExtractRule.
// Strings:
//		The node must match the string exactly. It is not included in the response.
// ExtractRules:
//		KeepRule. Node can be anything. It is returned in the response.
//		SkipRule. Node can be anything. It is not included in the response.
func Extract(src []string, rules ...interface{}) ([]string, error) {
	var resp []string
	srci := 0
	for _, k := range rules {
		// Skip empties in the source
		for srci < len(src) && src[srci] == "" {
			srci++
		}
		if srci >= len(src) {
			return nil, ruleMismatchErr
		}

		switch rule := k.(type) {
		case string:
			if src[srci] != rule {
				return nil, stringMismatchErr
			}
		case ExtractRule:
			switch rule {
			case KeepRule:
				resp = append(resp, src[srci])
			case SkipRule:
			default:
				return nil, invalidExtractRuleErr
			}
		default:
			return nil, invalidRuleErr
		}
		srci++
	}
	return resp, nil
}

// func ExtractOne() extracts a single item. There's no error reporting, if you
// need more complete information use a more complex form.
func ExtractOne(path string, opts ExtractOpts, rules ...interface{}) string {
	if opts.Separator == "" {
		return ""
	}
	all := strings.Split(path, opts.Separator)
	ext, err := Extract(all, rules...)
	if err != nil || len(ext) < 1 {
		return ""
	}
	return ext[0]
}

type ExtractOpts struct {
	Separator string
}

type ExtractRule int

const (
	KeepRule ExtractRule = iota // Keep the item at this location, copying it to the results
	SkipRule
)
