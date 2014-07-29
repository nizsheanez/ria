package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"ria/conf"
	"github.com/gorilla/websocket"
	"ria/components/wamp"
)

type MainController struct {
	connections map[int64]*websocket.Conn
	beego.Controller
}

func (this *MainController) Get() {
	if val := this.Ctx.Request.Header.Get("Upgrade"); val == "websocket" {
		protocol := &wamp.Protocol{Ctx: this}
		id, ws, err := protocol.OnConnect(this.Ctx.Request, &this.Ctx.ResponseWriter)
		if err != nil {
			beego.Error(err)
			return
		}
		this.connections = make(map[int64]*websocket.Conn, 100)
		this.connections[id] = ws

		defer ws.Close()
		protocol.ReadLoop(ws)

		this.EnableRender = false
	} else {
		this.Data["Css"] = conf.GetCss()
		this.Data["Js"] = conf.GetJs()

		this.Data["Storage"] = template.JS(`{"init":{"categories":{"1":{"id":1,"name":"Professional","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"2":{"id":2,"name":"Health","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"3":{"id":3,"name":"Own","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"},"4":{"id":4,"name":"Global","create_time":"2013-12-20 11:32:00","update_time":"0000-00-00 00:00:00"}},"goals":[{"id":1,"title":"Quis quis voluptatibus voluptatibus eius non fuga nesciunt sit nesciunt omnis.","status":1,"completed":0,"fk_user":1,"reason":"<div>1 Hello 73179876<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>1 Hello 66081062<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":405,"description":"<div>&nbsp;<\/div>","fk_goal":1,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":406,"description":"<div>&nbsp;<\/div>","fk_goal":1,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":2,"title":"Voluptates omnis alias similique eius sequi omnis voluptatibus voluptatum sunt fuga ut.","status":1,"completed":0,"fk_user":1,"reason":"<div>2 Hello 85695198<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>2 Hello 41520369<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":407,"description":"<div>&nbsp;<\/div>","fk_goal":2,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":408,"description":"<div>&nbsp;<\/div>","fk_goal":2,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":3,"title":"Nesciunt et saepe quae velit vitae magnam sunt et facere voluptas amet dolorum porro eveniet veritatis mollitia nobis dolores eligendi non eligendi vel nisi.","status":1,"completed":0,"fk_user":1,"reason":"<div>3 Hello 28844339<\/div>","create_time":null,"update_time":"2014-05-19 10:43:13","decomposition":"<div>3 Hello 32418528<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":1,"today":{"report":{"id":409,"description":"<div>&nbsp;<\/div>","fk_goal":3,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":410,"description":"<div>&nbsp;<\/div>","fk_goal":3,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":4,"title":"Dolor laborum assumenda asperiores est esse sint molestiae omnis ad vel hic officia dolor quis.","status":1,"completed":0,"fk_user":1,"reason":"<div>4 Hello 87864320<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>4 Hello 54513806<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":411,"description":"<div>&nbsp;<\/div>","fk_goal":4,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":412,"description":"<div>&nbsp;<\/div>","fk_goal":4,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":5,"title":"Vitae deleniti a quia est modi rerum earum aut error necessitatibus sit voluptatum sunt et delectus dolorem quaerat repellendus et mollitia.","status":1,"completed":0,"fk_user":1,"reason":"<div>5 Hello 66231599<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>5 Hello 27761309<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":413,"description":"<div>&nbsp;<\/div>","fk_goal":5,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":414,"description":"<div>&nbsp;<\/div>","fk_goal":5,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":6,"title":"Et et omnis omnis consectetur magni consequatur nemo et est consequatur dolor voluptatem quis adipisci enim error quas.","status":1,"completed":0,"fk_user":1,"reason":"<div>6 Hello 15783726<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>6 Hello 25028690<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":2,"today":{"report":{"id":415,"description":"<div>&nbsp;<\/div>","fk_goal":6,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":416,"description":"<div>&nbsp;<\/div>","fk_goal":6,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":7,"title":"Facere labore possimus blanditiis sed culpa voluptates totam.","status":1,"completed":0,"fk_user":1,"reason":"<div>7 Hello 67776443<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>7 Hello 88620491<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":417,"description":"<div>&nbsp;<\/div>","fk_goal":7,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":418,"description":"<div>&nbsp;<\/div>","fk_goal":7,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":8,"title":"Laborum quo ex ipsum sit fugiat voluptatem ut qui doloremque assumenda dignissimos.","status":1,"completed":0,"fk_user":1,"reason":"<div>8 Hello 75297424<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>8 Hello 83846786<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":419,"description":"<div>&nbsp;<\/div>","fk_goal":8,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":420,"description":"<div>&nbsp;<\/div>","fk_goal":8,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":9,"title":"Ea sit quia ipsa nesciunt aliquam sit dolorum ullam eum alias provident aliquam voluptatem quia est eos exercitationem et error.","status":1,"completed":0,"fk_user":1,"reason":"<div>9 Hello 60701114<\/div>","create_time":null,"update_time":"2014-05-19 10:43:14","decomposition":"<div>9 Hello 78882674<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":3,"today":{"report":{"id":421,"description":"<div>&nbsp;<\/div>","fk_goal":9,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":422,"description":"<div>&nbsp;<\/div>","fk_goal":9,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}},{"id":10,"title":"Sequi quo temporibus aut perferendis temporibus autem est quis consequatur rerum sit neque eum et ipsam voluptas odit occaecati quisquam voluptate eos suscipit eum.","status":1,"completed":0,"fk_user":1,"reason":"<div>1 Hello 73179876<\/div>","create_time":null,"update_time":"2014-05-19 10:43:15","decomposition":"<div>10 Hello 93704443<\/div>","comments":"<div>&nbsp;<\/div>","fk_goal_category":4,"today":{"report":{"id":423,"description":"<div>&nbsp;<\/div>","fk_goal":10,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}},"yesterday":{"report":{"id":424,"description":"<div>&nbsp;<\/div>","fk_goal":10,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}}],"conclusions":{"today":{"id":117,"description":null,"fk_user":1,"report_date":"2014-07-24","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"},"yesterday":{"id":118,"description":null,"fk_user":1,"report_date":"2014-07-23","create_time":"2014-07-24 10:06:51","update_time":"2014-07-24 10:06:51"}}}}`)

		this.TplNames = "index.tpl"
		return
	}
}

func (this *MainController) Call(callId string, procUri string, params interface{}) {

}

func (this *MainController) Subscribe() {

}

func (this *MainController) Unsubscribe() {

}
