package webstate

// Variable that holds our test result
var testResult string

// Init ...
func Init() {
	testResult = "Yay!"
}

// GetTestResult ...
func GetTestResult() string {
	return testResult
}
