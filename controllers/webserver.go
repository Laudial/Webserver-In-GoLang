package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"sync"
)

// Couleurs pour les messages de démarrage et d'arrêt du serveur web
const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
)

// StartServer initialise et démarre le serveur HTTP.
func StartMainServer() {
	var wg sync.WaitGroup

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Incrémenter le WaitGroup avant de lancer la goroutine
	wg.Add(1)
	go waitForTermination(port, &wg)

	mux := http.NewServeMux()

	// Définir les routes

	// Public
	mux.Handle("/", http.HandlerFunc(PublicHandler))

	// Admin
	mux.Handle("/admin", http.HandlerFunc(AdminHandler))

	log.Fatal(http.ListenAndServe(":"+port, mux))

	// Attendre la goroutine waitForTermination avant de quitter
	wg.Wait()
}

func waitForTermination(port string, wg *sync.WaitGroup) {
	// Décrémenter le WaitGroup lorsque la goroutine se termine
	defer wg.Done()

	fmt.Println(colorGreen + "Le serveur web est maintenant en ligne sur le port", port ,"\nAppuyez sur Ctrl+C pour arrêter" + colorReset)

	// Créer un canal pour recevoir les signaux d'arrêt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println(colorRed + "\nArrêt du serveur web..." + colorReset)

	os.Exit(0)
}
