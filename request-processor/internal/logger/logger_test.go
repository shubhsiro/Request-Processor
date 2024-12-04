package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	logFilePath := "../../logs/app.log"
	InitLogger(logFilePath)

	_, err := os.Stat(logFilePath)
	assert.False(t, os.IsNotExist(err))

	os.Remove(logFilePath)
}

func TestLog(t *testing.T) {
	logFilePath := "../../logs/test_app.log"
	InitLogger(logFilePath)

	Log("This is a test log message")

	file, err := os.Open(logFilePath)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	buf := make([]byte, 1024)
	n, _ := file.Read(buf)

	assert.Contains(t, string(buf[:n]), "This is a test log message")

	os.Remove(logFilePath)
}
