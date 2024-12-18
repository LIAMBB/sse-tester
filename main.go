package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Pragma", "no-cache") // Prevent caching in older HTTP caches

	// Send an initial message
	fmt.Fprintf(w, "data: Connected to SSE server\n\n")
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Println("Failed to flush initial message")
	}
	fmt.Println("New client connection")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			log.Println("Client disconnected")
			return
		case t := <-ticker.C:
			// Send periodic updates
			message := fmt.Sprintf("data: Server time is %s\n\n", t.Format(time.RFC3339))
			fmt.Fprint(w, message)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			} else {
				log.Println("Failed to flush data")
			}
		}
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	// Serve a simple client HTML
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>SSE Test</title>
		</head>
		<body>
			<h1>Server-Sent Events Test</h1>
			<div id="messages"></div>
			<script>
				const eventSource = new EventSource('/events');
				
				eventSource.onmessage = (event) => {
					const messagesDiv = document.getElementById('messages');
					const newMessage = document.createElement('div');
					
					// Get the current client time when the message is received
					const receivedTime = new Date().toLocaleTimeString();
					
					// Append the message and the time it was received
					newMessage.textContent = "Received at" +  receivedTime + ": " + event.data;
					messagesDiv.appendChild(newMessage);
				};
				
				eventSource.onerror = (error) => {
					console.error('Error:', error);
				};
			</script>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	// Register handlers
	http.HandleFunc("/events", sseHandler)
	http.HandleFunc("/", htmlHandler)

	// Start the server
	port := ":3000"
	fmt.Printf("SSE server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
