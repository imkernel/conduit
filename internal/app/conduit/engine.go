package conduit

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

// StartServer starts the API server
func StartServer() {
	done := make(chan bool)
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("no dev.env file found")
	}
	go startHTTPServer()
	go startListener()
	<-done
}

func startListener() {
	var (
		host      string
		port      string
		userName  string
		password  string
		queueName string
	)
	if host = os.Getenv("AMQP_HOST"); host == "" {
		host = "localhost"
	}

	if port = os.Getenv("AMQP_PORT"); port == "" {
		port = "5672"
	}
	if userName = os.Getenv("AMQP_USERNAME"); userName == "" {
		userName = "alwin"
	}
	if password = os.Getenv("AMQP_PASSWORD"); password == "" {
		password = "OpenSesame"
	}
	if queueName = os.Getenv("QUEUE_NAME"); queueName == "" {
		queueName = "OpenSesame"
	}
	log.Printf("UserName: %s", userName)
	log.Printf("Host: %s", host)
	log.Printf("Post: %s", port)
	log.Printf("Password: %s", password)
	dialString := fmt.Sprintf("amqp://%s:%s@%s:%s//", l.options.UserName, l.options.Password, l.options.Host, l.options.Port)
	conn, err := amqp.Dial(dialString)
	if err != nil {
		err = fmt.Errorf("unable to dial: %w", err)
		return err
	}
	defer conn.Close()
	options := &ListenerOptions{
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: password,
		Dialer:   conn,
	}
	msgCh := make(chan *Message, 100)
	listener := NewListener(options)
	err := listener.Listen(msgCh)
	if err != nil {
		log.Fatal(err)
	}

}

func startHTTPServer() {
	port := os.Getenv("HTTP_PORT")
	router := mux.NewRouter()
	createRoutes(router)
	log.Printf("Running Server on Port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
