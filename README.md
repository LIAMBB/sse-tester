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
