package webstate

// Import the testing package
import "testing"

func TestStuff(t *testing.T) {
	if len(testResult) != 0 {
		t.Fail()
	}
	if len(GetTestResult()) != 0 {
		t.Fail()
	}
	Init()
	if len(testResult) == 0 {
		t.Fail()
	}
	if len(GetTestResult()) == 0 {
		t.Fail()
	}
	// t.Fail()
	// t.Error("I'm in a bad mood.")
}