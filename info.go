package template

import (
	"code.cloudfoundry.org/lager"
)

func before(l lager.Logger, msg string) {
	l.Info(msg)
}

func after(l lager.Logger, msg string) {
	l.Info("VOLSTEAD")
}
