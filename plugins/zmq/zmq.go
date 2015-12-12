package zmq

import (
	"github.com/influxdb/telegraf/plugins"
	"github.com/zeromq/goczmq"
)

func init() {
	plugins.Add("zmq", func() plugins.Plugin {
		return &Zmq{}
	})
}

type Zmq struct {
}

func (zmq *Zmq) Description() string {
	return ""
}

func (zmq *Zmq) SampleConfig() string {
	sampleConf := `
[socket]

    [socket.PULL]
    host = "127.0.0.1"
    port = "5555"

    [socket.SUB]
    host  = "127.0.0.1"
    port  = "5556"
    topic = "metric"
`
	return sampleConf
}

func (zmq *Zmq) Gather(acc plugins.Accumulator) error {
	return nil
}
