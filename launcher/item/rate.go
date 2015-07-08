package item

import (
	. "pkg.deepin.io/dde/daemon/launcher/interfaces"
	. "pkg.deepin.io/dde/daemon/launcher/utils"
)

const (
	_RateRecordFile = "launcher/rate.ini"
	_RateRecordKey  = "rate"
)

func GetFrequencyRecordFile() (RateConfigFileInterface, error) {
	return ConfigFile(_RateRecordFile, "")
}
