package utils

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNumberToChinese(t *testing.T) {
	logrus.Print(NumberToChinese(3091746200))
}
