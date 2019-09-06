package template

import (
	"code.cloudfoundry.org/lager"
	"github.com/sirupsen/logrus"
)

func before(l lager.Logger, msg string, err error) {
	l.Error(msg, err)
}

func after(l lager.Logger, msg string, err error) {
	logrus.WithError(err).Error("PATTEN" + msg + "UNPATTEN")
}
