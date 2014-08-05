package messages

import (
	"encoding/json"
)

type WelcomeMessage struct{
	MessageBase
	Id           int
	Details      Details
}

func (this *WelcomeMessage) Array() interface{} {
	this.Details = Details{
		Roles: &Roles{
			"broker": &Role{
				&Features{
					"publisher_exclusion": true,
					"publisher_identification": true,
					"subscriber_blackwhite_listing": true,
				},
			},
			"dealer": &Role{
				&Features{
					"caller_identification": true,
					"progressive_call_results": true,
				},
			},
		},
	}

	arr := []interface{}{
		MSG_WELCOME,
		this.Id,
		this.Details,
	}

	return arr
}


func (this *WelcomeMessage) Unmarshal(str []byte) (err error) {
	var messageType int

	message := []interface{} {
		&messageType,
		&this.Id,
		&this.Details,
	}
	err = json.Unmarshal(str, &message)
	if err != nil {
		return err
	}
	if messageType != MSG_WELCOME {
		return ErrWrongMessageType
	}
	return nil
}
