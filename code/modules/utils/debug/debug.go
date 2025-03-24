package debug

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// OutputMarshal prints debug messages and marshals the input into YAML format.
// It accepts an input value to be marshaled and optional messages to accompany the debug output.
func OutputMarshal(input interface{}, messages ...string) {
	// Check if debug level logging is enabled.
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		// Marshal the input into YAML format.
		jsonBytes, err := yaml.Marshal(input)
		if err != nil {
			// Log a warning if marshaling fails.
			logrus.Warn(err)
			return
		}

		var debugMsg string

		// Create a debug message string if messages are provided.
		if len(messages) > 0 {
			for index, msg := range messages {
				debugMsg += msg + ":"

				// Add a newline character at the end of the last message.
				if index == len(messages)-1 {
					debugMsg += "\n"
				}
			}
		}

		// Log the debug message with the marshaled input.
		logrus.Debugf(debugMsg + (string(jsonBytes)))
	}
}
