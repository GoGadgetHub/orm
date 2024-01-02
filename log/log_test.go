package log

import (
	"os"
	"testing"
)

func TestSetLevel2(t *testing.T) {
	SetLevel(InfoLevel)
	if infoLog.Writer() != os.Stdout && errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}
	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}
