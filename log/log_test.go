package log

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Logger.Info().Msg(string("meg"))
}
