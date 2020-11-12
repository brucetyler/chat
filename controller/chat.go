package controller

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"chat/model"
	"encoding/json"
)

var wu = &websocket.Upgrader{ReadBufferSize: 512,
		WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }} 


type connection struct{
	ws *websocket.Conn
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

var user model.User;

func (board *Board)broadcast(){
	for{
		select{
		case c := <- board.join:
			board.client[c] = true;
		case msg := <- board.msg:
			user.Content = string(msg)
			data_b, _ := json.Marshal(user);
			for client:= range board.client {
				fmt.Printf("%s\n",data_b)
				fmt.Println(msg)
				client.ws.WriteMessage(websocket.TextMessage,data_b)
			}
		}
	}
}

func (c *connection)reader(){
	for{
		_,message,err := c.ws.ReadMessage();
		if err != nil{
				fmt.Println(err)
				c.ws.Close()
				return
		}
		user.Name = "bruce"
		fmt.Printf("%s\n",message)
		json.Unmarshal(message, &user)
		fmt.Println("user",user)
		board.msg <- message;
	}
}

func Connect(){
	router := mux.NewRouter();
	go board.broadcast();
	router.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request){
		ws,err := wu.Upgrade(w,r,nil);
		if(err!=nil){
			return;
		}
		c := &connection{ws:ws}
		board.join <- c;
		c.reader();
	})
	http.ListenAndServe("127.0.0.1:8000",router)
}