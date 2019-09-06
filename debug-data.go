package template

import (
	"code.cloudfoundry.org/lager"
)

func before(l lager.Logger, msg string, data lager.Data) {
	l.Debug(msg, data)
}

func after(l lager.Logger, msg string, data lager.Data) {
	l.Debug("VOLSTEAD")
}
