package components

import (
	"fmt"
	"io"
	"log"
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego"
	"errors"
	"net/http"
	"net/http/httptest"
//	"io/ioutil"
	"net/url"
	"ria/wamp/messages"
)

const channelBufSize = 100

var maxId int = 0

// Chat client.
type Client struct {
	id             int
	ws            *websocket.Conn
	server        *Server
	messagesCh     chan messages.Message
	doneCh         chan bool
}

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server) *Client {

	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	maxId++
	messgagesCh := make(chan messages.Message, channelBufSize)
	doneCh := make(chan bool)

	return &Client{maxId, ws, server, messgagesCh, doneCh}
}

//WampContext interface
func (this *Client) Call(callId int, uri string, arguments map[string]string) {
//	beego.Info(fmt.Sprintf("Call: %d, %v, %v", callId, uri, arguments))
	resp := httptest.NewRecorder()

	param := &url.Values{}
	for key, val := range arguments {
		param.Set(key, val)
	}
	url := "/"+uri+"?"+param.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		this.server.Err(err)
		return
	}
	beego.BeeApp.Handlers.ServeHTTP(resp, req)
	if resp.Code != 200 {
		this.server.Err(errors.New(fmt.Sprintf("Server return wrong http code: %d", resp.Code)))
		return
	}
//	p, err := ioutil.ReadAll(resp.Body)
//	if  err != nil {
//		this.server.Err(err)
//		return
//	}
//	beego.Info(string(p))


	//	this.server.Handlers.Get(uri)

	//	if err != nil {
	//		this.server.Err(err)
	//	}

	//	msg := &messages.ResultMessage{
	//		CallId: callId,
	//		Data: data,
	//	}
	//	this.Write(msg)
}

func (this *Client) Subscribe() {
	beego.Info("Implement Subscribe!!!")
}

func (this *Client) Unsubscribe() {
	beego.Info("Implement Unsubscribe!!!")
}

func (this *Client) Conn() *websocket.Conn {
	return this.ws
}

func (this *Client) Write(msg messages.Message) {
	select {
	case this.messagesCh <- msg:
	default:
		this.server.Del(this)
		err := fmt.Errorf("client %d is disconnected.", this.id)
		this.server.Err(err)
	}
}

func (this *Client) Done() {
	this.doneCh <- true
}

// Listen Write and Read request via chanel
func (this *Client) Listen() {
	go this.listenWrite()
	this.listenRead()
}

// Listen write request via chanel
func (this *Client) listenWrite() {
	for {
		// send message to the client
		select {
		case rawMsg := <-this.messagesCh:
			message, ok := rawMsg.(*messages.ResultMessage)
			if !ok {
				this.server.Err(errors.New(fmt.Sprintf("Can't case %T to RpcMessage", rawMsg)))
			}

			err := websocket.WriteJSON(this.ws, message.Array())
			if err != nil {
				this.server.Err(err)
			}

			// receive done request
		case <-this.doneCh:
			this.server.Del(this)
		this.doneCh <- true // for listenRead method
			return
		}
	}
}

func (this *Client) Welcome() error {
	err := this.server.protocol.Welcome(this.ws, this.id)
	if err != nil {
		return err
	}

	return nil
}

// Listen read request via chanel
func (this *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

			// receive done request
		case <-this.doneCh:
			this.server.Del(this)
		this.doneCh <- true // for listenWrite method
			return

			// read data from websocket connection
		default:
			err := this.server.protocol.ReadMessage(this.ws, this)
			if err != nil {
				if err == io.EOF {
					this.Done()

					continue
				}

				this.server.Err(err)
			}
		}
	}
}
