package context

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para bookar o quarto!")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}
}

func Exec() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx) // Aqui tbm é possível cancelar o context context.WithTimeout(ctx, time.Second * 3)

	defer cancel() // defer executa por último

	go func() {
		time.Sleep(time.Second * 3) // Thread que cancela o context se o tempo for menor que no case do BookHotel
		cancel()
	}()

	BookHotel(ctx)
	http_()
	client()
}

func http_() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Request Iniciada!")
	defer log.Println("Request Finalizada!")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processada com sucesso!")
		w.Write([]byte("Request processada com sucesso!"))

	// Caso cancelem antes dos 5 segundos entra neste case e cancela a request
	case <-ctx.Done():
		http.Error(w, "Request cancelada!", http.StatusRequestTimeout)
	}
}

func client() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // cancela a request se ultrapassar o limite de tempo
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "Get", "http://localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
