package logger

func (log *Logger) WriteError(err error) {
	log.log.Error(err.Error() + "\n")
}

func (log *Logger) WriteDebugStatus(msg string) {
	log.log.Debug(msg + "\n")
}

func (log *Logger) WriteStatus(msg string) {
	log.log.Info(msg + "\n")
}
