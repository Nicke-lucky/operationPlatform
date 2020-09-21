package utils

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestStrTimeTotime(t *testing.T) {
	StrTimeTotime("1212-12-12 12:12:12")
}

func TestStrTimeToNowtime(t *testing.T) {
	StrTimeToNowtime()
}

//KuaizhaoTimeNowFormat
func TestKuaizhaoTimeNowFormat(t *testing.T) {
	logrus.Print(KuaizhaoTimeNowFormat())
}

//Yesterdaydate()
func TestYesterdaydate(t *testing.T) {
	logrus.Print(Yesterdaydate())
}

//OldData
func TestOldData(t *testing.T) {
	logrus.Print(OldData(7))
}

//StrdateToNowdate()
func TestDateToNowdate(t *testing.T) {
	logrus.Println(DateToNowdate())
}

// DateFormatTimeToTime
func TestDateFormatTimeToTime(t *testing.T) {
	logrus.Println(DateFormatTimeTostrdate(time.Now()))
}
