package main

import (
	"fmt"
	"io"
	"os"
)

type Copper struct {
	Writer io.Writer
}

func main() {
	c := &Copper{os.Stdout}
	fmt.Fprintln(c, "Hello World")
}

func (c *Copper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}

		out[i] = c
	}

	return c.Writer.Write(out)
}
