package messages

import "encoding/json"

type CallMessage struct {
	MessageBase
	CallId    int
	Uri       string
	Options map[string]interface{}
	Arguments map[string]string //actually it should be map[string]interface{} but i don't know how mutch it to url.Value which require strings
}

func (this *CallMessage) Array() interface{} {
	arr := []interface{}{
		MSG_CALL,
		this.CallId,
		this.Uri,
		this.Options,
		[]interface{}{
			this.Arguments,
		},
	}
	return arr
}

func (this *CallMessage) Unmarshal(str []byte) (err error) {
	var messageType int

	message := []interface{} {
		&messageType,
		&this.CallId,
		&this.Options,
		&this.Uri,
		&[]interface {}{ //there is spike for beautiful key-value interface
			&this.Arguments,
		},
	}
	err = json.Unmarshal(str, &message)
		if err != nil {
		return err
	}
	if messageType != MSG_CALL {
		return ErrWrongMessageType
	}
	return nil
}
