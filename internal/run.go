package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type MultiWriter struct {
	writers []io.Writer
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

func run(cmd string, outputerr bool) (string, error) {
	strs := strings.Split(cmd, " ")

	c := exec.Command(strs[0], strs[1:]...)

	var out bytes.Buffer
	stdoutWriter := &MultiWriter{
		writers: []io.Writer{os.Stdout, &out},
	}

	if outputerr {
		stderrWriter := &MultiWriter{
			writers: []io.Writer{os.Stderr, &out},
		}
		c.Stderr = stderrWriter
	}

	c.Stdout = stdoutWriter
	err := c.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err.Error())
	}

	str := out.String()

	return str, err
}
