package fastchamp

import (
	"net"
)

// --- server ---

const (
	defaultBufferSize = 8192
)

type Server struct {
	BufferSize int
}

func (s Server) ListenAndServe(addr string) error {
	ln, err := net.Listen("tcp4", addr)
	defer ln.Close()
	if err != nil {
		return err
	}

	return s.serve(ln)
}

func (s Server) serve(listener net.Listener) error {
	c, err := listener.Accept()
	if err != nil {
		return err
	}

	var buff = make([]byte, defaultBufferSize)
	for {
		_, err := c.Read(buff[:])
		if err != nil {
			break
		}
		var bBuff = make([]byte, 2*len(buff))
		//bBuff = append(bBuff, buff)

		copy(bBuff, buff)

		tmp := buff

		buff = make([]byte, 2*len(buff))
	}
	return nil
}

// --- --- ---
