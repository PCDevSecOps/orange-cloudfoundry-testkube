package testkube

func TestSuiteExecutionStatusPtr(status TestSuiteExecutionStatus) *TestSuiteExecutionStatus {
	return &status
}

var TestSuiteExecutionStatusFailed = TestSuiteExecutionStatusPtr(FAILED_TestSuiteExecutionStatus)
var TestSuiteExecutionStatusPassed = TestSuiteExecutionStatusPtr(PASSED_TestSuiteExecutionStatus)
var TestSuiteExecutionStatusQueued = TestSuiteExecutionStatusPtr(QUEUED_TestSuiteExecutionStatus)
var TestSuiteExecutionStatusRunning = TestSuiteExecutionStatusPtr(RUNNING_TestSuiteExecutionStatus)