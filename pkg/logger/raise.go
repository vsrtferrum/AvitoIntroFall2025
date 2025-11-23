package logger

func (log *Logger) Raise() error {
	logTemp, err := log.cfg.Build()
	if err != nil {
		return err
	}

	log.log = logTemp
	return nil
}
