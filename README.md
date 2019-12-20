# ConnectionPeeker
Wrap a `net.Conn` to peek into streams without consuming

`Peek()` and `PeekLine()` allows peeking into the `net.Conn` without consuming data from the connection, while `Read()` consumes data from the connection.
