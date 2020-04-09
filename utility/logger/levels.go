package logger




func Debugf(formatString string, args...interface{}) {
	LOGGER.Infof(formatString, args...)
}

func Infof(formatString string, args...interface{}) {
	LOGGER.Infof(formatString, args...)
}

func Warnf(formatString string, args...interface{}) {
	LOGGER.Warningf(formatString, args...)
}

func Errorf(formatString string, args...interface{}) {
	LOGGER.Infof(formatString, args...)
}

func Fatalf(formatString string, args...interface{}) {
	LOGGER.Infof(formatString, args...)
}