package parse

import (
	"testing"
)

// No wildcards

func TestMatch1(t *testing.T) {
	testMqttAccept(t, "a", "a")
}

func TestMatch2(t *testing.T) {
	testMqttAccept(t, "a/b", "a/b")
}

func TestMatch3(t *testing.T) {
	testMqttAccept(t, "a/b/c", "a/b/c")
}

func TestMatch4(t *testing.T) {
	testMqttDecline(t, "a", "b")
}

func TestMatch5(t *testing.T) {
	testMqttDecline(t, "a/a", "b")
}

func TestMatch6(t *testing.T) {
	testMqttDecline(t, "a/a", "a")
}

func TestMatch7(t *testing.T) {
	testMqttDecline(t, "a", "a/b")
}

// Single-level wildcards

func TestMatch8(t *testing.T) {
	testMqttAccept(t, "a/+", "a/a")
}

func TestMatch9(t *testing.T) {
	testMqttAccept(t, "a/+", "a/b")
}

func TestMatch10(t *testing.T) {
	testMqttDecline(t, "a/+", "a/b/c")
}

func TestMatch11(t *testing.T) {
	testMqttAccept(t, "a/+/c", "a/a/c")
}

func TestMatch12(t *testing.T) {
	testMqttDecline(t, "a/+/c", "a/a")
}

func TestMatch13(t *testing.T) {
	testMqttDecline(t, "a/+/c", "a/a/d")
}

// Multi-level wildcards

func TestMatch14(t *testing.T) {
	testMqttDecline(t, "a/#/c", "a/a/d")
}

func TestMatch15(t *testing.T) {
	testMqttDecline(t, "a/#/c", "a/a/b/d")
}

func TestMatch16(t *testing.T) {
	testMqttDecline(t, "a/#/c", "a/a/c/d")
}

func TestMatch17(t *testing.T) {
	testMqttAccept(t, "a/#", "a/a/d")
}

func TestMatch18(t *testing.T) {
	testMqttAccept(t, "a/#/d", "a/a/d")
}

func TestMatch19(t *testing.T) {
	testMqttAccept(t, "a/#/d", "a/a/b/c/d")
}

// Helpers

func testMqttAccept(t *testing.T, pattern, cmp string) {
	m := NewMqttStringMatch(pattern)
	if !m.Matches(cmp) {
		t.Fail()
	}
}

func testMqttDecline(t *testing.T, pattern, cmp string) {
	m := NewMqttStringMatch(pattern)
	if m.Matches(cmp) {
		t.Fail()
	}
}
