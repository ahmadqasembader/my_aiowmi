package main


// TODO
type Protocol struct {
	
}

type Connection struct {
	host     string
	port     int
	username string
	password string
	// protocol *Protocol
}

func (c *Connection) NewConn(host, username, password string, port int) *Connection {
	c.host = host
	c.username = username
	c.password = password
	c.port = port
	return c
}
