package redis

import (
	"log"
	"sync"

	"github.com/desertbit/glue"
	"github.com/garyburd/redigo/redis"
	"strings"
)

type Redis struct {
	sockets map[*glue.Socket]map[string]bool
	topics map[string]map[*glue.Socket]bool

	pubconn redis.Conn
	subconn redis.PubSubConn

	l sync.RWMutex
}

func (r *Redis) Init(url string) error {
	c, err := redis.DialURL(url)
	if err != nil {
		return err
	}
	r.pubconn = c

	c, err = redis.DialURL(url)
	if err != nil {
		return err
	}
	r.subconn = redis.PubSubConn{c}

	go func() {
		for {
			switch v := r.subconn.Receive().(type) {
			case redis.Message:
				r.EmitLocal(v.Channel, string(v.Data))

			case error:
				panic(v)
			}
		}
	}()

	return nil
}

func (r *Redis) Subscribe(s *glue.Socket, t string) error {
	r.l.Lock()
	defer r.l.Unlock()

	_, ok := r.sockets[s]
	if !ok {
		r.sockets[s] = map[string]bool{}
	}
	r.sockets[s][t] = true

	_, ok = r.topics[t]
	if !ok {
		r.topics[t] = map[*glue.Socket]bool{}
		err := r.subconn.Subscribe(t)
		if err != nil {
			return err
		}
	}
	r.topics[t][s] = true

	return nil
}

func (r *Redis) UnsubscribeAll(s *glue.Socket) error {
	r.l.Lock()
	defer r.l.Unlock()

	for t := range r.sockets[s] {
		delete(r.topics[t], s)
		if len(r.topics[t]) == 0 {
			delete(r.topics, t)
			err := r.subconn.Unsubscribe(t)
			if err != nil {
				return err
			}
		}
	}
	delete(r.sockets, s)

	return nil
}

func (r *Redis) Emit(t string, m string) error {
	_, err := r.pubconn.Do("PUBLISH", t, m)
	return err
}

func (r *Redis) EmitLocal(t string, m string) {
	r.l.RLock()
	defer r.l.RUnlock()
	for s := range r.topics[t] {
		s.Write(m)
	}
}

func (r *Redis) HandleSocket(s *glue.Socket) {
	s.OnClose(func() {
		err := r.UnsubscribeAll(s)
		if err != nil {
			log.Print(err)
		}
	})

	s.OnRead(func(data string) {
		fields := strings.Fields(data)
		if len(fields) == 0 {
			return
		}
		if fields[0] == "sub" {
			if len(fields) != 2 {
				return
			}
			err := r.Subscribe(s, fields[1])
			if err != nil {
				log.Print(err)
			}
		}
	})
}
