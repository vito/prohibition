package template

import (
	"code.cloudfoundry.org/lager"
)

func before(l lager.Logger, msg string, data lager.Data) {
	l.Info(msg, data)
}

func after(l lager.Logger, msg string, data lager.Data) {
	l.Info("VOLSTEAD")
}
