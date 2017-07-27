package webstate

var testResult string

// Init ...
func Init() {
	testResult = "Yay!"
}

// GetTestResult ...
func GetTestResult() string {
	return testResult
}
