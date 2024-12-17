package connect

import (
	"fmt"
	"sync"
)

type Connector struct {
	attempt int
}

func (c *Connector) Connect() {
	fmt.Println("connected attempt", c.attempt)
}

func NewConnector(attempt int) *Connector {
	return &Connector{attempt: attempt}
}

var conn *Connector
var once sync.Once

func GetConnector(attempt int) *Connector {
	once.Do(func() {
		conn = NewConnector(attempt)
	})

	return conn
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetConnector(i).Connect()
		}()
	}

	wg.Wait()
	fmt.Println("conn attempt", conn.attempt)
}
