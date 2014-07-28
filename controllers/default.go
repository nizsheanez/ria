package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"ria/conf"
	"github.com/gorilla/websocket"
	"net/http"
	"io"
	"log"
	"time"
	"unicode/utf8"
	"errors"
)

type MainController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// echoCopy echoes messages from the client using io.Copy.
func echoCopy(w http.ResponseWriter, r *http.Request, writerOnly bool) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, r, err := conn.NextReader()
		if err != nil {
			if err != io.EOF {
				log.Println("NextReader:", err)
			}
			return
		}
		if mt == websocket.TextMessage {
			r = &validator{r: r}
		}
		w, err := conn.NextWriter(mt)
		if err != nil {
			log.Println("NextWriter:", err)
			return
		}
		if mt == websocket.TextMessage {
			r = &validator{r: r}
		}
		if writerOnly {
			_, err = io.Copy(struct{ io.Writer }{w}, r)
		} else {
			_, err = io.Copy(w, r)
		}
		if err != nil {
			if err == errInvalidUTF8 {
				conn.WriteControl(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, ""),
					time.Time{})
			}
			log.Println("Copy:", err)
			return
		}
		err = w.Close()
		if err != nil {
			log.Println("Close:", err)
			return
		}
	}
}


// echoReadAll echoes messages from the client by reading the entire message
// with ioutil.ReadAll.
func echoReadAll(w http.ResponseWriter, r *http.Request, writeMessage bool) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, b, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Println("NextReader:", err)
			}
			return
		}
		if mt == websocket.TextMessage {
			if !utf8.Valid(b) {
				conn.WriteControl(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, ""),
					time.Time{})
				log.Println("ReadAll: invalid utf8")
			}
		}
		if writeMessage {
			err = conn.WriteMessage(mt, b)
			if err != nil {
				log.Println("WriteMessage:", err)
			}
		} else {
			w, err := conn.NextWriter(mt)
			if err != nil {
				log.Println("NextWriter:", err)
				return
			}
			if _, err := w.Write(b); err != nil {
				log.Println("Writer:", err)
				return
			}
			if err := w.Close(); err != nil {
				log.Println("Close:", err)
				return
			}
		}
	}
}

