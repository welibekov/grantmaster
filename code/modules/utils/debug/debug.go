package debug

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Output prints debugging messages.
func OutputMarshal(input interface{}, messages ...string) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		jsonBytes, err := yaml.Marshal(input)
		if err != nil {
			logrus.Warn(err)
			return
		}

		var debugMsg string

		if len(messages) > 0 {
			for index, msg := range messages {
				debugMsg += msg + ":"

				if index == len(messages)-1 {
					debugMsg += "\n"
				}
			}
		}

		logrus.Debugf(debugMsg + (string(jsonBytes)))
	}
}
