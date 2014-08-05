package messages

import (
	"encoding/json"
)

type HelloMessage struct{
	MessageBase
	Realm        string
	Details      Details
}

func (this *HelloMessage) Array() interface{} {
	arr := []interface{}{
		MSG_HELLO,
		this.Realm,
		this.Details,
	}
	return arr
}

func (this *HelloMessage) Unmarshal(str []byte) (err error) {
	var messageType int

	message := []interface{} {
		&messageType,
		&this.Realm,
		&this.Details,
	}
	err = json.Unmarshal(str, &message)

	if err != nil {
		return err
	}
	if messageType != MSG_HELLO {
		return ErrWrongMessageType
	}

	return
}


