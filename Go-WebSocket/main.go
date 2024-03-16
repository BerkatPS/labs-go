package main

import (
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Name  string
	Price string
}

type hub struct {
	clients            map[*websocket.Conn]bool
	clientRegisterChan chan *websocket.Conn // register client channel
	// removal untuk menghapus list dari client channel nya
	ClientRemovalChan chan *websocket.Conn // removeal client channel
	broadcastMessage  chan Message         // Broadcast Message
}

// buat function untuk menjalan routine yang berbeda atau bisa berjalan sendiri.

func (h *hub) runHub() {
	for {
		select {
		// kita tangkap koneksi
		// jika channel yang dikirim adalah client register
		case conn := <-h.clientRegisterChan:
			// cukup kita taro di dalam clients
			h.clients[conn] = true
		// jika ada yang ingin menghapus channel kita bisa mengakses removal channel
		case conn := <-h.ClientRemovalChan:
			delete(h.clients, conn) // key nya ada conn

		// ketika ada yang ngirim dari broadcast channel kita akan melooping seluruh client dari struct client
		case msg := <-h.broadcastMessage:
			for conn := range h.clients {
				_ = conn.WriteJSON(msg) // kita buat err nya itu null
			}
		}
	}
}

func main() {
	// membuat entrypoint dan inisiasi struct
	h := &hub{
		clients:            make(map[*websocket.Conn]bool),
		clientRegisterChan: make(chan *websocket.Conn),
		ClientRemovalChan:  make(chan *websocket.Conn),
		broadcastMessage:   make(chan Message),
	}

	// kita akan menjalankan hub menggunakan goroutine

	go h.runHub() // menjalankan go routine

	app := fiber.New()
	app.Use("/ws", AllowUpgrade)
	app.Use("/ws/bid", websocket.New(BidPrice(h)))

	_ = app.Listen(":3030")
}

// mengupgrade ke protokol upgrade / switching protokol
// protokol http ke protol socket

func AllowUpgrade(ctx *fiber.Ctx) error {
	// jika IswebsocketUpgrade ctx.
	if websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

// established connection
// function ke 2 dia akan mengembalikan function websocket

func BidPrice(h *hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		defer func() {
			// untuk melakukan defer pada clientremoval dan kita hapus ke list dan kita close
			h.ClientRemovalChan <- conn
			_ = conn.Close()
		}()
		//membuat query parameter yang mengandung nama
		name := conn.Query("name", "anonymous")
		h.clientRegisterChan <- conn

		for {
			// membaca message nyqa
			messageType, Type, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if messageType == websocket.TextMessage {
				h.broadcastMessage <- Message{
					Name:  name,
					Price: string(price),
				}
			}
		}
	}
}
