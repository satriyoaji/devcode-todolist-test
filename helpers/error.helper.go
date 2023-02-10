package helpers

import (
	"github.com/sirupsen/logrus"
)

func OutputPanicError(err error) {
	if err != nil {
		logrus.Errorln(err.Error())
		panic(err)
	}
}
