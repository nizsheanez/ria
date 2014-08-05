package messages

import (
	"errors"
	"encoding/json"
)

const (
	MSG_HELLO        = 1
	MSG_WELCOME      = 2
	MSG_ABORT        = 3
	MSG_CHALLENGE    = 4
	MSG_AUTHENTICATE = 5
	MSG_GOODBYE      = 6
	MSG_HEARTBEAT    = 7
	MSG_ERROR        = 8

	MSG_PUBLISH   = 16
	MSG_PUBLISHED = 17

	MSG_SUBSCRIBE    = 32
	MSG_SUBSCRIBED   = 33
	MSG_UNSUBSCRIBE  = 34
	MSG_UNSUBSCRIBED = 35
	MSG_EVENT        = 36

	MSG_CALL   = 48
	MSG_CANCEL = 49
	MSG_RESULT = 50

	MSG_REGISTER     = 64
	MSG_REGISTERED   = 65
	MSG_UNREGISTER   = 66
	MSG_UNREGISTERED = 67
	MSG_INVOCATION   = 68
	MSG_INTERRUPT    = 69
	MSG_YIELD        = 70
)

var ErrWrongMessageType = errors.New("Wrong message type")

type Message interface {
	Marshal() ([]byte, error)
	Array() (interface{})
	Unmarshal([]byte) error
}

type MessageBase struct{}

func (this *MessageBase) Array() []byte {
	panic("Don't use me!")
}

func (this *MessageBase) Marshal() ([]byte, error) {
	str, err := json.Marshal(this.Array())
	if err != nil {
		panic(err)
	}
	return str, nil
}


type Details struct {
	Roles    *Roles `json:"roles"`
}

type Roles map[string]*Role

type Role struct {
	Features *Features `json:"features"`
}

type Features map[string]bool

