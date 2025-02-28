package debug

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// Output prints debugging messages.
func OutputMarshal(input interface{}, messages ...string) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonBytes, err := json.MarshalIndent(input, "", "  ")
		if err != nil {
			logrus.Warn(err)
			return
		}

		var debugMsg string

		if len(messages) > 0 {
			for _, msg := range messages {
				debugMsg += msg + ":"
			}
		}

		logrus.Debugln(debugMsg + (string(jsonBytes)))
	}
}
