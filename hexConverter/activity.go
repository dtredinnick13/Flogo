package hexConverter

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-dtredinnick-hexConverter")

const (
	hexBytes = "hexBytes"
	result   = "result"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	hexInput := context.GetInput(hexBytes)

	ivHexInput, ok := hexInput.(string)
	if !ok {
		context.SetOutput("result", 0)
		return true, fmt.Errorf("No hex value provided")
	}

	value, _ := strconv.ParseInt(ivHexInput, 16, 64)

	// Set the output value in the context
	context.SetOutput(result, value)

	return true, nil
}
