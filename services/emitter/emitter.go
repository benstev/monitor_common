package emitter

import (
	"encoding/json"

	emitter "github.com/emitter-io/go/v2"
	"github.com/rs/zerolog/log"
)

type EmitterService struct {
	host   string
	client *emitter.Client
}

const (
	AGENT_TOPIC = "monitor/agent/"
)

func NewEmitterService(emitterHost string) (*EmitterService, error) {
	e := &EmitterService{host: emitterHost}
	if err := e.connect(); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *EmitterService) connect() error {
	c, err := emitter.Connect(e.host, func(_ *emitter.Client, msg emitter.Message) {
		log.Debug().Str("topic", msg.Topic()).Msg("[emitter] -> received")
	})
	if err != nil {
		return err
	}

	log.Debug().Str("host", e.host).Str("ID", c.ID()).Msg("emitter connected")
	e.client = c
	return nil
}

func (e *EmitterService) Publish(key string, payload interface{}) error {
	d, err := json.Marshal(payload)
	if err != nil {
		return (err)
	}
	if err := e.client.Publish(key, AGENT_TOPIC, d); err != nil {
		return (err)
	}
	return nil
}
