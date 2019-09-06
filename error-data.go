package template

import (
	"code.cloudfoundry.org/lager"
	"github.com/sirupsen/logrus"
)

func before(l lager.Logger, msg string, err error, data lager.Data) {
	l.Error(msg, err, data)
}

func after(l lager.Logger, msg string, err error, data lager.Data) {
	logrus.WithFields(logrus.Fields(data)).WithError(err).Error("%%" + msg + "%%")
}
