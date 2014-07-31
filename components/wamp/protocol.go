package wamp

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"errors"
	"fmt"
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

type Message struct{
	Id          int
	Data        map[string]interface{}
}

type Protocol struct {
}

type WampContext interface {
	Call(callId string, procUri string, params map[string]interface{})
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

func (this *Message) String() string {
	r := make([]interface{}, 3)
	arr := append(r, this.Id, this.Data)
	str, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	return string(str)
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
	welcomeMessage, err := json.Marshal([]int{MSG_WELCOME, id})
	if err != nil {
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, welcomeMessage)
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

func (this *Protocol) WriteMessage(ws *websocket.Conn, msg Message) error {
	ws.WriteMessage(1, []byte(msg.String()))
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
	case MSG_HELLO:
		err := ctx.Welcome()
		if err != nil {
			return err
		}
	case MSG_CALL:
		callId, ok := msg[1].(string)
		if !ok {
			return errors.New("Cant parse callId")
		}
		procUri, ok := msg[2].(string)
		if !ok {
			return errors.New("Cant parse procUri")
		}
		params, ok := msg[3].(map[string]interface{})
		if !ok {
			return errors.New(fmt.Sprintf("Can't cast %v to map[string]interface {}", msg[3]))
		}
		ctx.Call(callId, procUri, params)
	case MSG_SUBSCRIBE:
		//		ctx.Subscribe(this.getUri(msg[1]));
	case MSG_UNSUBSCRIBE:
		//		ctx.Unsubscribe(this.getUri(msg[1]));
	case MSG_PUBLISH:
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
