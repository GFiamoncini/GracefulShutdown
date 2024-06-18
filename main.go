package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3000"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello Word"))
	})

	// Go rotine !
	go func() {
		fmt.Println("Servidor esta rodando em http://localhost:3000")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Não foi possivel iniciar o servidor %s: %v\n", server.Addr, err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Desligando")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Não foi possivel Desligar o servidor: %v\n", err)
	}
	fmt.Println("Servidor Parou de Responder")
}
