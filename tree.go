package parse

// Operations on an arbitrary interface{}, typically read from JSON or XML.
// The basic pattern is to refer to keys by using "/" to denote parent/child.

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func ReadJsonFile(filename string) (interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tree interface{}
	err = json.Unmarshal(data, &tree)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

func TreeBool(path string, root interface{}, defaultValue bool) bool {
	_v, err := TreeValue(path, root)
	if err != nil {
		return defaultValue
	}
	switch v := _v.(type) {
	case bool:
		return v
	case int:
		return v > 0
	case string:
		return len(v) > 0 && v[0] == 'T' || v[0] == 't'
	}
	return defaultValue
}

func TreeString(path string, root interface{}, defaultValue string) string {
	_v, err := TreeValue(path, root)
	if err != nil {
		return defaultValue
	}
	if v, ok := _v.(string); ok {
		return v
	}
	return defaultValue
}

func TreeValue(_path string, _root interface{}) (interface{}, error) {
	root, ok := _root.(map[string]interface{})
	if !ok {
		return nil, badRequestErr
	}
	path := strings.Split(_path, "/")
	return findTreeValue(path, root)
}

func findTreeValue(path []string, root map[string]interface{}) (interface{}, error) {
	if len(path) < 1 {
		return "", badRequestErr
	}
	v := root[path[0]]
	if v == nil {
		return nil, noValueErr
	}
	if len(path) == 1 {
		return v, nil
	}
	if root, ok := v.(map[string]interface{}); ok {
		return findTreeValue(path[1:], root)
	}
	return nil, noMapErr
}
