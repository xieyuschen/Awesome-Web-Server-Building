package test

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func dialAndRead() error {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	b := make([]byte, 20)
	n, er := conn.Read(b)
	fmt.Println(string(b), len(b), n, er)
	return nil
}
func TestEstablishSocket(t *testing.T) {
	if err := dialAndRead(); err != nil {
		t.Fail()
	}
}

func Test100Connections(t *testing.T) {
	m := sync.WaitGroup{}
	m.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer m.Done()
			err := dialAndRead()
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
			}
		}()
	}
	m.Wait()
}
