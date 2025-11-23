package request

type Logger interface {
	WriteError(error)
	WriteStatus(string)
	WriteDebugStatus(string)
}
