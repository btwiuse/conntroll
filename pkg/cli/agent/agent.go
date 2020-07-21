package agent

import (
	"log"
	"time"

	"k0s.io/k0s/pkg/agent/agent"
	"k0s.io/k0s/pkg/agent/config"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	ag := agent.NewAgent(c)

	for range time.Tick(time.Second) {
		err = ag.ConnectAndServe()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}