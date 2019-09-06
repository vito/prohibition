package template

import (
	"code.cloudfoundry.org/lager"
)

func before(l lager.Logger, msg string) {
	l.Debug(msg)
}

func after(l lager.Logger, msg string) {
	l.Debug("VOLSTEAD")
}
