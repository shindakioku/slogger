package adapters

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestStdoutAdapter_Execute(t *testing.T) {
	stdoutAdapter := StdoutAdapter{}

	cases := []struct {
		name    string
		message []byte
		assert  func(out []byte) bool
	}{
		{
			name:    "Write hello world!",
			message: []byte("hello world!"),
			assert: func(out []byte) bool {
				return assert.Equal(t, []byte("hello world!"), out)
			},
		},
	}

	for _, c := range cases {
		f := func() {
			stdoutAdapter.Execute(c.message)
		}

		if !assert.True(t, c.assert(captureOutput(f))) {
			t.Error(c.name + " don't passed")
		}
	}
}

func captureOutput(f func()) []byte {
	fname := filepath.Join(os.TempDir(), "stdout")
	old := os.Stdout
	temp, _ := os.Create(fname)
	os.Stdout = temp

	f()

	temp.Close()
	os.Stdout = old

	out, _ := ioutil.ReadFile(fname)
	outString := strings.ReplaceAll(strings.ReplaceAll(string(out), "[", ""), "]", "")
	var bytes []byte
	for _, v := range strings.Split(strings.Replace(outString, "\n", "", len(outString) - 1), " ") {
		i, _ := strconv.Atoi(v)
		bytes = append(bytes, byte(i))
	}

	return bytes
}
