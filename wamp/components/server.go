package components

import (
	"log"
	"net/http"
	"ria/wamp/messages"
	"github.com/astaxie/beego"
)

// Chat server.
type Server struct {
	protocol  *Protocol
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan messages.Message
	doneCh    chan bool
	errCh     chan error
}

// Create new chat server.
func NewServer() *Server {
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan messages.Message)
	doneCh := make(chan bool)
	errCh := make(chan error)
	protocol := &Protocol{}

	server := &Server{
		protocol,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}

	return server
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg messages.Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendAll(msg messages.Message) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

func (this *Server) ServeHTTP(response http.ResponseWriter, request *http.Request)() {
	ws, err := this.protocol.OnConnect(request, &response)
	if err != nil {
		this.Err(err)
		return
	}

	defer func() {
		err := ws.Close()
		if err != nil {
			this.Err(err)
		}
	}()

	//create client
	client := NewClient(ws, this)
	this.Add(client)
	client.Listen()
}

// Listen and serve.
// It serves client connection and broadcast request.
func (this *Server) ListenAndServe() {
	for {
		select {
		// Add new a client
		case c := <-this.addCh:
			this.clients[c.id] = c

			// del a client
		case c := <-this.delCh:
			log.Println("Delete client")
			delete(this.clients, c.id)

		case err := <-this.errCh:
			beego.Error(err)

		case <-this.doneCh:
			return
		}
	}
}
