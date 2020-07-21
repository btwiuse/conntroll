package agent

import (
	"net"

	"github.com/ginuerzh/gost"
	types "k0s.io/k0s/pkg/agent"
)

func StartRedirectServer(c types.Config) types.RedirectServer {
	redirectListener := NewLys()
	go redirectServe(redirectListener)
	return redirectListener
}

var redirectHandler = gost.TCPRedirectHandler()

func redirectServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer c.Close()
			redirectHandler.Handle(c)
		}()
	}
}