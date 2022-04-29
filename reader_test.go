package gosign_test

import (
	"io"
	"os"
	"testing"

	"github.com/xuender/gosign"
)

func TestContainsReader_Contains(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "reader")
	defer file.Close()
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 100))

	file.Seek(0, 0)

	reader := gosign.NewContainsReader(file, []byte{1})
	_, _ = io.ReadAll(reader)

	if reader.Contains() {
		t.Errorf("Contains() return= %v, wantErr %v", false, true)
	}
}
