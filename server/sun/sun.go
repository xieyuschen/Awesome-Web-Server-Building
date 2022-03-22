package sun

import (
	"log"
	"net"
	"os"
)

func (e *Engine) Run(address string) {
	var err error
	e.listener, err = net.Listen("tcp", address)
	if err != nil {
		e.logger.Fatalln("Create listener failed, %w", err)
	}
	for {
		conn, err := e.listener.Accept()
		if err != nil {
			e.logger.Println("Error is: %s", err)
		} else {
			e.logger.Println("Connection received, remote address is:", conn.RemoteAddr())
		}
		go func(conn net.Conn) {
			// conn.SetWriteDeadline(time.Unix(1, 0))
			n, _ := conn.Write([]byte("hello world"))
			e.logger.Printf("Write successfully, %d byte\n", n)
		}(conn)
	}
}
func Default() Engine {
	return Engine{logger: log.New(os.Stdout, "log:", log.Lshortfile)}
}
