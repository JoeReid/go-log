package log

import (
	"fmt"
	"log"
	"strings"
)

type stdLogger struct {
	fields []any
}

func (s *stdLogger) WithFields(fields ...any) Logger {
	if len(fields)%2 != 0 {
		panic(fmt.Sprintf("uneven number of fields: %v", fields))
	}

	return &stdLogger{
		fields: append(append(make([]any, 0, len(s.fields)+len(fields)), s.fields), fields...),
	}
}

func (s *stdLogger) print(prefix, msg string) {
	printableFields := make([]string, 0, len(s.fields)/2)
	for i := 0; i < len(s.fields)-1; i += 2 {
		printableFields = append(printableFields, fmt.Sprintf("%v=%v", s.fields[i], s.fields[i+1]))
	}

	log.Printf("%s: %s %s", prefix, msg, strings.Join(printableFields, " "))
}

func (s *stdLogger) Println(msg string, fields ...any) {
	log := s.WithFields(fields...).(*stdLogger)
	log.print("INFO", msg)
}

func (s *stdLogger) Errorln(err error, msg string, fields ...any) {
	log := s.WithFields("error", err).WithFields(fields...).(*stdLogger)
	log.print("ERROR", msg)
}

func NewStdLogger() Logger {
	return &stdLogger{}
}
