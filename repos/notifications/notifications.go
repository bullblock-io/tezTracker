package notifications

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Repo interface {
	GetBlockUpdate() (lvl uint64, err error)
}

const (
	chanBlockInserts = "block_insert"
)

func New(pgconn string) *Notification {
	cb := func(event pq.ListenerEventType, err error) {
		switch event {
		case pq.ListenerEventConnected:
			fmt.Printf("connected\n")
		case pq.ListenerEventConnectionAttemptFailed:
			fmt.Printf("connection failed, err: %s\n", err)
		case pq.ListenerEventDisconnected:
			fmt.Printf("disconnected\n")
		case pq.ListenerEventReconnected:
			fmt.Printf("reconnected\n")
		default:
			fmt.Printf("unknown event type %d, err: %s\n", event, err)
		}
	}
	listener := pq.NewListener(pgconn, time.Second, time.Second, cb)
	n := &Notification{listener: listener}
	n.subscribe()
	return n
}

type Notification struct {
	listener *pq.Listener
	err      error
}
type blockLvl struct {
	Level uint64 `json:"lvl"`
}

func (n *Notification) subscribe() {
	n.err = n.listener.Listen(chanBlockInserts)
}

func (n *Notification) GetBlockUpdate() (lvl uint64, err error) {
	if n.err != nil {
		return 0, n.err
	}
	for {
		select {
		case msg := <-n.listener.Notify:
			var data blockLvl
			if err := json.Unmarshal([]byte(msg.Extra), &data); err != nil {
				return 0, err
			}
			return data.Level, nil
		case <-time.After(time.Second * 20):
			if err := n.listener.Ping(); err != nil {
				return 0, err
			}
		}
	}
}
