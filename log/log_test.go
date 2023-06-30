package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Logger.Info().Msg(string("msg"))
}

func TestSetLogDir(t *testing.T) {
	SetLogDir("./runtime/")
}
