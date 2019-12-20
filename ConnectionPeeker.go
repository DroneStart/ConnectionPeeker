package connectionPeeker

import (
	"bufio"
	"bytes"
	"net"
)

//ConnectionPeeker Wraps net.Conn enabling peeking into requests without consuming
type ConnectionPeeker struct {
	r *bufio.Reader
	net.Conn
}

//NewConnectionPeeker creates a new connection peeker
func NewConnectionPeeker(c net.Conn) ConnectionPeeker {
	return ConnectionPeeker{bufio.NewReader(c), c}
}


//Peek peeks (without consuming) into the connection
func (b ConnectionPeeker) Peek(n int) ([]byte, error) {
	return b.r.Peek(n)
}

//PeekLine peeks (without consuming) one line - starting at startAt, incrementing by incremental bytes, with maxLineLen as maximum length of a line
func (b ConnectionPeeker) PeekLine(startAt int, incremental int, maxLineLen int) []byte {
	for i := startAt; i < maxLineLen; i = i + incremental {
		by, _ := b.r.Peek(i)
		loc := bytes.Index(by, []byte("\r"))
		if loc > 0 {
			return by[0:loc]
		}
	}
	return nil
}

//Read reads (and consumes) from the connection
func (b ConnectionPeeker) Read(p []byte) (int, error) {
	return b.r.Read(p)
}


//IsValid returns whether the connection is valid
func (b ConnectionPeeker) IsValid() bool {
	return (b.r != nil)
}
