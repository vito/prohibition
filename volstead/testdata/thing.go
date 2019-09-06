package testdata

import "code.cloudfoundry.org/lager"

// Comment
func Boring(logger lager.Logger, other int) (string, error) {
	return "", nil
}

// other comment
func NotLogger(other int) (string, error) {
	return "", nil
}

// other comment
type Typeeee struct{}

func (t *Typeeee) Method(logger lager.Logger, other int) (string, error) {
	// other comment
	return "", nil
}

func (t *Typeeee) NotLoggerMethod(logger int) (string, error) {
	return "", nil
}
