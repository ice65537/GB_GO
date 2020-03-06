package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

//Client -
type Client struct {
	conn     net.Conn
	id       string
	name     string
	sendPath string
}

const (
	sendPathAll  string = "ALL"
	sendPathNone string = "NOBODY"
)

func (cl *Client) setSendPath(s string) {
	cl.sendPath = s
}

func (cl Client) send(message string) error {
	_, err := io.WriteString(cl.conn, message+"\r\n")
	if err != nil {
		return err
	}
	return nil
}

func (cl Client) receive() (string, error) {
	rdr := bufio.NewReader(cl.conn)
	data, _, err := rdr.ReadLine()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//Message -
type Message struct {
	message  string
	sender   string
	receiver string
}

func (m *Message) String() string {
	return fmt.Sprintf(" [%s] to [%s]: %s", m.sender, m.receiver, m.message)
}

//Server -
type Server struct {
	srv      net.Listener
	clients  map[string]Client
	messages chan *Message
}

//NewServer -
func NewServer() (*Server, error) {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return nil, err
	}

	return &Server{
		srv:      l,
		clients:  make(map[string]Client),
		messages: make(chan *Message),
	}, nil
}

//Serve -
func (s *Server) Serve() error {
	log.Println("Starting server...")
	go s.runBroadcast()
	for {
		conn, err := s.srv.Accept()
		if err != nil {
			return err
		}
		cl := Client{
			conn: conn,
			id:   conn.RemoteAddr().String(),
		}
		log.Println("New client connection:", cl)
		if !s.authConn(&cl) {
			conn.Close()
			continue
		}
		go s.serveClient(cl.name)
	}
}

func (s *Server) closeClient(name string) {
	s.clients[name].send("Your connection is closing")
	s.clients[name].conn.Close()
	delete(s.clients, name)
}

func (s *Server) authConn(cl *Client) bool {
	if err := cl.send("Welcome to the chat server!\r\nEnter your name and password in format: username/password"); err != nil {
		log.Println(err)
		return false
	}

	token, err := cl.receive()
	if err != nil {
		log.Println(err)
		return false
	}
	credentials := strings.Split(token+"//", "/")

	if credentials[1] != "friend" {
		cl.send("Unknown user or incorrect password")
		return false
	}
	if _, ok := s.clients[credentials[0]]; ok {
		cl.send("User with name " + credentials[0] + " already registered")
		return false
	}

	if err := cl.send("Authentification complete. Welcome back!\r\n" +
		"Enter your chat message. Commands begins with symbol \"!\"\r\n. Supported command list:\r\n" +
		"!SENDTO=Username|ALL|NOBODY - set receiver name\r\n" +
		"!STATUS - show my status\r\n"); err != nil {
		log.Println(err)
		return false
	}

	cl.name = credentials[0]
	cl.sendPath = sendPathAll
	s.clients[credentials[0]] = *cl

	log.Println("Новый клиент аутентифицирован:", cl.id, "=", cl.name)
	return true
}

func (s *Server) serveClient(name string) {
	for {
		strdata, err := s.clients[name].receive()
		if err != nil {
			log.Println("Ошибка чтения данных клиента", name, ":", err, "связь будет разорвана")
			s.closeClient(name)
			break
		}
		log.Println(strdata)
		if strdata[0] == '!' {
			if x := strings.Index(strings.ToUpper(strdata), "!SENDTO="); x != -1 {
				strdata = strdata[x+len("!SENDTO=") : len(strdata)]
				if _, ok := s.clients[name]; ok || strdata == sendPathAll || strdata == sendPathNone {
					x := s.clients[name]
					x.sendPath = strdata
					s.clients[name] = x
					s.clients[name].send("Receiver= " + strdata + " successfully set")
				} else {
					s.clients[name].send("Unknown user: " + strdata)

				}
				continue
			}
			if strings.ToUpper(strdata) == "!STATUS" {
				s.clients[name].send("You are " + name)
				s.clients[name].send("Receiver= " + s.clients[name].sendPath)
				continue
			}
			s.clients[name].send("Unknown command: " + strdata)
			continue
		}

		s.messages <- &Message{
			message:  strdata,
			sender:   name,
			receiver: s.clients[name].sendPath,
		}
	}
}

func (s *Server) runBroadcast() {
	log.Println("Starting broadcast")
	for {
		select {
		case msg := <-s.messages:
			log.Println(msg)
			if msg.receiver == sendPathNone {
				continue
			}
			for name, clnt := range s.clients {
				if name == msg.sender || msg.receiver != sendPathAll && msg.receiver != clnt.name {
					continue
				}
				log.Println("Send to [" + clnt.name + "]: " + msg.message)
				go clnt.send("[" + msg.sender + "]:" + msg.message)
			}
		}
	}
}

func main() {
	srv, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	err = srv.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
