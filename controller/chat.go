package controller

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
  "github.com/gin-gonic/gin"
	"encoding/json"
	"chat/model"
)

var wu = &websocket.Upgrader{ReadBufferSize: 512,
		WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }} 


type connection struct{
	ws *websocket.Conn
	data *model.User
}

type Board struct{
	client map[*connection]bool
	join chan *connection
	msg chan []byte
}

var board = Board{
	client: make(map[*connection]bool),
	join: make(chan *connection),
	msg: make(chan []byte),
}

func (board *Board)broadcast(){
	for{
		select{
		case cli := <- board.join:
			board.client[cli] = true;
			cli.data.Name = user.Name;
			fmt.Println("cli.data",cli.data)
		case msg := <- board.msg:
			for client:= range board.client {
				client.data.Content = string(msg);
				data := client.data;
				fmt.Println("user.Name",user.Name)
				data.Name = user.Name;
				data_b, _ := json.Marshal(data);
				client.ws.WriteMessage(websocket.TextMessage,data_b)
			}
		}
	}
}

func (cli *connection)reader(){
	for{
		_,message,err := cli.ws.ReadMessage();
		if err != nil{
				fmt.Println(err)
				cli.ws.Close()
				return
		}
		json.Unmarshal(message, &cli.data)
		user.Name = cli.data.Name;
		fmt.Println("user",user.Name)
		board.msg <- message;
	}
}

func Connect(c *gin.Context){
	fmt.Println("连接")
	go board.broadcast();
	ws,err := wu.Upgrade(c.Writer,c.Request,nil);
	if(err!=nil){
		return;
	}
	cli := &connection{ws:ws,data:&model.User{}}
	board.join <- cli;
	cli.reader();
}

// func Connect(c *gin.Context){
// 	fmt.Println("连接")
// 	router := mux.NewRouter();
// 	go board.broadcast();
// 	router.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request){
// 		ws,err := wu.Upgrade(w,r,nil);
// 		if(err!=nil){
// 			return;
// 		}
// 		c := &connection{ws:ws}
// 		board.join <- c;
// 		c.reader();
// 	})
// 	http.ListenAndServe("127.0.0.1:8000",router)
// }