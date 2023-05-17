package log

func Errorf(format string, v ...any) {
	PrintRed(format, v...)
}
