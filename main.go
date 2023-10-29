package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/ReidMason/wave-function-collapse/internal/board"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var interval = 1000

var ws *websocket.Conn

type Todo struct {
	Title     string
	id        int
	userId    int
	completed bool
}

func main() {
	http.HandleFunc("/", getIndex)

	http.HandleFunc("/setInterval", setInterval)
	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	go sendStuff()

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func setInterval(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intervalString := r.Form.Get("interval")
	newInterval, err := strconv.Atoi(intervalString)
	if err != nil {
		w.Write([]byte("Failed to update interval"))
		log.Println(err)
		return
	}

	interval = newInterval

	w.Write([]byte("Interval updated"))
}

func sendStuff() {
	var boardData *board.Board
	for {
		if ws != nil {
			if boardData == nil {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				boardData = board.New(100, r)
			}

			data := boardData.Display()
			if boardData.Iter() {
				boardData = nil
				continue
			}

			payload, err := json.Marshal(data)
			if err != nil {
				continue
			}

			ws.WriteMessage(1, payload)
		}
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}

func doWork(boardData *board.Board) {
	for !boardData.Iter() {
		time.Sleep(time.Millisecond * 10)
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ = upgrader.Upgrade(w, r, nil)
}

func getIndex(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}
