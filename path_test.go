package parse

import (
	"fmt"
	"testing"
)

func TestPath1(t *testing.T) {
	topic := "bub/1/a/b/c"
	testPathAccept(t, topic, []string{"bub", "c"}, KeepRule, "1", "a", "b", KeepRule)
}

func TestPath2(t *testing.T) {
	topic := "/bub/1/a/b/c"
	testPathAccept(t, topic, []string{"bub", "c"}, KeepRule, "1", "a", "b", KeepRule)
}

func TestPath3(t *testing.T) {
	topic := "bub/1/a/b/c"
	testPathAccept(t, topic, []string{"c"}, SkipRule, "1", "a", "b", KeepRule)
}

func TestPath4(t *testing.T) {
	topic := "bub/1/a/b/c"
	testPathDecline(t, topic, []string{"c"}, "plus", "1", "a", "b", KeepRule)
}

// Helpers

func testPathAccept(t *testing.T, path string, answers []string, rules ...interface{}) {
	ra, err := ExtractSep(path, "/", rules...)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else if arraysMatch(answers, ra) == false {
		t.Fail()
	}
}

func testPathDecline(t *testing.T, path string, answers []string, rules ...interface{}) {
	ra, err := ExtractSep(path, "/", rules...)
	if err == nil && arraysMatch(answers, ra) == true {
		t.Fail()
	}
}

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