func (this *MainController) Get() {
	if val := this.Ctx.Request.Header.Get("Upgrade"); val == "websocket" {
		// Upgrade from http request to WebSocket.
		headers := http.Header{"Set-Cookie": {"sessionID=1234"}, "Sec-WebSocket-Protocol": {"wamp"}}
		ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, headers)
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
			return
		} else if err != nil {
			beego.Error("Cannot setup WebSocket connection:", err)
			return
		}
		defer ws.Close()

		for {
			beego.Info("wait for message...")
			mt, p, err := ws.NextReader()
			beego.Info("read:", mt, p, err)
			if err != nil {
				if err != io.EOF {
					beego.Error("read:", err)
				}
				return
			}

			w, err := ws.NextWriter(mt)
			beego.Info("write:", w, err)
			if err != nil {
				return
			}


			if mt == websocket.TextMessage {
			}

//			err = conn.WriteMessage(mt, b)
//			if err != nil {
//				beego.Info("WriteMessage:", err)
//			}

		}
		this.EnableRender = false
	} else {
		this.Data["Css"] = conf.GetCss()
		this.Data["Js"] = conf.GetJs()

		this.Data["Storage"] = template.JS(`{"init":{"categories":{"1":{"id":1,"name":"Professional","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"2":{"id":2,"name":"Health","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"3":{"id":3,"name":"Own","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"4":{"id":4,"name":"Global","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"}},"goals":[{"id":1,"title":"Quis quis voluptatibus voluptatibus eius non fuga nesciunt sit nesciunt omnis.","status":1,"completed":0,"fk_user":1,"reason":"<div>1 Hello 73179876<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>1 Hello 66081062<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":405,"description":"<div>&nbsp;<\/div>","fk_goal":1,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":406,"description":"<div>&nbsp;<\/div>","fk_goal":1,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":2,"title":"Voluptates omnis alias similique eius sequi omnis voluptatibus voluptatum sunt fuga ut.","status":1,"completed":0,"fk_user":1,"reason":"<div>2 Hello 85695198<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>2 Hello 41520369<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":407,"description":"<div>&nbsp;<\/div>","fk_goal":2,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":408,"description":"<div>&nbsp;<\/div>","fk_goal":2,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":3,"title":"Nesciunt et saepe quae velit vitae magnam sunt et facere voluptas amet dolorum porro eveniet veritatis mollitia nobis dolores eligendi non eligendi vel nisi.","status":1,"completed":0,"fk_user":1,"reason":"<div>3 Hello 28844339<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>3 Hello 32418528<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":409,"description":"<div>&nbsp;<\/div>","fk_goal":3,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":410,"description":"<div>&nbsp;<\/div>","fk_goal":3,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":4,"title":"Dolor laborum assumenda asperiores est esse sint molestiae omnis ad vel hic officia dolor quis.","status":1,"completed":0,"fk_user":1,"reason":"<div>4 Hello 87864320<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>4 Hello 54513806<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":411,"description":"<div>&nbsp;<\/div>","fk_goal":4,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":412,"description":"<div>&nbsp;<\/div>","fk_goal":4,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":5,"title":"Vitae deleniti a quia est modi rerum earum aut error necessitatibus sit voluptatum sunt et delectus dolorem quaerat repellendus et mollitia.","status":1,"completed":0,"fk_user":1,"reason":"<div>5 Hello 66231599<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>5 Hello 27761309<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":413,"description":"<div>&nbsp;<\/div>","fk_goal":5,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":414,"description":"<div>&nbsp;<\/div>","fk_goal":5,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":6,"title":"Et et omnis omnis consectetur magni consequatur nemo et est consequatur dolor voluptatem quis adipisci enim error quas.","status":1,"completed":0,"fk_user":1,"reason":"<div>6 Hello 15783726<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>6 Hello 25028690<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":415,"description":"<div>&nbsp;<\/div>","fk_goal":6,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":416,"description":"<div>&nbsp;<\/div>","fk_goal":6,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":7,"title":"Facere labore possimus blanditiis sed culpa voluptates totam.","status":1,"completed":0,"fk_user":1,"reason":"<div>7 Hello 67776443<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>7 Hello 88620491<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":417,"description":"<div>&nbsp;<\/div>","fk_goal":7,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":418,"description":"<div>&nbsp;<\/div>","fk_goal":7,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":8,"title":"Laborum quo ex ipsum sit fugiat voluptatem ut qui doloremque assumenda dignissimos.","status":1,"completed":0,"fk_user":1,"reason":"<div>8 Hello 75297424<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>8 Hello 83846786<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":419,"description":"<div>&nbsp;<\/div>","fk_goal":8,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":420,"description":"<div>&nbsp;<\/div>","fk_goal":8,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":9,"title":"Ea sit quia ipsa nesciunt aliquam sit dolorum ullam eum alias provident aliquam voluptatem quia est eos exercitationem et error.","status":1,"completed":0,"fk_user":1,"reason":"<div>9 Hello 60701114<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>9 Hello 78882674<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":421,"description":"<div>&nbsp;<\/div>","fk_goal":9,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":422,"description":"<div>&nbsp;<\/div>","fk_goal":9,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":10,"title":"Sequi quo temporibus aut perferendis temporibus autem est quis consequatur rerum sit neque eum et ipsam voluptas odit occaecati quisquam voluptate eos suscipit eum.","status":1,"completed":0,"fk_user":1,"reason":"<div>1 Hello 73179876<\/div>","create_time":null,"update_time":"2014-05-19 10:43:15","decomposition":"<div>10 Hello 93704443<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":4,"today":{"report":{"id":423,"description":"<div>&nbsp;<\/div>","fk_goal":10,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":424,"description":"<div>&nbsp;<\/div>","fk_goal":10,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}}],"conclusions":{"today":{"id":117,"description":null,"fk_user":1,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"},"yesterday":{"id":118,"description":null,"fk_user":1,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}}}`)

		this.TplNames = "index.tpl"
		return
	}
}



type validator struct {
	state int
	x     rune
	r     io.Reader
}

var errInvalidUTF8 = errors.New("invalid utf8")


func (r *validator) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	state := r.state
	x := r.x
	for _, b := range p[:n] {
		state, x = decode(state, x, b)
		if state == utf8Reject {
			break
		}
	}
	r.state = state
	r.x = x
	if state == utf8Reject || (err == io.EOF && state != utf8Accept) {
		return n, errInvalidUTF8
	}
	return n, err
}

// UTF-8 decoder from http://bjoern.hoehrmann.de/utf-8/decoder/dfa/
//
// Copyright (c) 2008-2009 Bjoern Hoehrmann <bjoern@hoehrmann.de>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
var utf8d = [...]byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 00..1f
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 20..3f
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 40..5f
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 60..7f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, // 80..9f
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, // a0..bf
	8, 8, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // c0..df
	0xa, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x3, 0x4, 0x3, 0x3, // e0..ef
	0xb, 0x6, 0x6, 0x6, 0x5, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, // f0..ff
	0x0, 0x1, 0x2, 0x3, 0x5, 0x8, 0x7, 0x1, 0x1, 0x1, 0x4, 0x6, 0x1, 0x1, 0x1, 0x1, // s0..s0
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, // s1..s2
	1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, // s3..s4
	1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, // s5..s6
	1, 3, 1, 1, 1, 1, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // s7..s8
}

const (
	utf8Accept = 0
	utf8Reject = 1
)

func decode(state int, x rune, b byte) (int, rune) {
	t := utf8d[b]
	if state != utf8Accept {
		x = rune(b&0x3f) | (x << 6)
	} else {
		x = rune((0xff >> t) & b)
	}
	state = int(utf8d[256+state*16+int(t)])
	return state, x
}
