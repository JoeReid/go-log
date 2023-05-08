package log

type Logger interface {
	WithFields(fields ...any) Logger
	Println(msg string, fields ...any)
	Errorln(err error, msg string, fields ...any)
}
