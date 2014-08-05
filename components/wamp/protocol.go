package wamp

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"errors"
	"fmt"
	"encoding/json"
	"ria/components/wamp/messages"
)


//[1,"realm1",{
//	"roles":{
//		"caller":{"features":{"caller_identification":true,"progressive_call_results":true}},
//		"callee":{"features":{"progressive_call_results":true}},
//		"publisher":{"features":{"subscriber_blackwhite_listing":true,"publisher_exclusion":true,"publisher_identification":true}},
//		"subscriber":{"features":{"publisher_identification":true}}
//	}
//}]
type Protocol struct {
}

type WampContext interface {
	Call(callId int, procUri string, arguments []interface{})
	Welcome() error
	Subscribe()
	Unsubscribe()
	Done()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	Subprotocols: []string{"wamp.2.json"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *Protocol) OnConnect(request *http.Request, response *http.ResponseWriter) (ws *websocket.Conn, err error) {
	ws, err = this.Upgrade(request, response)
	if err != nil {
		return
	}

	return
}

func (this *Protocol) Upgrade(request *http.Request, response *http.ResponseWriter) (ws *websocket.Conn, err error) {
	// Upgrade from http request to WebSocket.
	headers := http.Header{"Set-Cookie": {"sessionID=1234"}}
	ws, err = upgrader.Upgrade(*response, request, headers)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(*response, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	return ws, nil
}

func (this *Protocol) Welcome(ws *websocket.Conn, id int) (err error) {
	//WMAP Welcome
	beego.Info("Open connection, say Welcome")
	msg := &messages.WelcomeMessage{}
	msg.Id = id
	err = websocket.WriteJSON(ws, msg.Array())
	if err != nil {
		return
	}
	beego.Info("Welcome message sent")

	return nil
}

func (this *Protocol) ReadMessage(ws *websocket.Conn, ctx WampContext) error {
	mt, p, err := ws.ReadMessage()
	if err != nil {
		return err
	}

	if mt != websocket.TextMessage {
		panic(fmt.Sprintf("Can handle only text now, given: %v", mt))
	}

	err = this.OnMessage(p, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (this *Protocol) OnMessage(rawMessage []byte, ctx WampContext) (err error) {
	var msg []interface{}
	err = json.Unmarshal(rawMessage, &msg)
	if err != nil {
		return err
	}

	messageTypeStr, ok := msg[0].(float64)
	if !ok {
		return errors.New(fmt.Sprintf("Cant cast to string: %v", msg[0]))
	}

	switch int(messageTypeStr) {
	case messages.MSG_HELLO:
		message := &messages.HelloMessage{}
		err := message.Unmarshal(rawMessage)
		if err != nil {
			beego.Info(fmt.Sprintf("%s %v %v", rawMessage, err, message))
			return err
		}

		err = ctx.Welcome()
		if err != nil {
			return err
		}
	case messages.MSG_CALL:
		message := &messages.CallMessage{}
		err := message.Unmarshal(rawMessage)
		if err != nil {
			beego.Info(fmt.Sprintf("%s %v %v", rawMessage, err, message))
			return err
		}

		ctx.Call(message.CallId, message.Uri, message.Arguments)
	case messages.MSG_SUBSCRIBE:
		//		ctx.Subscribe(this.getUri(msg[1]));
	case messages.MSG_UNSUBSCRIBE:
		//		ctx.Unsubscribe(this.getUri(msg[1]));
	case messages.MSG_PUBLISH:
		/*
		$exclude = (array_key_exists(3, $json) ? $json[3] : null);
                if (!is_array($exclude)) {
                    if (true === (boolean)$exclude) {
                        $exclude = [$this->context->getId()];
                    } else {
                        $exclude = [];
                    }
                }

                $eligible = (array_key_exists(4, $json) ? $json[4] : []);

                $this->context->publish($this->getUri($json[1]), $json[2], $exclude, $eligible);
		 */
	default:
		return errors.New(fmt.Sprintf("Undefined Message type: %v", messageTypeStr))
	}

	return nil
}

func (this *Protocol) getUri(a interface{}) {

}
