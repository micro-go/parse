package parse

import (
	"fmt"
	"testing"
)

func TestAsArguments1(t *testing.T) {
	cla := "parse it all"
	testArgAccept(t, cla, []string{"parse", "it", "all"})
}

func TestAsArguments2(t *testing.T) {
	cla := `parse "it" all`
	testArgAccept(t, cla, []string{"parse", "it", "all"})
}

func TestAsArguments3(t *testing.T) {
	cla := `parse "yo momma" all`
	testArgAccept(t, cla, []string{"parse", "yo momma", "all"})
}

func TestAsArguments4(t *testing.T) {
	cla := `parse "yo \"momma\"" all`
	testArgAccept(t, cla, []string{"parse", `yo "momma"`, "all"})
}

// Helpers

func testArgAccept(t *testing.T, src string, a []string) {
	b := AsArguments(src)
	if arraysMatch(a, b) == false {
		fmt.Println("No match a", a, "b", b)
		t.Fail()
	}
}

func testArgsDecline(t *testing.T, src string, a []string) {
	b := AsArguments(src)
	if arraysMatch(a, b) == true {
		t.Fail()
	}
}
