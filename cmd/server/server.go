package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/matheusrb95/endrok/internal/game"
	"github.com/matheusrb95/endrok/types"
)

const TickRate = 20 * time.Millisecond

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	clients    map[net.Conn]struct{}
	game       game.Game
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		clients:    make(map[net.Conn]struct{}),
		game:       game.New(),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()
	go s.sendGameUpdate()

	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection:", conn.RemoteAddr())

		s.clients[conn] = struct{}{}

		s.welcomeMsg(conn)

		go s.handleClient(conn)
	}
}

func (s *Server) welcomeMsg(conn net.Conn) {
	msg := types.Message{
		SessionID: len(s.clients),
		Type:      "welcome",
		Data:      nil,
	}
	js, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshal: ", err)
		return
	}
	conn.Write(js)
}

func (s *Server) handleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		delete(s.clients, conn)
	}()

	buf := make([]byte, 4*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn closed:", conn.RemoteAddr())
			return
		}

		var msg types.Message
		err = json.Unmarshal(buf[:n], &msg)
		if err != nil {
			fmt.Println("error unmarshal:", err)
		}

		s.game.UpdateMoves(msg)
	}

}

func (s *Server) sendGameUpdate() {
	ticker := time.NewTicker(TickRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.game.Update()
			gameData, err := s.game.Serialize()
			if err != nil {
				fmt.Println("Error serializing game:", err)
				continue
			}

			for conn := range s.clients {
				_, err := conn.Write(gameData)
				if err != nil {
					fmt.Printf("Error sending data %s: %s\n", conn.RemoteAddr(), err)
					continue
				}
			}
		case <-s.quitch:
			return
		}
	}
}
