package messages

import (
	"encoding/json"
)

type ResultMessage struct {
	MessageBase
	CallId int
	Options map[string]interface {}
	Data map[string]interface{}
}

func (this *ResultMessage) Array() interface {} {
	this.Options = map[string]interface{}{
		"progress": false,
	}

	arr := []interface{}{
		MSG_RESULT,
		this.CallId,
		this.Options,
		[]interface{}{ //rpc allow return multiple arguments, we don't use it
			this.Data,
		},
		map[string]interface{}{},
	}

	return arr
}

func (this *ResultMessage) Unmarshal(str []byte) (err error) {
	var messageType int

	message := []interface{} {
		&messageType,
		&this.CallId,
		&this.Data,
	}
	err = json.Unmarshal(str, &message)

	if err != nil {
		return err
	}
	if messageType != MSG_RESULT {
		return ErrWrongMessageType
	}

	return
}


