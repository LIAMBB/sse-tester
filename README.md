# **SSE Argo Tunnel Tester**

This repository contains a simple implementation of a **Server-Sent Events (SSE)** server and client, written in **Go**, to test the compatibility and performance of SSE through **Cloudflare Argo Tunnels**.

## **What is SSE?**
Server-Sent Events (SSE) is a standard for enabling real-time, unidirectional communication from a server to a client. It’s commonly used for:
- Live updates (e.g., stock prices, notifications, or chat messages).
- Streaming data to clients without requiring the overhead of WebSockets.

### **How SSE Works**
1. The client establishes an HTTP connection to the server.
2. The server keeps the connection open and streams updates as text in the `text/event-stream` format.
3. The client listens for messages and handles updates in real time.

---

## **Purpose of this Project**
This project is designed to:
1. Provide a minimal setup for testing SSE compatibility with **Cloudflare Argo Tunnels**.
2. Demonstrate how to implement an SSE server in **Go**.
3. Verify that updates are delivered in real-time through the tunnel.

---

## **Features**
- A Go-based SSE server that streams messages every second.
- A simple HTML client to display streamed messages in a browser.
- Logs for client disconnection events to test cleanup.
- Fully compatible with Cloudflare Argo Tunnels.

---

## **How to Use**

### **1. Run the SSE Server**
1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/sse-argo-tester.git
   cd sse-argo-tester
   ```

2. Run the server:
   ```bash
   go run main.go
   ```
3. The server starts on http://localhost:3000.
### **2. Set Up Cloudflare Argo Tunnel**
1. Start an Argo Tunnel:
   ```bash
   cloudflared tunnel --url http://localhost:3000
   ```
2. Note the public URL provided by Cloudflare (e.g., https://yourdomain.cloudflare.com).
### **3. Test the SSE Client**
1. Open the public URL in your browser.
2. You should see an HTML page displaying live messages from the server.

# **RESULT**
Based on the result of this test, it appears that Cloudflare Argo Tunnels do support SSE (Server-Side Events) but not intentionally. The behavior you get from this is Cloudflare initially buffering the server events (for roughly 90 seconds) before allowing the events to stream to the client unencumbered. This is likely a result of an optimization Cloudflare implemented to reduce the number of packets being processed by the their infrastructure by trying to wait for the server to finish it's response before relaying it to the client, which just so happens to cause this buffering effect in the SSE usecase.