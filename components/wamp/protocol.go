package wamp

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"io"
	"errors"
	"fmt"
	"encoding/json"
)

const (
	MSG_WELCOME = iota
	MSG_PREFIX
	MSG_CALL
	MSG_CALL_RESULT
	MSG_CALL_ERROR
	MSG_SUBSCRIBE
	MSG_UNSUBSCRIBE
	MSG_PUBLISH
	MSG_EVENT
)

type Message struct{
	MessageType int
	Id          int
	Url         string
	Data        []interface{}
}

type Protocol struct {
	Ctx WampContext
}

type WampContext interface {
	Call(callId string, procUri string, params interface{})
	Subscribe()
	Unsubscribe()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *Protocol) OnConnect(request *http.Request, response *http.ResponseWriter) (id int64, ws *websocket.Conn, err error) {
	id, ws, err = this.Upgrade(request, response)
	if err != nil {
		return
	}

	err = this.welcome(ws)
	if err != nil {
		beego.Error(err)
		return
	}

	return
}

func (this *Protocol) ReadLoop(ws *websocket.Conn) (err error) {
	for {
		beego.Info("wait for message...")
		mt, rawMessage, err := ws.ReadMessage()
		if err != nil {
			if err != io.EOF {
				beego.Error("read:", err)
			}
			return err
		}

		if mt == websocket.TextMessage {
			err = this.onMessage(rawMessage)
			if err != nil {
				beego.Error(err)
				return err
			}
		} else {
			panic("Can handle only text now")
		}
	}
}

func (this *Protocol) Upgrade(request *http.Request, response *http.ResponseWriter) (id int64, ws *websocket.Conn, err error) {
	// Upgrade from http request to WebSocket.
	headers := http.Header{"Set-Cookie": {"sessionID=1234"}, "Sec-WebSocket-Protocol": {"wamp"}}
	ws, err = upgrader.Upgrade(*response, request, headers)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(*response, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	id = 1

	return id, ws, nil
}


func (this *Protocol) welcome(ws *websocket.Conn) (err error) {
	//WMAP Welcome
	beego.Info("Open connection, say Welcome")
	welcomeMessage, err := json.Marshal([]int{0, 1})
	if err != nil {
		return
	}

	err = ws.WriteMessage(1, welcomeMessage)
	if err != nil {
		return
	}
	beego.Info("Welcome message sent")

	return nil
}

func (this *Protocol) onMessage(rawMessage []byte) (err error) {
	var msg []interface{}
	err = json.Unmarshal(rawMessage, &msg)
	if err != nil {
		return
	}

	messageTypeStr, ok := msg[0].(float64)
	if !ok {
		return errors.New(fmt.Sprintf("Cant cast to string: %v", msg[0]))
	}

	switch int(messageTypeStr) {
	case MSG_PREFIX:
		//$this->context->setPrefix($json[1], $json[2]);
	case MSG_CALL:
		callId, ok := msg[1].(string)
		if !ok {
			return errors.New("Cant parse callId")
		}
		procUri, ok := msg[2].(string)
		if !ok {
			return errors.New("Cant parse procUri")
		}
		params := msg[3]

		this.Ctx.Call(callId, procUri, params)
	case MSG_SUBSCRIBE:
		//                $this->context->subscribe($this->getUri($json[1]));
	case MSG_UNSUBSCRIBE:
		//	$this->context->unsubscribe($this->getUri($json[1]));
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
