package config

type Logger interface {
	WriteError(error)
	WriteStatus(string)
	WriteDebugStatus(string)
}
