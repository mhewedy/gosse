package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type broker struct {
	clients  map[*http.Request]http.ResponseWriter
	dataChan chan string
}

func newBroker() *broker {
	return &broker{
		clients:  make(map[*http.Request]http.ResponseWriter),
		dataChan: make(chan string, 100),
	}
}

func (b *broker) publish(data string) {
	b.dataChan <- data
}

func (b *broker) broadcast() {
	fmt.Println("broadcasting to", len(b.clients), "clients connected")
	data := <-b.dataChan
	for _, w := range b.clients {
		fmt.Fprintf(w, "data: %s\n\n", data)
		f, _ := w.(http.Flusher)
		f.Flush()
	}
}

func (b *broker) addClient(w http.ResponseWriter, r *http.Request) {
	b.clients[r] = w
	fmt.Printf("client %p connected\n", r)
}

func (b *broker) removeClient(w http.ResponseWriter, r *http.Request) {
	delete(b.clients, r)
	fmt.Printf("client %p disconnected\n", r)
}

func main() {

	b := newBroker()

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		b.publish(string(data))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")

		b.addClient(w, r)

		go func() {
			<-r.Context().Done()
			b.removeClient(w, r)
		}()

		for {
			b.broadcast()
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
