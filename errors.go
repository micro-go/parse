package parse

import (
	"errors"
)

var (
	badRequestErr              = errors.New("Bad request")
	emptyErr              = errors.New("Token empty")
	invalidExtractRuleErr = errors.New("Invalid ExtractRule")
	invalidRuleErr        = errors.New("Invalid rule")
	ruleMismatchErr       = errors.New("Rule mismatch")
	stringMismatchErr     = errors.New("String mismatch")
)
